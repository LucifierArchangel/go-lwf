package golwf

func bfs(n *Node, method string, path string, rs *Routes) {
	if len(n.handlers) > 0 {
		rs.Append(Route{method, path, n.handlers})
	}

	for _, child := range n.children {
		var segment string

		if child.isFixed {
			segment = child.segment
		} else {
			segment = "{" + child.segment + "}"
		}

		bfs(child, method, path+"/"+segment, rs)
	}
}
