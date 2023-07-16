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

	Properties FieldMap  `json:"properties"`
	Required   *[]string `json:"required,omitempty"`
	Nullable   *[]string `json:"nullable,omitempty"`
}

type ArrayDefinition struct {
	LexiconFieldType

	Items     FieldDefinition `json:"items"`
	MinLength *int            `json:"minLength,omitempty"`
	MaxLength *int            `json:"maxLength,omitempty"`
}

type UnionDefinition struct {
	LexiconFieldType

	Refs   []string `json:"refs"`
	Closed *bool    `json:"closed,omitempty"`
}

type BooleanDefinition struct {
	LexiconFieldType

	Default *bool `json:"default,omitempty"`
	Const   *bool `json:"const,omitempty"`
}

type StringDefinition struct {
	LexiconFieldType

	Format       *string   `json:"format,omitempty"`
	MaxLength    *int      `json:"maxLength,omitempty"`
	MinLength    *int      `json:"minLength,omitempty"`
	MaxGraphemes *int      `json:"maxGraphemes,omitempty"`
	MinGraphemes *int      `json:"minGraphemes,omitempty"`
	KnownValues  *[]string `json:"knownValues,omitempty"`
	Enum         *[]string `json:"enum,omitempty"`
	Default      *string   `json:"default,omitempty"`
	Const        *string   `json:"const,omitempty"`
}

type ReferenceDefinition struct {
	LexiconFieldType

	Ref string `json:"ref"`
}

type IntegerDefinition struct {
	LexiconFieldType

	Minimum *int   `json:"minimum,omitempty"`
	Maximum *int   `json:"maximum,omitempty"`
	Enum    *[]int `json:"enum,omitempty"`
	Default *int   `json:"default,omitempty"`
	Const   *int   `json:"const,omitempty"`
}

type BlobDefinition struct {
	LexiconFieldType

	Accept  *[]string `json:"accept,omitempty"`
	MaxSize *int      `json:"maxSize,omitempty"`
}

type BytesDefinition struct {
	LexiconFieldType

	MaxLength *int `json:"maxLength,omitempty"`
	MinLength *int `json:"minLength,omitempty"`
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
