// +build gofuzz

package main

import (
	"errors"

	"github.com/ericcornelissen/wordrow/internal/cli"
	"github.com/ericcornelissen/wordrow/internal/strings"
	"github.com/ericcornelissen/wordrow/internal/wordmaps"
)

const (
	csv      = "csv"
	markdown = "md"
)

func dataToInputs(data []byte) ([]string, error) {
	s := string(data)
	inputs := strings.Split(s, "\n")

	if len(inputs) < 4 {
		return nil, errors.New("too little input data")
	}

	return inputs, nil
}

func _processMapFile(s, format string, wordmap *wordmaps.WordMap) {
	s = strings.ReplaceAll(s, ";", "\n")
	mapfileReader := strings.NewReader(s)
	processMapFile(mapfileReader, format, wordmap)
}

func _doReplace(s string, wordmap *wordmaps.WordMap) {
	inputfileReader := strings.NewReader(s)
	doReplace(inputfileReader, wordmap)
}

func Fuzz(data []byte) int {
	x, err := dataToInputs(data)
	if err != nil {
		return -1
	}

	_, args := cli.ParseArgs(strings.Split(x[0], ";"))

	var wordmap wordmaps.WordMap
	processInlineMappings(args.Mappings, &wordmap)
	_processMapFile(x[1], csv, &wordmap)
	_processMapFile(x[2], markdown, &wordmap)

	if args.Invert {
		wordmap.Invert()
	}

	_doReplace(x[3], &wordmap)

	if wordmap.Size() > 0 && len(x[3]) > 0 {
		return 1
	}

	return 0
}
