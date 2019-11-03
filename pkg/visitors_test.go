package gophql

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockedFieldVisitor struct {
	mock.Mock
}

func (m *MockedFieldVisitor) VisitField() error {
	m.Called()
	return nil
}

func TestWrapperCallsVisitor(t *testing.T) {
	assert := assert.New(t)
	testObj := new(MockedFieldVisitor)
	w, err := NewVisitor(testObj)
	assert.NoError(err)
	assert.NotNil(w)
	testObj.On("VisitField").Return(nil)
	w.VisitField(&FieldNode{Name: "xxx"})
	testObj.AssertCalled(t, "VisitField")
}
