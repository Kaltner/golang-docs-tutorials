package main

import (
	"fmt"
	"testing"

	"golang.org/x/tour/tree"
)

func Test_Walk(t *testing.T) {
	cases := map[int][]int{
		1:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		2:   []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20},
		3:   []int{21, 24, 27, 30, 12, 6, 3, 9, 15, 18},
		255: []int{1785, 1020, 510, 1275, 255, 765, 1530, 2295, 2040, 2550},
	}

	for i, c := range cases {
		i, c := i, c
		t.Run(fmt.Sprintf("Test case %d", i), func(t *testing.T) {
			t.Parallel()
			ch := make(chan int, 10)
			tree := tree.New(i)

			go Walk(tree, ch)

			nums := make([]int, 10)
			for i := 0; i < 10; i++ {
				nums[i] = <-ch
			}

			for _, expectedN := range c {
				foundValue := false
				for _, n := range nums {
					if n == expectedN {
						foundValue = true
						break
					}
				}
				if !foundValue {
					t.Fatalf("Tree do not match with the expected result \n Expected: %v \n Got: %v", c, nums)
				}
			}
		})
	}
}

// func Test_Same(t *testing.T) {
// cases := struct()
// }
