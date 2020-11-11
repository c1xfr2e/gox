package algorithm

// Color represents the type for color of node.
type Color bool

const red = true
const black = false

// RBNode represents node in a RBTree.
type RBNode struct {
	Color  Color
	Value  int
	Parent *RBNode
	Left   *RBNode
	Right  *RBNode
}

// RBTree represents Red-Black Tree structure.
type RBTree struct {
	Root      *RBNode
	ElemCount int
}

// Insert creates a new node in the tree for the value.
// It inserts the node to the tree and maintains Red-Black property.
func (tree *RBTree) Insert(value int) {
	// Your code goes here.
}
