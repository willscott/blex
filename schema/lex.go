package schema

// Schema represents a lexicon top-level definition
type Schema struct {
	Lexicon     int           `json:"lexicon"`
	ID          string        `json:"id"`
	Revision    *int          `json:"revision,omitempty"`
	Description string        `json:"description,omitempty"`
	Defs        DefinitionMap `json:"defs"`
}

// LexiconPrimaryType represents the common fields found across query, procedure, subscription, and record.
type LexiconPrimaryType struct {
	Type        SchemaType `json:"type"`
	Description string     `json:"description,omitempty"`
}

func (lpt *LexiconPrimaryType) GetType() SchemaType {
	return lpt.Type
}

type DefinitionMap map[string]*Definition

// Definition represents a single lexicon definition
type Definition interface {
	GetType() SchemaType
}

type FieldMap map[string]*Field

// Field represents a single lexicon field
type Field interface {
	GetType() SchemaType
}

type FieldDefinition struct {
	Field
}

// A SchemaType represents a lexicon [Primary Type Definition](https://atproto.com/specs/lexicon#primary-type-definitions)
type SchemaType string

const (
	Query        SchemaType = "query"
	Procedure    SchemaType = "procedure"
	Subscription SchemaType = "subscription"
	Record       SchemaType = "record"

	Parameters SchemaType = "params"

	Reference SchemaType = "ref"
	Union     SchemaType = "union"
	Unknown   SchemaType = "unknown"
	Token     SchemaType = "token"

	// Basic Types
	Object   SchemaType = "object"
	String   SchemaType = "string"
	Array    SchemaType = "array"
	Integer  SchemaType = "integer"
	Float    SchemaType = "float"
	Blob     SchemaType = "blob"
	Boolean  SchemaType = "boolean"
	DateTime SchemaType = "datetime"
	CidLink  SchemaType = "cid-link"
	Bytes    SchemaType = "bytes"
)

func NewDefinition(s SchemaType) Definition {
	switch s {
	case Query:
		return &QueryDefinition{}
	case Procedure:
		return &ProcedureDefinition{}
	case Record:
		return &RecordDefinition{Record: &ObjectDefinition{Properties: make(FieldMap)}}
	case Subscription:
		return &SubscriptionDefinition{}

	case Object:
		return &ObjectDefinition{Properties: make(FieldMap)}
	case Array:
		return &ArrayDefinition{}
	case String:
		return &StringDefinition{}
	case Token:
		return &TokenDefinition{}
	default:
		return nil
	}
}

func NewField(s SchemaType) Field {
	switch s {
	case Object:
		return &ObjectDefinition{Properties: make(FieldMap)}
	case Blob:
		return &BlobDefinition{}
	case Boolean:
		return &BooleanDefinition{}
	case String:
		return &StringDefinition{}
	case Reference:
		return &ReferenceDefinition{}
	case Integer:
		return &IntegerDefinition{}
	case Unknown:
		return &UnknownDefinition{}
	case Array:
		return &ArrayDefinition{}
	case Union:
		return &UnionDefinition{}
	case Token:
		return &TokenDefinition{}
	case CidLink:
		return &CidLinkDefinition{}
	case Bytes:
		return &BytesDefinition{}
	default:
		return nil
	}
}
