package graphql

/*
#include "c/GraphQLAst.h"
#include "c/GraphQLAstNode.h"
#include "c/GraphQLAstVisitor.h"
#include "c/GraphQLParser.h"
int visit_field_cgo(struct GraphQLAstField *field, void *visitorPtr);
*/
import "C"

import (
	"unsafe"

	gopointer "github.com/mattn/go-pointer"
	gophql "github.com/shuttl-io/gophql/pkg"
)

//export visitField
func visitField(field *C.struct_GraphQLAstField, visitorPtr unsafe.Pointer) int {
	v := gopointer.Restore(visitorPtr).(*visitor)
	node := &gophql.FieldNode{
		Name: C.GoString(C.GraphQLAstName_get_value(C.GraphQLAstField_get_name(field))),
	}
	if err := v.visitor.VisitField(node); err != nil {
		v.err = err
		return 0
	}
	return 1
}

//export visitDocument
func visitDocument(field *C.struct_GraphQLAstField, visitorPtr unsafe.Pointer) int {
	v := gopointer.Restore(visitorPtr).(*visitor)
	if err := v.visitor.VisitDocument(); err != nil {
		v.err = err
		return 0
	}
	return 1
}

//export visitOperation
func visitOperation(field *C.struct_GraphQLAstField, visitorPtr unsafe.Pointer) int {
	v := gopointer.Restore(visitorPtr).(*visitor)
	if err := v.visitor.VisitOperation(); err != nil {
		v.err = err
		return 0
	}
	return 1
}

//export visitVariableDefinition
func visitVariableDefinition(field *C.struct_GraphQLAstField, visitorPtr unsafe.Pointer) int {
	v := gopointer.Restore(visitorPtr).(*visitor)
	if err := v.visitor.VisitVariableDefinition(); err != nil {
		v.err = err
		return 0
	}
	return 1
}

//export visitSelectionSetVisitor
func visitSelectionSetVisitor(field *C.struct_GraphQLAstField, visitorPtr unsafe.Pointer) int {
	v := gopointer.Restore(visitorPtr).(*visitor)
	if err := v.visitor.VisitSelectionSet(); err != nil {
		v.err = err
		return 0
	}
	return 1
}

//export visitArgument
func visitArgument(field *C.struct_GraphQLAstField, visitorPtr unsafe.Pointer) int {
	v := gopointer.Restore(visitorPtr).(*visitor)
	if err := v.visitor.VisitArgument(); err != nil {
		v.err = err
		return 0
	}
	return 1
}

//export visitFragmentSpread
func visitFragmentSpread(field *C.struct_GraphQLAstField, visitorPtr unsafe.Pointer) int {
	v := gopointer.Restore(visitorPtr).(*visitor)
	if err := v.visitor.VisitFragmentSpread(); err != nil {
		v.err = err
		return 0
	}
	return 1
}

//export visitVariable
func visitVariable(field *C.struct_GraphQLAstField, visitorPtr unsafe.Pointer) int {
	v := gopointer.Restore(visitorPtr).(*visitor)
	if err := v.visitor.VisitVariable(); err != nil {
		v.err = err
		return 0
	}
	return 1
}

//export visitInteger
func visitInteger(field *C.struct_GraphQLAstField, visitorPtr unsafe.Pointer) int {
	v := gopointer.Restore(visitorPtr).(*visitor)
	if err := v.visitor.VisitInteger(); err != nil {
		v.err = err
		return 0
	}
	return 1
}

//export visitFloat
func visitFloat(field *C.struct_GraphQLAstField, visitorPtr unsafe.Pointer) int {
	v := gopointer.Restore(visitorPtr).(*visitor)
	if err := v.visitor.VisitFloat(); err != nil {
		v.err = err
		return 0
	}
	return 1
}

//export visitString
func visitString(field *C.struct_GraphQLAstField, visitorPtr unsafe.Pointer) int {
	v := gopointer.Restore(visitorPtr).(*visitor)
	if err := v.visitor.VisitString(); err != nil {
		v.err = err
		return 0
	}
	return 1
}

//export visitBoolean
func visitBoolean(field *C.struct_GraphQLAstField, visitorPtr unsafe.Pointer) int {
	v := gopointer.Restore(visitorPtr).(*visitor)
	if err := v.visitor.VisitBoolean(); err != nil {
		v.err = err
		return 0
	}
	return 1
}

//export visitNull
func visitNull(field *C.struct_GraphQLAstField, visitorPtr unsafe.Pointer) int {
	v := gopointer.Restore(visitorPtr).(*visitor)
	if err := v.visitor.VisitNull(); err != nil {
		v.err = err
		return 0
	}
	return 1
}

//export visitEnum
func visitEnum(field *C.struct_GraphQLAstField, visitorPtr unsafe.Pointer) int {
	v := gopointer.Restore(visitorPtr).(*visitor)
	if err := v.visitor.VisitEnum(); err != nil {
		v.err = err
		return 0
	}
	return 1
}

//export visitList
func visitList(field *C.struct_GraphQLAstField, visitorPtr unsafe.Pointer) int {
	v := gopointer.Restore(visitorPtr).(*visitor)
	if err := v.visitor.VisitList(); err != nil {
		v.err = err
		return 0
	}
	return 1
}

//export visitObject
func visitObject(field *C.struct_GraphQLAstField, visitorPtr unsafe.Pointer) int {
	v := gopointer.Restore(visitorPtr).(*visitor)
	if err := v.visitor.VisitObject(); err != nil {
		v.err = err
		return 0
	}
	return 1
}

