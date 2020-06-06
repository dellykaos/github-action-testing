package hello_test

import "testing"

import hello "github.com/dellykaos/github-action-testing"

import "github.com/stretchr/testify/assert"

func TestAdd(t *testing.T) {
	assert := assert.New(t)
	t.Run("Add with positive numbers", func(t *testing.T) {
		assert.Equal(20, hello.Add(10, 10))
	})

	t.Run("Add with negative numbers", func(t *testing.T) {
		assert.Equal(0, hello.Add(10, -10))
	})
}

func TestName(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("Hello, delly!", hello.Name("delly"))
	assert.Equal("Hello, darren!", hello.Name("darren"))
	assert.Equal("Hello, world!", hello.Name(""))
}
