package golwf

import "net/http"

type Handler func(ctx *Context)

func HandlerFunc(f http.HandlerFunc) Handler {
	return func(ctx *Context) {
		f(ctx.Writer, ctx.Request)
	}
}

func defaultNotFoundHandler(ctx *Context) {
	ctx.Writer.WriteHeader(http.StatusNotFound)
	ctx.Writer.Write([]byte("404 page not found"))
}
