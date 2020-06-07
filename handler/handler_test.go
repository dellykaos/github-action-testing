package handler_test

import (
	"testing"

	"github.com/dellykaos/github-action-testing/handler"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert := assert.New(t)

	assert.NotNil(handler.New())
}
