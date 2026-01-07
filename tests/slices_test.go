package tests

import (
	"fmt"
	"slices"
	"testing"
)

func TestSlice_3(t *testing.T) {
	my_slice := make([]int, 3, 5)
	println(my_slice)
}

func TestSlice_4(t *testing.T) {
	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	mySlice = slices.Delete(mySlice, 2, 3)
	fmt.Println(mySlice)
}

// TestSlice_5 removes first item
func TestSlice_5(t *testing.T) {
	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	mySlice = slices.Delete(mySlice, 0, 1)
	fmt.Println(mySlice)
}

// TestSlice_6 removes last item
func TestSlice_6(t *testing.T) {
	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	mySlice = slices.Delete(mySlice, len(mySlice)-1, len(mySlice))
	fmt.Println(mySlice)
}
