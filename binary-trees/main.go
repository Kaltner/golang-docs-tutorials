package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		go Walk(t.Left, ch)
	}
	if t.Right != nil {
		go Walk(t.Right, ch)
	}
	ch <- t.Value
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	return t1 == t2
}

func main() {
	ch1 := make(chan int, 10)
	t1 := tree.New(1)

	go Walk(t1, ch1)

	ch2 := make(chan int, 10)
	t2 := tree.New(255)

	go Walk(t2, ch2)

	nums1 := make([]int, 10)
	for i := 0; i < 10; i++ {
		nums1[i] = <-ch1
	}

	// fmt.Println(nums1)

	nums2 := make([]int, 10)
	for i := 0; i < 10; i++ {
		nums2[i] = <-ch2
	}
	fmt.Println(nums2)

	// fmt.Println(Same(t1, t1))
	// fmt.Println(Same(t1, t2))
}
