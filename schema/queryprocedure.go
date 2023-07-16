package schema

// QueryDefinition represents an atproto 'query'
type QueryDefinition struct {
	LexiconPrimaryType
	Parameters ParametersDefinition
	Output     OutputDefinition
	Errors     []Error
}

// ProcedureDefinition represents an atproto 'procedure'
type ProcedureDefinition struct {
	LexiconPrimaryType
	Parameters ParametersDefinition
	Output     OutputDefinition
	Input      InputDefintion
	Errors     []Error
}

type OutputDefinition struct {
	LexiconFieldType

	Encoding string
	Schema   FieldDefinition
}

type InputDefintion struct {
	LexiconFieldType

	Encoding string
	Schema   FieldDefinition
}

type ParametersDefinition struct {
	LexiconFieldType
}

type Error struct {
	Name        string
	Description string
}
