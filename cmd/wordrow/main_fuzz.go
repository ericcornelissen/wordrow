// +build gofuzz

package main

import (
	"errors"

	"github.com/ericcornelissen/stringsx"
	"github.com/ericcornelissen/wordrow/internal/cli"
	"github.com/ericcornelissen/wordrow/internal/wordmaps"
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

func _processMapFile(s, format string, wordmap *wordmaps.WordMap) {
	s = stringsx.ReplaceAll(s, ";", "\n")
	mapfileReader := stringsx.NewReader(s)
	processMapFile(mapfileReader, format, wordmap)
}

func _doReplace(s string, wordmap *wordmaps.WordMap) string {
	s = stringsx.ReplaceAll(s, ";", "\n")
	inputfileReader := stringsx.NewReader(s)
	output, _ := doReplace(inputfileReader, wordmap)
	return output
}

func Fuzz(data []byte) int {
	inputs, err := dataToInputs(data)
	if err != nil {
		return -1
	}

	rawArgs := stringsx.Split(inputs[0], ";")
	_, args := cli.ParseArgs(rawArgs)

	var wordmap wordmaps.WordMap
	processInlineMappings(args.Mappings, &wordmap)
	_processMapFile(inputs[1], csv, &wordmap)
	_processMapFile(inputs[2], markdown, &wordmap)

	if args.Invert {
		wordmap.Invert()
	}

	output := _doReplace(inputs[3], &wordmap)
	if output != inputs[3] {
		return 1
	}

	return 0
}
