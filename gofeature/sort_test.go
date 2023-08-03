package gostudy

import (
	"fmt"
	"sort"
	"testing"
)

type Int64Slice []int64

func (a Int64Slice) Len() int           { return len(a) }
func (a Int64Slice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Int64Slice) Less(i, j int) bool { return a[i] < a[j] }

func TestSortInt64(t *testing.T) {
	s := Int64Slice{10, 9, 1,2,5,4,7}
	sort.Sort(s)
	for _, element := range s {
		fmt.Println(element)
	}
	fmt.Println(s[len(s) -1])
}

func Sort(nums []int, i, j int) {
	if i >= j {
		return
	}

	left := i
	right := j
	middle := (left + right) / 2
	stub := nums[middle]

	for i < j {
		for i < j {
			if nums[i] < stub {
				i++
				continue
			}
			break
		}

		for i < j {
			if nums[j] > stub {
				j--
				continue
			}
			break
		}

		nums[i], nums[j] = nums[j], nums[i]
	}

	Sort(nums, left, i)
	Sort(nums, i+1, right)
}

func TestSort(t *testing.T) {
	nums := []int{2, 4, 1, 5, 8}
	Sort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}

func TestSortInt(t *testing.T) {
	nums := []int{2, 4, 1, 5, 8}
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	fmt.Printf("%v", nums)
}