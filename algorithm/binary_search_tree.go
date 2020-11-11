package algorithm

import "fmt"

type Tree struct {
	left   *Tree
	right  *Tree
	parent *Tree
	value  int
}

func create_tree(A []int) *Tree {
	tree := Tree{nil, nil, nil, 0}
	for i, v := range A {
		if i == 0 {
			tree.value = v
		} else {
			cur := &tree
			var parent *Tree = nil
			to_left := false
			for cur != nil {
				parent = cur
				to_left = v <= cur.value
				if to_left {
					cur = cur.left
				} else {
					cur = cur.right
				}
			}
			if to_left {
				parent.left = &Tree{nil, nil, parent, v}
			} else {
				parent.right = &Tree{nil, nil, parent, v}
			}
		}
	}

	return &tree
}

func walk(tree *Tree) {
	if tree == nil {
		return
	}
	walk(tree.left)
	fmt.Println(tree.value)
	walk(tree.right)
}

func delete_node(node *Tree) {
	if node.left == nil {
		if node == node.parent.left {
			node.parent.left = node.right
		} else {
			node.parent.right = node.right
		}
		if node.right != nil {
			node.right.parent = node.parent
		}
	} else if node.right == nil {
		if node == node.parent.left {
			node.parent.left = node.left
		} else {
			node.parent.right = node.left
		}
		if node.left != nil {
			node.left.parent = node.parent
		}
	} else {
		x := node.right
		for x.left != nil {
			x = x.left
		}
		node.value = x.value
		delete_node(x)
	}
}

func TestBST() {
	tree := create_tree([]int{3, 1, 5, 7, 4, 2, 8, 9, 6, 0})
	tree.parent = &Tree{tree, tree, nil, 0}
	walk(tree)

	fmt.Println("-----------------------")

	fmt.Printf("%d %d\n", tree.right.value, tree.right.left.value)

	delete_node(tree.right.left) // 4
	delete_node(tree.right)      // 5

	delete_node(tree.right) // 7

	fmt.Printf("%d %d\n", tree.right.value, tree.right.right.value)
	walk(tree)
}
