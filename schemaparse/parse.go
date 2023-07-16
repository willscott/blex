package schemaparse

import (
	"encoding/json"
	"io"
	"os"

	"github.com/willscott/blex/schema"
)

func ParseLexiconFromFile(name string) (*schema.Schema, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return ParseLexicon(file)
}

func ParseLexicon(stream io.Reader) (*schema.Schema, error) {
	decoder := json.NewDecoder(stream)

	schema := schema.Schema{Defs: make(schema.DefinitionMap)}
	err := decoder.Decode(&schema)
	if err != nil {
		return nil, err
	}
	return &schema, nil
}
