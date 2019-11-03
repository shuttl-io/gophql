package gophql

import (
	"fmt"
	"reflect"
)

// VisitorWrapper allows for partial implementation of the Visitor interface. The
// Wrapper implements the full Visitor interface and realVisitor could implement
// the individaul Visitors (like the FieldVisitor)
type VisitorWrapper struct {
	realVisitor interface{}
}

// NewVisitor takes an object that perhaps implements a visitor interface and
// returns a new wrapper that will implement the full visitor interface. You
// must pass this function a pointer
func NewVisitor(realVisitor interface{}) (*VisitorWrapper, error) {
	if reflect.ValueOf(realVisitor).Kind() != reflect.Ptr {
		return nil, fmt.Errorf("must provide a pointer to gophql.NewVisitor")
	}
	return &VisitorWrapper{
		realVisitor: realVisitor,
	}, nil
}
