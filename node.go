package golwf

type Node struct {
	isFixed  bool
	segment  string
	children []*Node
	handlers []Handler
}

func (n *Node) equal(other *Node) bool {
	return n.isFixed == other.isFixed && n.segment == other.segment
}

func (n *Node) hasChild(c *Node) int {
	for i, child := range n.children {
		if c.equal(child) {
			return i
		}
	}

	return -1
}

func merge(n1 *Node, n2 *Node) bool {
	if !n1.equal(n2) {
		return false
	}

	n1.handlers = append(n1.handlers, n2.handlers...)

	for _, child := range n2.children {
		if i := n1.hasChild(child); i != -1 {
			merge(n1.children[i], child)
		} else {
			n1.children = append(n1.children, child)
		}
	}

	return true
}
