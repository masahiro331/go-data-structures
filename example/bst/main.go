package main

import (
	"fmt"

	"github.com/masahiro331/go-data-structures/pkg/bst"
)

func main() {
	root := bst.Node{Value: 5}
	root.Insert(3)
	root.Insert(10)
	root.Insert(4)
	root.Insert(8)
	root.Insert(11)
	root.Inorder()

	fmt.Println(root.Search(6))
	fmt.Println(root.Search(12))
	fmt.Println(root.Search(7))
	fmt.Println(root.Search(4))
	fmt.Println(root.Search(11))

	root.Remove(3)
	root.Inorder()
}
