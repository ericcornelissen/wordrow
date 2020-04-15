package main

import "os"

import "github.com/ericcornelissen/wordrow/internal/cli"
import "github.com/ericcornelissen/wordrow/internal/fs"
import "github.com/ericcornelissen/wordrow/internal/logger"
import "github.com/ericcornelissen/wordrow/internal/replacer"
import "github.com/ericcornelissen/wordrow/internal/wordmap"

func run(args cli.Arguments) {
  mapFiles, err := fs.ResolveGlobs(args.MapFiles...)
  wordmap, err := wordmap.WordMapFrom(mapFiles...)
  if err != nil {
    logger.Error(err)
    return
  }

  if args.Invert {
    wordmap.Invert()
  }

  inputFiles, err := fs.ResolveGlobs(args.InputFiles...)
  paths := fs.ResolvePaths(inputFiles...)
  for i := 0; i < len(paths); i++ {
    filePath := paths[i]
    logger.Debugf("Processing '%s'", filePath)

    binaryFileData, err := fs.ReadFile(filePath)
    if err != nil {
      continue
    }

    originalFileData := string(binaryFileData)
    fixedFileData := replacer.ReplaceAll(originalFileData, wordmap)

    if !args.DryRun {
      fs.WriteFile(filePath, fixedFileData)
    } else {
      logger.Printf("Before:\n-------\n%s\n", originalFileData)
      logger.Printf("After:\n------\n%s", fixedFileData)
    }
  }
}

func setLogLevel(args cli.Arguments) {
  if args.Silent {
    logger.SetLogLevel(logger.ERROR)
  } else if args.Verbose {
    logger.SetLogLevel(logger.DEBUG)
  }
}

func main() {
  shouldRun, args := cli.ParseArgs(os.Args)
  if shouldRun {
    setLogLevel(args)
    run(args)
  }
}
