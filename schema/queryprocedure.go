package schema

// QueryDefinition represents an atproto 'query'
type QueryDefinition struct {
	LexiconPrimaryType
	Parameters *ParametersDefinition `json:"parameters,omitempty"`
	Output     OutputDefinition      `json:"output,omitempty"`
	Errors     []Error               `json:"errors,omitempty"`
}

// ProcedureDefinition represents an atproto 'procedure'
type ProcedureDefinition struct {
	LexiconPrimaryType
	Parameters *ParametersDefinition `json:"parameters,omitempty"`
	Output     *OutputDefinition     `json:"output,omitempty"`
	Input      *InputDefintion       `json:"input,omitempty"`
	Errors     []Error               `json:"errors,omitempty"`
}

type OutputDefinition struct {
	Description string           `json:"description,omitempty"`
	Encoding    string           `json:"encoding"`
	Schema      *FieldDefinition `json:"schema,omitempty"`
}

type InputDefintion struct {
	Description string           `json:"description,omitempty"`
	Encoding    string           `json:"encoding"`
	Schema      *FieldDefinition `json:"schema,omitempty"`
}

type ParametersDefinition struct {
	LexiconFieldType

	Required   *[]string `json:"required,omitempty"`
	Properties FieldMap  `json:"properties"`
}

type Error struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}
