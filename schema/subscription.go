package schema

type SubscriptionDefinition struct {
	LexiconPrimaryType

	Parameters ParametersDefinition
	Message    MessageDefinition
	Errors     []Error
}

type MessageDefinition struct {
	LexiconFieldType

	Schema UnionDefinition
}
