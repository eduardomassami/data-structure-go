package main

import (
	"fmt"

	binarytree "github.com/eduardomassami/data-structure-go/binary-tree"
)

func main() {
	root := binarytree.NewNode(10)
	root.Add(5)
	root.Add(15)
	root.Add(3)
	root.Add(7)
	root.Add(12)
	root.Add(18)
	root.Add(20)
	root.Add(5)

	fmt.Println("Árvore original em ordem:")
	root.PrintInOrder(root, 0)

	fmt.Println("Árvore invertida:")
	root.InvertBTree()
	root.PrintInOrder(root, 0)

	heigth := root.Heigth()
	fmt.Println("Altura de {}", heigth)
}
