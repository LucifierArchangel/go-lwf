package golwf

import (
	"fmt"
	"net/http"
	"sort"
	"strings"
)

type Router struct {
	NotFoundHandler Handler

	middlewares []Handler
	routeTable  map[string][]*Node
}

func NewRouter() *Router {
	return &Router{NotFoundHandler: defaultNotFoundHandler, routeTable: make(map[string][]*Node)}
}

func (router *Router) Print() {
	var rs Routes

	for method, roots := range router.routeTable {
		for _, root := range roots {
			var segment string

			if root.isFixed {
				segment = root.segment
			} else {
				segment = "{" + root.segment + "}"
			}
			bfs(root, method, "/"+segment, &rs)
		}
	}

	sort.Sort(rs)

	for _, r := range rs {
		fmt.Println(r)
	}
}

func (router *Router) Use(middlewares ...Handler) {
	router.middlewares = append(router.middlewares, middlewares...)
}

func (router *Router) AddRoute(method string, path string, handlers ...Handler) {
	segments := splitPath(path, make([]string, 0, defaultNumOfSegments))
	nodes := make([]*Node, len(segments))

	for i, segment := range segments {
		switch {
		case varSegmentPattern1.MatchString(segment):
			nodes[i] = &Node{isFixed: false, segment: segment[1:]}
		case varSegmentPattern2.MatchString(segment):
			nodes[i] = &Node{isFixed: false, segment: segment[1 : len(segment)-1]}
		default:
			nodes[i] = &Node{isFixed: true, segment: segment}
		}
	}

	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].children = []*Node{nodes[i+1]}
	}

	nodes[len(nodes)-1].handlers = handlers
	roots := router.routeTable[method]
	nodeExist := false

	for _, root := range roots {
		if merge(root, nodes[0]) {
			nodeExist = true
			break
		}
	}

	if !nodeExist {
		router.routeTable[method] = append(roots, nodes[0])
	}
}

func (router *Router) Get(path string, handlers ...Handler) {
	router.AddRoute("GET", path, handlers...)
}

func (router *Router) Post(path string, handlers ...Handler) {
	router.AddRoute("POST", path, handlers...)
}

func (router *Router) Put(path string, handlers ...Handler) {
	router.AddRoute("PUT", path, handlers...)
}

func (router *Router) Patch(path string, handlers ...Handler) {
	router.AddRoute("PATCH", path, handlers...)
}

func (router *Router) Delete(path string, handlers ...Handler) {
	router.AddRoute("DELETE", path, handlers...)
}

func (router *Router) Any(path string, handlers ...Handler) {
	router.AddRoute("GET", path, handlers...)
	router.AddRoute("POST", path, handlers...)
	router.AddRoute("PUT", path, handlers...)
	router.AddRoute("PATCH", path, handlers...)
	router.AddRoute("DELETE", path, handlers...)
}

func (router *Router) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	ctx := getContext()
	defer putContext(ctx)

	ctx.Params.Reset()

	ctx.Request = req
	ctx.Writer = res

	roots := router.routeTable[req.Method]

	if len(roots) == 0 {
		router.NotFoundHandler(ctx)

		return
	}

	segments := splitPath(req.URL.Path, make([]string, 0, defaultNumOfSegments))
	path := findPath(make([]NodeWrapper, 0, defaultNumOfSegments), roots, segments)

	if len(path) == 0 || len(path[len(path)-1].handlers) == 0 {
		router.NotFoundHandler(ctx)

		return
	}

	for i, n := range path {
		if !n.isFixed {
			ctx.Params.Set(n.segment, segments[i])
		}
	}

	for _, middleware := range router.middlewares {
		middleware(ctx)
	}

	for _, handler := range path[len(path)-1].handlers {
		handler(ctx)
	}

	ctx.Writer.Write([]byte(ctx.Response.Body))
}

func (router *Router) Group(path string) *RouterGroup {
	return &RouterGroup{r: router, path: strings.Trim(path, "/") + "/"}
}
