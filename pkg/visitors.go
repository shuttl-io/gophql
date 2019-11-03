package gophql

// VisitField visits fields on the graphql AST
func (w *VisitorWrapper) VisitField(node *FieldNode) error {
	if visitor, ok := w.realVisitor.(FieldVisitor); ok {
		return visitor.VisitField(node)
	}
	return nil
}

// VisitDocument checks to see if the real visitor implements the Document
// visitor and calls it if it does or does nothing
func (w *VisitorWrapper) VisitDocument() error {
	if visitor, ok := w.realVisitor.(DocumentVisitor); ok {
		return visitor.VisitDocument()
	}
	return nil
}

func (w *VisitorWrapper) VisitOperation() error {
	if visitor, ok := w.realVisitor.(OperationVisitor); ok {
		return visitor.VisitOperation()
	}
	return nil
}

func (w *VisitorWrapper) VisitVariableDefinition() error {
	if visitor, ok := w.realVisitor.(VariableDefVisitor); ok {
		return visitor.VisitVariableDefinition()
	}
	return nil
}

func (w *VisitorWrapper) VisitSelectionSet() error {
	if visitor, ok := w.realVisitor.(SelectionSetVisitor); ok {
		return visitor.VisitSelectionSet()
	}
	return nil
}

func (w *VisitorWrapper) VisitArgument() error {
	if visitor, ok := w.realVisitor.(ArgumentVisitor); ok {
		return visitor.VisitArgument()
	}
	return nil
}

func (w *VisitorWrapper) VisitFragmentSpread() error {
	if visitor, ok := w.realVisitor.(FragmentSpreadVisitor); ok {
		return visitor.VisitFragmentSpread()
	}
	return nil
}

func (w *VisitorWrapper) VisitVariable() error {
	if visitor, ok := w.realVisitor.(VariableVisitor); ok {
		return visitor.VisitVariable()
	}
	return nil
}

func (w *VisitorWrapper) VisitInteger() error {
	if visitor, ok := w.realVisitor.(IntegerVisitor); ok {
		return visitor.VisitInteger()
	}
	return nil
}

func (w *VisitorWrapper) VisitFloat() error {
	if visitor, ok := w.realVisitor.(FloatVisitor); ok {
		return visitor.VisitFloat()
	}
	return nil
}

func (w *VisitorWrapper) VisitString() error {
	if visitor, ok := w.realVisitor.(StringVisitor); ok {
		return visitor.VisitString()
	}
	return nil
}

func (w *VisitorWrapper) VisitBoolean() error {
	if visitor, ok := w.realVisitor.(BooleanVisitor); ok {
		return visitor.VisitBoolean()
	}
	return nil
}

func (w *VisitorWrapper) VisitNull() error {
	if visitor, ok := w.realVisitor.(NullVisitor); ok {
		return visitor.VisitNull()
	}
	return nil
}

func (w *VisitorWrapper) VisitEnum() error {
	if visitor, ok := w.realVisitor.(EnumVisitor); ok {
		return visitor.VisitEnum()
	}
	return nil
}

func (w *VisitorWrapper) VisitList() error {
	if visitor, ok := w.realVisitor.(ListVisitor); ok {
		return visitor.VisitList()
	}
	return nil
}

func (w *VisitorWrapper) VisitObject() error {
	if visitor, ok := w.realVisitor.(ObjectVisitor); ok {
		return visitor.VisitObject()
	}
	return nil
}

func (w *VisitorWrapper) VisitObjectField() error {
	if visitor, ok := w.realVisitor.(ObjectFieldVisitor); ok {
		return visitor.VisitObjectField()
	}
	return nil
}

func (w *VisitorWrapper) VisitDirective() error {
	if visitor, ok := w.realVisitor.(DirectiveVisitor); ok {
		return visitor.VisitDirective()
	}
	return nil
}

func (w *VisitorWrapper) VisitNamedType() error {
	if visitor, ok := w.realVisitor.(NamedTypeVisitor); ok {
		return visitor.VisitNamedType()
	}
	return nil
}

func (w *VisitorWrapper) VisitListType() error {
	if visitor, ok := w.realVisitor.(ListTypeVisitor); ok {
		return visitor.VisitListType()
	}
	return nil
}

func (w *VisitorWrapper) VisitNonNullType() error {
	if visitor, ok := w.realVisitor.(NonNullTypeVisitor); ok {
		return visitor.VisitNonNullType()
	}
	return nil
}

func (w *VisitorWrapper) VisitName() error {
	if visitor, ok := w.realVisitor.(NameVisitor); ok {
		return visitor.VisitName()
	}
	return nil
}

func (w *VisitorWrapper) VisitSchemaDefinition() error {
	if visitor, ok := w.realVisitor.(SchemaDefinitionVisitor); ok {
		return visitor.VisitSchemaDefinition()
	}
	return nil
}

func (w *VisitorWrapper) VisitOperationType() error {
	if visitor, ok := w.realVisitor.(OperationTypeVisitor); ok {
		return visitor.VisitOperationType()
	}
	return nil
}

func (w *VisitorWrapper) VisitFieldDefinition() error {
	if visitor, ok := w.realVisitor.(FieldDefinitionVisitor); ok {
		return visitor.VisitFieldDefinition()
	}
	return nil
}

func (w *VisitorWrapper) VisitInputValue() error {
	if visitor, ok := w.realVisitor.(InputValueDefinitionVisitor); ok {
		return visitor.VisitInputValue()
	}
	return nil
}

func (w *VisitorWrapper) VisitInterfaceValueDefinition() error {
	if visitor, ok := w.realVisitor.(InterfaceValueDefinitionVisitor); ok {
		return visitor.VisitInterfaceValueDefinition()
	}
	return nil
}

func (w *VisitorWrapper) VisitUnionTypeDefinition() error {
	if visitor, ok := w.realVisitor.(UnionTypeDefinitionVisitor); ok {
		return visitor.VisitUnionTypeDefinition()
	}
	return nil
}

func (w *VisitorWrapper) VisitEnumTypeDefinition() error {
	if visitor, ok := w.realVisitor.(EnumTypeDefinitionVisitor); ok {
		return visitor.VisitEnumTypeDefinition()
	}
	return nil
}

func (w *VisitorWrapper) VisitEnumValueDefinition() error {
	if visitor, ok := w.realVisitor.(EnumValueDefinitionVisitor); ok {
		return visitor.VisitEnumValueDefinition()
	}
	return nil
}

func (w *VisitorWrapper) VisitInputObjectTypeDefinition() error {
	if visitor, ok := w.realVisitor.(InputObjectTypeDefinitionVisitor); ok {
		return visitor.VisitInputObjectTypeDefinition()
	}
	return nil
}

func (w *VisitorWrapper) VisitTypeExtensionDefinition() error {
	if visitor, ok := w.realVisitor.(TypeExtensionDefinitionVisitor); ok {
		return visitor.VisitTypeExtensionDefinition()
	}
	return nil
}

func (w *VisitorWrapper) VisitDirectiveDefinition() error {
	if visitor, ok := w.realVisitor.(DirectiveDefinitionVisitor); ok {
		return visitor.VisitDirectiveDefinition()
	}
	return nil
}
