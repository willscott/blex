package main

import (
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
	}

	fmt.Printf("accumulated %d schemas\r\n", len(schemas))

	return nil
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
