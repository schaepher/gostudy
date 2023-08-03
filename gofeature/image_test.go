package gostudy

import (
	"image"
	"image/draw"
	"image/jpeg"
	"io/fs"
	"log"
	"os"
	"path"
	"path/filepath"
	"sync"
	"testing"

	"github.com/panjf2000/ants/v2"
)

func Trasfer(filename string, split int) (err error) {
	f, err := os.Open(filename)
	if err != nil {
		return
	}
	defer f.Close()

	ext := path.Ext(filename)
	switch ext {
	case ".jpg", ".jpeg":
		var jpg image.Image
		jpg, err = jpeg.Decode(f)
		if err != nil {
			return err
		}

		rgba := jpg.(*image.YCbCr)
		dx := rgba.Rect.Dx()
		dy := rgba.Rect.Dy()
		each := dy / split

		dst := image.NewRGBA(rgba.Rect)

		for currentDy := dy; currentDy > 0; currentDy -= each {
			sub := rgba.SubImage(image.Rect(0, currentDy-each, dx, currentDy)).(*image.YCbCr)
			sub.Rect = image.Rect(0, dy-currentDy, dx, dy-currentDy+each)
			draw.Draw(dst, dst.Bounds(), sub, image.Point{}, draw.Over)
		}

		out, _ := os.Create("p_" + filename)
		defer out.Close()
		err = jpeg.Encode(out, dst, nil)
	default:
		return
	}

	return
}

func TestRefactorImage(t *testing.T) {
	split := 10
	// if len(os.Args) == 2 {
	// 	splitS := os.Args[1]
	// 	split, _ = strconv.Atoi(splitS)
	// }

	wg := &sync.WaitGroup{}
	p, _ := ants.NewPool(10)

	filepath.WalkDir("/e/Codes/Go/golang-test", func(filename string, d fs.DirEntry, err error) error {
		ext := path.Ext(filename)
		switch ext {
		case ".jpg", ".jpeg":
			wg.Add(1)
			p.Submit(func() {
				err := Trasfer(filename, split)
				if err != nil {
					log.Println(err)
				}
				wg.Done()
			})
		}
		return nil
	})

	wg.Wait()
}
