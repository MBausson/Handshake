package graph

type Node struct {
	Data string
	Root *Node
	Next []*Node
}
