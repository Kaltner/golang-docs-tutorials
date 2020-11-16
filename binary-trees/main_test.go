package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

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

func Test_Same(t *testing.T) {
	timeNow := time.Now()
	rand1 := rand.New(rand.NewSource(timeNow.Unix()))
	rand2 := rand.New(rand.NewSource(timeNow.Unix()))

	cases := []struct {
		tree1          *tree.Tree
		tree2          *tree.Tree
		expectedResult bool
	}{
		{
			tree1:          tree.New(1),
			tree2:          tree.New(1),
			expectedResult: true,
		},
		{
			tree1:          tree.New(1),
			tree2:          tree.New(2),
			expectedResult: false,
		},
		{
			tree1:          tree.New(rand1.Int()),
			tree2:          tree.New(rand2.Int()),
			expectedResult: true,
		},
		{
			tree1:          tree.New(rand1.Int()),
			tree2:          tree.New(rand1.Int()),
			expectedResult: false,
		},
	}

	for i, c := range cases {
		i, c := i, c
		t.Run(fmt.Sprintf("Test case N%d", i), func(t *testing.T) {
			if Same(c.tree1, c.tree2) != c.expectedResult {
				t.Fatalf("The two test trees are not the same. \n Tree 1: %v \n Tree 2: %v", c.tree1, c.tree2)
			}
		})
	}
}
