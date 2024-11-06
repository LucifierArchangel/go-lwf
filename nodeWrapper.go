package golwf

type NodeWrapper struct {
	*Node
	nextChild int
}

func findPath(path []NodeWrapper, roots []*Node, segments []string) []NodeWrapper {
	for _, root := range roots {
		if root.isFixed && root.segment != segments[0] {
			continue
		}
		path = append(path, NodeWrapper{root, 0})

	loop:
		for len(path) > 0 {
			if len(path) == len(segments) {
				return path
			}

			top := &path[len(path)-1]
			for top.nextChild < len(top.children) {
				child := top.children[top.nextChild]
				top.nextChild++
				if !child.isFixed || child.segment == segments[len(path)] {
					path = append(path, NodeWrapper{child, 0})
					continue loop
				}
			}

			path = path[:len(path)-1]
		}

	}

	return nil
}

func splitPath(path string, segments []string) []string {
	var (
		begin = 0
		end   = len(path) - 1
	)

	for begin < len(path) && path[begin] == '/' {
		begin++
	}

	for end >= 0 && path[end] == '/' {
		end--
	}

	if begin >= end {
		return append(segments, "")
	}

	start := begin

	for i := begin; i < end; i++ {
		if path[i] == '/' {
			segments = append(segments, path[start:i])
			start = i + 1
		}
	}

	return append(segments, path[start:end+1])
}
