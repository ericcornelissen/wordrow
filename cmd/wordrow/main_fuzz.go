// +build gofuzz

package main

import (
	"errors"

	"github.com/ericcornelissen/stringsx"
	"github.com/ericcornelissen/wordrow/internal/cli"
	"github.com/ericcornelissen/wordrow/internal/common"
)

const (
	csv      = "csv"
	markdown = "md"
)

func dataToInputs(data []byte) ([]string, error) {
	s := string(data)

	inputs := stringsx.Split(s, "\n")
	if len(inputs) < 4 {
		return nil, errors.New("too little input data")
	}

	return inputs, nil
}

func _processMapFile(s, format string, mapping map[string]string) {
	s = stringsx.ReplaceAll(s, ";", "\n")
	mapfileReader := stringsx.NewReader(s)
	newMapping, err := processMapFile(mapfileReader, format)
	if err == nil {
		common.MergeMaps(mapping, newMapping)
	}
}

func _doReplace(s string, mapping map[string]string) string {
	s = stringsx.ReplaceAll(s, ";", "\n")
	inputfileReader := stringsx.NewReader(s)
	output, _ := doReplace(inputfileReader, mapping)
	return output
}

func Fuzz(data []byte) int {
	inputs, err := dataToInputs(data)
	if err != nil {
		return -1
	}

	rawArgs := stringsx.Split(inputs[0], ";")
	_, args := cli.ParseArgs(rawArgs)

	mapping := make(map[string]string)
	forEach(args.Mappings, processInlineMappingWith(mapping))
	_processMapFile(inputs[1], csv, mapping)
	_processMapFile(inputs[2], markdown, mapping)

	if args.Invert {
		mapping = invert(mapping)
	}

	output := _doReplace(inputs[3], mapping)
	if output != inputs[3] {
		return 1
	}

	return 0
}
