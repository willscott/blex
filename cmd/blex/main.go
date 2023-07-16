package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/urfave/cli/v2"
	"github.com/willscott/blex/schema"
	"github.com/willscott/blex/schemaparse"

	cjson "github.com/docker/go/canonical/json"
)

func main() {
	app := &cli.App{
		Name:  "blex",
		Usage: "Work with lexicons",
		Commands: []*cli.Command{
			{
				Name:   "parse",
				Action: parseLex,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func parseLex(ctx *cli.Context) error {
	args := ctx.Args()
	if args.Len() == 0 {
		return errors.New("usage: parse <source schema>")
	}

	// get relevant files within the source directory tree.
	sources, err := expandArgs(args.Slice())
	if err != nil {
		return err
	}

	schemas := []*schema.Schema{}
	for _, source := range sources {
		schema, err := schemaparse.ParseLexiconFromFile(source)
		if err != nil {
			return fmt.Errorf("could not parse %s: %w", source, err)
		}
		schemas = append(schemas, schema)

		// ensure schema is fully parsed.
		parsedJSON, err := json.MarshalIndent(schema, "", "  ")
		if err != nil {
			return fmt.Errorf("representation of %s is not serializable: %w", source, err)
		}
		originalData, err := os.ReadFile(source)
		if err == nil {
			if !bytes.Equal(canonicalRepresentation(parsedJSON), canonicalRepresentation(originalData)) {
				fmt.Printf("==parsed==\n%s\n==orig==\n%s\n", parsedJSON, originalData)
				return fmt.Errorf("parse of %s was not complete", source)
			}
		}
	}

	fmt.Printf("found %d schemas\r\n", len(schemas))

	return nil
}

func canonicalRepresentation(jsonData []byte) []byte {
	var repr interface{}
	if err := json.Unmarshal(jsonData, &repr); err != nil {
		return nil
	}
	out, err := cjson.MarshalCanonical(repr)
	if err != nil {
		return nil
	}
	return out
}

func findSchemas(dir string) ([]string, error) {
	var out []string
	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if strings.HasSuffix(path, ".json") {
			out = append(out, path)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return out, nil

}

func expandArgs(args []string) ([]string, error) {
	var out []string
	for _, a := range args {
		st, err := os.Stat(a)
		if err != nil {
			return nil, err
		}
		if st.IsDir() {
			s, err := findSchemas(a)
			if err != nil {
				return nil, err
			}
			out = append(out, s...)
		} else if strings.HasSuffix(a, ".json") {
			out = append(out, a)
		}
	}

	return out, nil
}
