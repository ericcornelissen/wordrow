package main

import "os"

import "github.com/ericcornelissen/wordrow/internal/cli"
import "github.com/ericcornelissen/wordrow/internal/fs"
import "github.com/ericcornelissen/wordrow/internal/logger"
import "github.com/ericcornelissen/wordrow/internal/replacer"
import "github.com/ericcornelissen/wordrow/internal/wordmap"


func run(args cli.Arguments) {
  var wm wordmap.WordMap

  mapFiles, err := fs.ReadFiles(args.MapFiles)
  if err != nil {
    logger.Error(err)
    return
  }

  for _, mapFile := range mapFiles {
    err := wm.AddFile(mapFile)
    if err != nil {
      logger.Error(err)
      return
    }
  }

  if args.Invert {
    wm.Invert()
  }

  inputFiles, err := fs.ReadFiles(args.InputFiles)
  if err != nil {
    logger.Error(err)
    return
  }

  for _, file := range inputFiles {
    logger.Debugf("Processing '%s'", file.Path)
    fixedFileData := replacer.ReplaceAll(file.Content, wm)

    if !args.DryRun {
      fs.WriteFile(file.Path, fixedFileData)
    } else {
      logger.Printf("Before:\n-------\n%s\n", file.Content)
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
