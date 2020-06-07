package hello_test

import (
	"testing"

	hello "github.com/dellykaos/github-action-testing"
	"github.com/stretchr/testify/assert"
)

func TestName(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("Hello, delly!", hello.Name("delly"))
	assert.Equal("Hello, darren!", hello.Name("darren"))
	assert.Equal("Hello, world!", hello.Name(""))
}