//export visitObjectField
func visitObjectField(field *C.struct_GraphQLAstField, visitorPtr unsafe.Pointer) int {
	v := gopointer.Restore(visitorPtr).(*visitor)
	if err := v.visitor.VisitObjectField(); err != nil {
		v.err = err
		return 0
	}
	return 1
}

//export visitDirective
func visitDirective(field *C.struct_GraphQLAstField, visitorPtr unsafe.Pointer) int {
	v := gopointer.Restore(visitorPtr).(*visitor)
	if err := v.visitor.VisitDirective(); err != nil {
		v.err = err
		return 0
	}
	return 1
}

//export visitNamedType
func visitNamedType(field *C.struct_GraphQLAstField, visitorPtr unsafe.Pointer) int {
	v := gopointer.Restore(visitorPtr).(*visitor)
	if err := v.visitor.VisitNamedType(); err != nil {
		v.err = err
		return 0
	}
	return 1
}

//export visitListType
func visitListType(field *C.struct_GraphQLAstField, visitorPtr unsafe.Pointer) int {
	v := gopointer.Restore(visitorPtr).(*visitor)
	if err := v.visitor.VisitListType(); err != nil {
		v.err = err
		return 0
	}
	return 1
}

//export visitNonNullType
func visitNonNullType(field *C.struct_GraphQLAstField, visitorPtr unsafe.Pointer) int {
	v := gopointer.Restore(visitorPtr).(*visitor)
	if err := v.visitor.VisitNonNullType(); err != nil {
		v.err = err
		return 0
	}
	return 1
}

//export visitName
func visitName(field *C.struct_GraphQLAstField, visitorPtr unsafe.Pointer) int {
	v := gopointer.Restore(visitorPtr).(*visitor)
	if err := v.visitor.VisitName(); err != nil {
		v.err = err
		return 0
	}
	return 1
}

//export visitSchemaDefinition
func visitSchemaDefinition(field *C.struct_GraphQLAstField, visitorPtr unsafe.Pointer) int {
	v := gopointer.Restore(visitorPtr).(*visitor)
	if err := v.visitor.VisitSchemaDefinition(); err != nil {
		v.err = err
		return 0
	}
	return 1
}

//export visitOperationType
func visitOperationType(field *C.struct_GraphQLAstField, visitorPtr unsafe.Pointer) int {
	v := gopointer.Restore(visitorPtr).(*visitor)
	if err := v.visitor.VisitOperationType(); err != nil {
		v.err = err
		return 0
	}
	return 1
}

//export visitFieldDefinition
func visitFieldDefinition(field *C.struct_GraphQLAstField, visitorPtr unsafe.Pointer) int {
	v := gopointer.Restore(visitorPtr).(*visitor)
	if err := v.visitor.VisitFieldDefinition(); err != nil {
		v.err = err
		return 0
	}
	return 1
}

//export visitInputValue
func visitInputValue(field *C.struct_GraphQLAstField, visitorPtr unsafe.Pointer) int {
	v := gopointer.Restore(visitorPtr).(*visitor)
	if err := v.visitor.VisitInputValue(); err != nil {
		v.err = err
		return 0
	}
	return 1
}

//export visitInterfaceValueDefinition
func visitInterfaceValueDefinition(field *C.struct_GraphQLAstField, visitorPtr unsafe.Pointer) int {
	v := gopointer.Restore(visitorPtr).(*visitor)
	if err := v.visitor.VisitInterfaceValueDefinition(); err != nil {
		v.err = err
		return 0
	}
	return 1
}

//export visitUnionTypeDefinition
func visitUnionTypeDefinition(field *C.struct_GraphQLAstField, visitorPtr unsafe.Pointer) int {
	v := gopointer.Restore(visitorPtr).(*visitor)
	if err := v.visitor.VisitUnionTypeDefinition(); err != nil {
		v.err = err
		return 0
	}
	return 1
}

//export visitEnumTypeDefinition
func visitEnumTypeDefinition(field *C.struct_GraphQLAstField, visitorPtr unsafe.Pointer) int {
	v := gopointer.Restore(visitorPtr).(*visitor)
	if err := v.visitor.VisitEnumTypeDefinition(); err != nil {
		v.err = err
		return 0
	}
	return 1
}

//export visitEnumValueDefinition
func visitEnumValueDefinition(field *C.struct_GraphQLAstField, visitorPtr unsafe.Pointer) int {
	v := gopointer.Restore(visitorPtr).(*visitor)
	if err := v.visitor.VisitEnumValueDefinition(); err != nil {
		v.err = err
		return 0
	}
	return 1
}

//export visitInputObjectTypeDefinition
func visitInputObjectTypeDefinition(field *C.struct_GraphQLAstField, visitorPtr unsafe.Pointer) int {
	v := gopointer.Restore(visitorPtr).(*visitor)
	if err := v.visitor.VisitInputObjectTypeDefinition(); err != nil {
		v.err = err
		return 0
	}
	return 1
}

//export visitTypeExtensionDefinition
func visitTypeExtensionDefinition(field *C.struct_GraphQLAstField, visitorPtr unsafe.Pointer) int {
	v := gopointer.Restore(visitorPtr).(*visitor)
	if err := v.visitor.VisitTypeExtensionDefinition(); err != nil {
		v.err = err
		return 0
	}
	return 1
}

//export visitDirectiveDefinition
func visitDirectiveDefinition(field *C.struct_GraphQLAstField, visitorPtr unsafe.Pointer) int {
	v := gopointer.Restore(visitorPtr).(*visitor)
	if err := v.visitor.VisitDirectiveDefinition(); err != nil {
		v.err = err
		return 0
	}
	return 1
}
