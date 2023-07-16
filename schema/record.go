package schema

type RecordDefinition struct {
	LexiconPrimaryType

	Key    string
	Record *ObjectDefinition
}
