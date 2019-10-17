package ast

type NodeStack struct {
	nodes []*Node
}

func (s *NodeStack) Push(node *Node) {
	s.nodes = append(s.nodes, node)
}

func (s *NodeStack) Pop() *Node {
	if len(s.nodes) > 0 {
		node := s.nodes[len(s.nodes)-1]
		s.nodes = s.nodes[:len(s.nodes)-1]
		return node
	}
	return nil
}

func (s *NodeStack) Last() *Node {
	if len(s.nodes) > 0 {
		return s.nodes[len(s.nodes)-1]
	}
	return nil
}
