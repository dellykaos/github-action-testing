package hello_test

import "testing"

import hello "github.com/dellykaos/github-action-testing"

import "github.com/stretchr/testify/assert"

func TestAdd(t *testing.T) {
	assert := assert.New(t)
	t.Run("Add with positive numbers", func(t *testing.T) {
		assert.Equal(hello.Add(10, 10), 20)
	})

	t.Run("Add with negative numbers", func(t *testing.T) {
		assert.Equal(hello.Add(10, -10), 0)
	})
}
