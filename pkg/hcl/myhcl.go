package hcl

import (
	"bytes"
	"log"

	"github.com/hashicorp/hcl/hcl/printer"
	jsonParser "github.com/hashicorp/hcl/json/parser"
	"github.com/tmccombs/hcl2json/convert"
)

func Json2bytes(json *string) (*bytes.Buffer, error) {
	ast, err := jsonParser.Parse([]byte(*json))
	if err != nil {
		log.Printf("error json parse: %s", err)
		return nil, err
	}

	stdin := bytes.NewBufferString("")
	err = printer.Fprint(stdin, ast)
	if err != nil {
		log.Printf("error stdin print: %s", err)
		return nil, err
	}

	return stdin, err
}

func Json2hcl(json *string) (*string, error) {
	hclBytes, err := Json2bytes(json)
	if err != nil {
		return nil, err
	}
	hcl := hclBytes.String()

	return &hcl, nil
}

func Hcl2json(hcl *string) (string, error) {
	var options convert.Options
	b, err := convert.Bytes([]byte(*hcl), "", options)
	if err != nil {
		return "", err
	}

	return string(b), nil

}
