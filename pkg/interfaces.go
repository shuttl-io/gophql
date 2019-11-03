package gophql

// FieldVisitor visits a graphql field node
type FieldVisitor interface {
	VisitField(*FieldNode) error
}

// DocumentVisitor Visits the document nodes
type DocumentVisitor interface {
	VisitDocument() error
}

// OperationVisitor Visits the operation nodes
type OperationVisitor interface {
	VisitOperation() error
}

// VariableDefVisitor visits the variable definition nodes
type VariableDefVisitor interface {
	VisitVariableDefinition() error
}

//SelectionSetVisitor visits selection set nodes
type SelectionSetVisitor interface {
	VisitSelectionSet() error
}

//ArgumentVisitor visits argument nodes
type ArgumentVisitor interface {
	VisitArgument() error
}

// FragmentSpreadVisitor visits fragment spreads
type FragmentSpreadVisitor interface {
	VisitFragmentSpread() error
}

//VariableVisitor visits variable nodes
type VariableVisitor interface {
	VisitVariable() error
}

type IntegerVisitor interface {
	VisitInteger() error
}

type FloatVisitor interface {
	VisitFloat() error
}

type StringVisitor interface {
	VisitString() error
}

type BooleanVisitor interface {
	VisitBoolean() error
}

type NullVisitor interface {
	VisitNull() error
}

type EnumVisitor interface {
	VisitEnum() error
}

type ListVisitor interface {
	VisitList() error
}

type ObjectVisitor interface {
	VisitObject() error
}

type ObjectFieldVisitor interface {
	VisitObjectField() error
}

type DirectiveVisitor interface {
	VisitDirective() error
}

type NamedTypeVisitor interface {
	VisitNamedType() error
}

type ListTypeVisitor interface {
	VisitListType() error
}

type NonNullTypeVisitor interface {
	VisitNonNullType() error
}

type NameVisitor interface {
	VisitName() error
}

type SchemaDefinitionVisitor interface {
	VisitSchemaDefinition() error
}

type OperationTypeVisitor interface {
	VisitOperationType() error
}

type FieldDefinitionVisitor interface {
	VisitFieldDefinition() error
}

type InputValueDefinitionVisitor interface {
	VisitInputValue() error
}

type InterfaceValueDefinitionVisitor interface {
	VisitInterfaceValueDefinition() error
}

type UnionTypeDefinitionVisitor interface {
	VisitUnionTypeDefinition() error
}

type EnumTypeDefinitionVisitor interface {
	VisitEnumTypeDefinition() error
}

type EnumValueDefinitionVisitor interface {
	VisitEnumValueDefinition() error
}

type InputObjectTypeDefinitionVisitor interface {
	VisitInputObjectTypeDefinition() error
}

type TypeExtensionDefinitionVisitor interface {
	VisitTypeExtensionDefinition() error
}

type DirectiveDefinitionVisitor interface {
	VisitDirectiveDefinition() error
}

// Visitor is an interface that is composed of all of the other visitors.
type Visitor interface {
	FieldVisitor
	DocumentVisitor
	OperationVisitor
	VariableDefVisitor
	SelectionSetVisitor
	ArgumentVisitor
	FragmentSpreadVisitor
	VariableVisitor
	IntegerVisitor
	FloatVisitor
	StringVisitor
	BooleanVisitor
	NullVisitor
	EnumVisitor
	ListVisitor
	ObjectVisitor
	ObjectFieldVisitor
	DirectiveVisitor
	NamedTypeVisitor
	ListTypeVisitor
	NonNullTypeVisitor
	NameVisitor
	SchemaDefinitionVisitor
	OperationTypeVisitor
	FieldDefinitionVisitor
	InputValueDefinitionVisitor
	InterfaceValueDefinitionVisitor
	UnionTypeDefinitionVisitor
	EnumTypeDefinitionVisitor
	EnumValueDefinitionVisitor
	InputObjectTypeDefinitionVisitor
	TypeExtensionDefinitionVisitor
	DirectiveDefinitionVisitor
}
