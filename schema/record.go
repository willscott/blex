package schema

type RecordDefinition struct {
	LexiconPrimaryType

	Key    string            `json:"key"`
	Record *ObjectDefinition `json:"record,omitempty"`
}
