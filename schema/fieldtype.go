package schema

// LexiconFieldType represents a field
type LexiconFieldType struct {
	Type        SchemaType `json:"type"`
	Description string     `json:"description,omitempty"`
}

func (lpt *LexiconFieldType) GetType() SchemaType {
	return lpt.Type
}

type ObjectDefinition struct {
	LexiconFieldType

	Properties FieldMap
	Required   *[]string
	Nullable   *[]string
}

type ArrayDefinition struct {
	LexiconFieldType

	Items     FieldDefinition
	MinLength *int
	MaxLength *int
}

type UnionDefinition struct {
	LexiconFieldType

	Refs   []string
	Closed *bool
}

type BooleanDefinition struct {
	LexiconFieldType

	Default *bool
	Const   *bool
}

type StringDefinition struct {
	LexiconFieldType

	Format       *string
	MaxLength    *int
	MinLength    *int
	MaxGraphemes *int
	MinGraphemes *int
	KnownValues  *[]string
	Enum         *[]string
	Default      *string
	Const        *string
}

type ReferenceDefinition struct {
	LexiconFieldType

	Ref string
}

type IntegerDefinition struct {
	LexiconFieldType

	Minimum *int
	Maximum *int
	Enum    *[]int
	Default *int
	Const   *int
}

type BlobDefinition struct {
	LexiconFieldType

	Accept  *[]string
	MaxSize *int
}

type BytesDefinition struct {
	LexiconFieldType

	MaxLength *int
	MinLength *int
}

type UnknownDefinition struct {
	LexiconFieldType
}

type TokenDefinition struct {
	LexiconFieldType
}

type CidLinkDefinition struct {
	LexiconFieldType
}
