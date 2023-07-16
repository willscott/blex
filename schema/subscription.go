package schema

type SubscriptionDefinition struct {
	LexiconPrimaryType

	Parameters *ParametersDefinition `json:"parameters,omitempty"`
	Message    MessageDefinition     `json:"message"`
	Errors     []Error               `json:"errors"`
}

type MessageDefinition struct {
	Description string          `json:"description,omitempty"`
	Schema      UnionDefinition `json:"schema"`
}
