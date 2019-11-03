package graphql

/*
#cgo pkg-config: libgraphqlparser
#include "c/GraphQLAst.h"
#include "c/GraphQLAstNode.h"
#include "c/GraphQLAstVisitor.h"
#include "c/GraphQLParser.h"
#include <stdlib.h>

*/
import "C"

import (
	"errors"
	"io"
	"io/ioutil"
	"unsafe"

	gophql "github.com/shuttl-io/gophql/pkg"
)

// A private internal visitor to track error states
type visitor struct {
	visitor gophql.Visitor
	err     error
}

func parse(query string) (*C.struct_GraphQLAstNode, error) {
	graphql := C.CString(query)
	cError := (*C.char)(nil)
	ast := C.graphql_parse_string(graphql, &cError)
	C.free(unsafe.Pointer(graphql))

	if ast == nil {
		err := errors.New(C.GoString(cError))
		C.graphql_error_free(cError)
		return nil, err
	}
	return ast, nil
}

// Parse takes a reader and parses the graphql string. Returns a graphql.AST and
// an error if applicable
func Parse(rc io.ReadCloser) (*AST, error) {
	b, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, err
	}
	ast, err := parse(string(b))
	if err != nil {
		return nil, err
	}
	return &AST{
		ast: ast,
	}, nil
}
