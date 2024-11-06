package golwf

import (
	"encoding/json"
	"net/http"
	"sync"
)

type Context struct {
	Params   Params
	Request  *http.Request
	Writer   http.ResponseWriter
	Response Response
}

var ctxPool = sync.Pool{
	New: func() interface{} {
		return new(Context)
	},
}

func (ctx *Context) JSON(data map[string]interface{}) {
	body, err := json.Marshal(data)

	if err != nil {
		ctx.InternalServerError()
		return
	}

	if ctx.Response.Headers == nil {
		ctx.Response.Headers = make(map[string]string)
	}

	ctx.Response.Headers["Content-Type"] = "application/json"
	ctx.Response.Body = string(body)

	if ctx.Response.Status == 0 {
		ctx.Response.Status = 200
	}

}

func (ctx *Context) InternalServerError() {
	ctx.Response.Status = 500
	ctx.Response.Body = "Internal Server Error"
}

func getContext() *Context {
	return ctxPool.Get().(*Context)
}

func putContext(ctx *Context) {
	ctxPool.Put(ctx)
}
