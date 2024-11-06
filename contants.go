package golwf

import (
	"net/http"
	"regexp"
)

const defaultNumOfSegments = 16

var (
	varSegmentPattern1 = regexp.MustCompile("^:[^/]+$")
	varSegmentPattern2 = regexp.MustCompile("^{[^/]+}$")
)

var methodOrder = map[string]int{
	http.MethodGet:     1,
	http.MethodHead:    2,
	http.MethodPost:    3,
	http.MethodPut:     4,
	http.MethodPatch:   5,
	http.MethodDelete:  6,
	http.MethodConnect: 7,
	http.MethodOptions: 8,
	http.MethodTrace:   9,
}
