package graph

type Edge struct {
	Dest  *Node
	Src   *Node
	Value int
}

type Node struct {
	Value int
	Edges []*Edge
}

func (n *Node) AddEdge(n2 *Node, value int) {
	e := &Edge{Src: n, Dest: n2, Value: value}
	n.Edges = append(n.Edges, e)
}
