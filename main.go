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

  fmt.Println("---------------------------------------")
	fmt.Println("Árvore original em ordem:")
	root.PrintInOrder(root, 0)

	heigth := root.Heigth()
	fmt.Println("---------------------------------------")
  fmt.Println("Altura de {}", heigth)

	smallest := root.Smallest()
	fmt.Println("---------------------------------------")
  fmt.Println("Menor: {}", smallest)

	greatest := root.Greatest()
	fmt.Println("---------------------------------------")
  fmt.Println("Maior: {}", greatest)
  
  fmt.Println("---------------------------------------")
  fmt.Println("Checa se a árvore é uma BST válida")
  fmt.Println("Is BST: {}", root.IsBST(nil, nil))
	
  fmt.Println("---------------------------------------")
  fmt.Println("Árvore invertida:")
	root.InvertBTree()
	root.PrintInOrder(root, 0)
}
