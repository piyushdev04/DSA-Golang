package main

import "fmt"

// TreeNode represents a node in the binary search tree.
type TreeNode struct {
	Key   int
	Left  *TreeNode
	Right *TreeNode
}

// BinaryTree represents a binary search tree.
type BinaryTree struct {
	Root *TreeNode
}

// Insert inserts a new key into the binary search tree.
func (bt *BinaryTree) Insert(key int) {
	newNode := &TreeNode{Key: key}

	if bt.Root == nil {
		bt.Root = newNode
	} else {
		insertNode(bt.Root, newNode)
	}
}

func insertNode(root, newNode *TreeNode) {
	if newNode.Key < root.Key {
		if root.Left == nil {
			root.Left = newNode
		} else {
			insertNode(root.Left, newNode)
		}
	} else {
		if root.Right == nil {
			root.Right = newNode
		} else {
			insertNode(root.Right, newNode)
		}
	}
}

func main() {
	bst := BinaryTree{}

	// Insert some nodes
	keys := []int{50, 30, 20, 40, 70, 60, 80}
	for _, key := range keys {
		bst.Insert(key)
	}

	// Print inorder traversal to verify insertion
	fmt.Println("Inoder traversal after insertion:")
	inorderTraversal(bst.Root) // Output: 20 30 40 50 60 70 80
	fmt.Println()
}

// Inorder traversal: Left -> Root -> Right
func inorderTraversal(node *TreeNode) {
	if node != nil {
		inorderTraversal(node.Left)
		fmt.Print(node.Key, " ")
		inorderTraversal(node.Right)
	}
}
