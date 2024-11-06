package golwf

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

type Route struct {
	method   string
	path     string
	handlers []Handler
}

func (r Route) String() string {
	handlerNames := make([]string, len(r.handlers))

	for i := range r.handlers {
		handlerNames[i] = runtime.FuncForPC(reflect.ValueOf(r.handlers[i]).Pointer()).Name()
	}

	return fmt.Sprintf("%-6s %s -> %s", r.method, r.path, strings.Join(handlerNames, " + "))
}
