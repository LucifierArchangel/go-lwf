package golwf

import "strings"

type RouterGroup struct {
	r            *Router
	path         string
	middlewwares []Handler
}

func (group *RouterGroup) Use(middlewares ...Handler) {
	group.middlewwares = append(group.middlewwares, middlewares...)
}

func (group *RouterGroup) AddRoute(method string, path string, handlers ...Handler) {
	group.r.AddRoute(method, group.path+strings.Trim(path, "/"), append(group.middlewwares, handlers...)...)
}

func (group *RouterGroup) Get(path string, handlers ...Handler) {
	group.AddRoute("GET", path, handlers...)
}

func (group *RouterGroup) Post(path string, handlers ...Handler) {
	group.AddRoute("POST", path, handlers...)
}

func (group *RouterGroup) Put(path string, handlers ...Handler) {
	group.AddRoute("PUT", path, handlers...)
}

func (group *RouterGroup) Patch(path string, handlers ...Handler) {
	group.AddRoute("PATCH", path, handlers...)
}

func (group *RouterGroup) Delete(path string, handlers ...Handler) {
	group.AddRoute("DELETE", path, handlers...)
}
