package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

// inspired by https://gist.github.com/kaipakartik/8120855?permalink_comment_id=4161303#gistcomment-4161303
func _Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		_Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		_Walk(t.Right, ch)
	}
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	_Walk(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for {
		t1_value, ok1 := <-ch1
		t2_value, ok2 := <-ch2

		if !ok1 && !ok2 {
			return true
		}
		
		if !ok1 || !ok2 {
			return false
		}
		
		if t1_value != t2_value {
			return false
		}
	}
}

func main() {
	fmt.Println("Same(tree.New(1), tree.New(1))", Same(tree.New(1), tree.New(1)))
	fmt.Println("Same(tree.New(1), tree.New(2))", Same(tree.New(1), tree.New(2)))
}
