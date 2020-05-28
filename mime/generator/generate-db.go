package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/function61/gokit/ezhttp"
	"github.com/function61/gokit/mime"
)

//go:generate go run generate-db.go

func main() {
	if err := logic(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func logic() error {
	resp := map[string]mime.Spec{}
	if _, err := ezhttp.Get(
		context.Background(),
		"https://raw.githubusercontent.com/jshttp/mime-db/master/db.json",
		ezhttp.RespondsJson(&resp, false),
	); err != nil {
		return err
	}

	goCode := []string{}
	code := func(line string) {
		goCode = append(goCode, line)
	}

	code("package mime")

	// for boolean pointers (not w/o hacks: https://stackoverflow.com/a/28818489)
	code("var trueVal = true")
	code("var falseVal = false")

	code("")

	code("// source https://github.com/jshttp/mime-db")
	code("var mimeTypes = map[string]*Spec{")

	for contentType, spec := range resp {
		if len(spec.Extensions) == 0 {
			continue // we're only interested in content-types that have file extensions defined
		}

		code(fmt.Sprintf(`"%s": &Spec{`, contentType))

		if spec.Source != "" {
			code(fmt.Sprintf(`Source: "%s",`, spec.Source))
		}

		if spec.CharEncoding != "" {
			code(fmt.Sprintf(`CharEncoding: "%s",`, spec.CharEncoding))
		}

		if spec.Compressible != nil {
			if *spec.Compressible {
				code("Compressible: &trueVal,")
			} else {
				code("Compressible: &falseVal,")
			}
		}

		extensionsAsCode := ""

		for _, ext := range spec.Extensions {
			extensionsAsCode = extensionsAsCode + fmt.Sprintf(`"%s",`, ext)
		}

		code(fmt.Sprintf("Extensions: []string{%s},", extensionsAsCode))

		code("},")
	}

	code("}")

	generated := strings.Join(goCode, "\n")

	return ioutil.WriteFile("../db.go", []byte(generated), 0755)
}
