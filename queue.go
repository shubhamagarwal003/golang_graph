package graph

type Queue struct {
	node []*Node
	head int
}

func (q *Queue) Enqueue(node *Node) {
	q.node = append(q.node, node)
}

func (q *Queue) Dequeue() *Node {
	if q.head < len(q.node) {
		n := q.node[q.head]
		q.head++
		return n
	}
	return nil
}

func (q *Queue) IsEmpty() bool {
	if q.head < len(q.node) {
		return false
	}
	return true
}
