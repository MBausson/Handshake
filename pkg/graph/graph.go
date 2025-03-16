package graph

type Node struct {
	Data string
	Root *Node
	Next []*Node
}

func FindNode(nodes []*Node, data string) *Node {
	for _, node := range nodes {
		if node.Data == data {
			return node
		}
	}

	return nil
}
