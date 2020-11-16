package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walk(t, ch)
	close(ch)
}

func walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		walk(t.Left, ch)
	}
	if t.Right != nil {
		walk(t.Right, ch)
	}
	ch <- t.Value
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	// ch1 := make(chan int, 10)
	// ch2 := make(chan int, 10)

	// for n := range <-ch1 {

	// }
	return true
}

func main() {
	ch1 := make(chan int)
	t1 := tree.New(1)

	go Walk(t1, ch1)

	// ch2 := make(chan int, 10)
	// t2 := tree.New(255)

	// go Walk(t2, ch2)

	nums1 := make([]int, 0)
	// for i := 0; i < 10; i++ {
	// 	nums1[i] = <-ch1
	// }
	for val := range ch1 {
		nums1 = append(nums1, val)
	}

	fmt.Println(nums1)

	// nums2 := make([]int, 10)
	// for i := 0; i < 10; i++ {
	// 	nums2[i] = <-ch2
	// }
	// fmt.Println(nums2)

	// fmt.Println(Same(t1, t1))
	// fmt.Println(Same(t1, t2))
}
