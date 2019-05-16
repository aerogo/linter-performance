package performance_test

import (
	"testing"

	"github.com/aerogo/http/client"
	performance "github.com/aerogo/linter-performance"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	linter := performance.New()
	linter.Begin("eduardurbach.com", "eduardurbach.com")
	response, err := client.Get("https://eduardurbach.com").End()
	assert.NoError(t, err)
	assert.NotNil(t, response)
	linter.End("eduardurbach.com", "eduardurbach.com", response)
}
