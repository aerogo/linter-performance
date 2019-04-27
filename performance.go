package performance

import (
	"time"

	"github.com/aerogo/http/client"
)

// Linter is a performance linter for routes showing size and response time.
type Linter struct {
	start map[string]time.Time
}

// New creates a new performance linter.
func New() *Linter {
	return &Linter{
		start: make(map[string]time.Time),
	}
}

// Begin is called before a page is requested.
func (linter *Linter) Begin(route string, uri string) {
	linter.start[route] = time.Now()
}

// End is called after the page has responded.
func (linter *Linter) End(route string, uri string, response *client.Response) {
	responseTime := time.Since(linter.start[route]).Nanoseconds() / 1000000
	responseSize := float64(len(response.RawBytes())) / 1024
	printResult(route, responseTime, responseSize)
}
