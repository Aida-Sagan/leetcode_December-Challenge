package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func (node *TreeNode) insert(value int) {

	if node == nil {
		return
	}  else if value >= node.Val {

		if node.Left == nil {
			node.Left = &TreeNode{
				Val : value,
				Left : nil,
				Right : nil,
			}
		} else {
			node.Left.insert(value)
		}



	} else {

		if node.Right == nil {
			node.Right = &TreeNode {
				Val : value,
				Left : nil,
				Right : nil,
			}
		} else {
			node.Right.insert(value)
		}
	}
	/*if node.Left == nil {
		node.Left = &TreeNode{
			Left : nil,
			Right : nil,
			Val : value,
		}
	} else {
		node.Left.insert(value)
	}

	if node.Right == nil {
		node.Right = &TreeNode {
			Left : nil,
			Right : nil,
			Val : value,
		}
	} else {
		node.Right.insert(value)
	}*/

}


func (node *TreeNode) printTree() {

	if node != nil {
		fmt.Printf("%d ", node.Val)
		node.Right.printTree()
		node.Left.printTree()
		
	}
	return
}


func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	//если оба дерева пустые => output nil
	if root1 == nil && root2 == nil {
		return nil
	}

	//если один из деревьев пуст, то выводим то дерево, которое не пустое
	if root1 == nil && root2 != nil {
		return root2
	}

	if root1 != nil && root2 == nil {
		return root1
	}

	//если все есть, то суммируем узлы
	if root1 != nil && root2 != nil {
		root1.Val += root2.Val
	}

	if root1.Left != nil && root2.Left != nil {
		mergeTrees(root1.Left, root2.Left)
	}

	if root1.Right != nil && root2.Right != nil {
		mergeTrees(root1.Right, root2.Right)
	}

	//если у первого дерева нет левого потомка,а у второго есть, то присваиваем первому левый потомок второго дерева
	if root1.Left == nil && root2.Left != nil {
		root1.Left = root2.Left
	}

	if root1.Right == nil && root2.Right != nil {
		root1.Right = root2.Right
	}
	

	return root1
}

func main() {
	root1 := &TreeNode{Val: 1}
	//root1.insert(1)
	root1.insert(3)
	root1.insert(2)
	root1.insert(5)


	root2 := &TreeNode{Val : 2}
	root2.insert(1)
	root2.insert(3)
	root2.insert(4)
	
	root2.insert(7)
	root1.printTree()
	fmt.Println("---------------------------")

	root2.printTree()
	fmt.Println("---------------------------")
	res := mergeTrees(root1, root2)
	res.printTree()
	
}