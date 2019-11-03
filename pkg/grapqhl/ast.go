package graphql

/*
#include "c/GraphQLAst.h"
#include "c/GraphQLAstNode.h"
#include "c/GraphQLAstVisitor.h"
#include "c/GraphQLParser.h"
#include <stdlib.h>
int visit_field_cgo(struct GraphQLAstField *field, void *unused);
*/
import "C"
import (
	gopointer "github.com/mattn/go-pointer"
	gophql "github.com/shuttl-io/gophql/pkg"
)

// AST is the go implementation of the GraphQL ast
type AST struct {
	ast *C.struct_GraphQLAstNode
}

// Accept takes a visitor and allows it to walk the tree
func (a *AST) Accept(v gophql.Visitor) error {
	_visitor := &visitor{
		visitor: v,
		err:     nil,
	}
	p := gopointer.Save(_visitor)
	defer gopointer.Unref(p)
	visitorCallbacks := C.struct_GraphQLAstVisitorCallbacks{
		visit_field: (C.visit_field_func)(C.visit_field_cgo),
	}
	C.graphql_node_visit(a.ast, &visitorCallbacks, p)
	return _visitor.err
}

// Free frees up the C Resources the AST has been using. This must be called at some point
func (a *AST) Free() {
	C.graphql_node_free(a.ast)
	a.ast = nil
}
