package gostudy

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func TestSimpleServer(t *testing.T) {
	mtx := http.NewServeMux()
	mtx.HandleFunc("/api/app", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Header.Get("X-Real-IP"))
		byt, _ := ioutil.ReadAll(r.Body)
		fmt.Println(string(byt))
		fmt.Println(r.RemoteAddr)
		time.Sleep(10 * time.Second)
		fmt.Fprint(w, "test")
	})

	srv := &http.Server{
		Addr:        ":8090",
		ReadTimeout: 1 * time.Second,
		IdleTimeout: 2 * time.Second,
		Handler:     mtx,
	}

	srv.ListenAndServe()
}
