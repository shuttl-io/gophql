package graphql

/*
#include <stdio.h>
struct GraphQLAstField;
int visitField(struct GraphQLAstField *field, void *visitorPtr);
int visit_field_cgo(struct GraphQLAstField *field, void *visitorPtr) {
  return visitField(field, visitorPtr);
}
*/
import "C"
