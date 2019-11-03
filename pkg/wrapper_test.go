package gophql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewVisitorErrorsWhenNoPointerIsGiven(t *testing.T) {
	assert := assert.New(t)
	wrapper, err := NewVisitor(2)
	assert.Error(err)
	assert.Nil(wrapper)
}

func TestNewVisitorReturnsANewVisitor(t *testing.T) {
	assert := assert.New(t)
	ptr := "wwww"
	wrapper, err := NewVisitor(&ptr)
	assert.NoError(err)
	assert.NotNil(wrapper)
}

func TestWrapperImplementsVisitorInterface(t *testing.T) {
	assert := assert.New(t)
	assert.Implements((*Visitor)(nil), new(VisitorWrapper))
}
