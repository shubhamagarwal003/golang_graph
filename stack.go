package graph

type Stack struct {
	nodes []*Node
	head  int
}

func (s *Stack) Push(node *Node) {
	s.nodes = append(s.nodes, node)
	s.head += 1
}

func (s *Stack) Pop() *Node {
	if s.head >= 0 {
		temp := s.nodes[s.head]
		s.nodes = append(s.nodes[:s.head], s.nodes[s.head+1:]...)
		s.head -= 1
		return temp
	} else {
		return nil
	}
}
