package main

import "os"

import "github.com/ericcornelissen/wordrow/internal/cli"
import "github.com/ericcornelissen/wordrow/internal/fs"
import "github.com/ericcornelissen/wordrow/internal/logger"
import "github.com/ericcornelissen/wordrow/internal/replacer"
import "github.com/ericcornelissen/wordrow/internal/wordmap"


func getWordMap(mapFilesPaths []string) (wm wordmap.WordMap, err error) {
  logger.Debug("Reading specified mapping files")
  mapFiles, err := fs.ReadFiles(mapFilesPaths)
  if err != nil {
    return wm, err
  }

  for _, mapFile := range mapFiles {
    logger.Debugf("Processing '%s' as mapping file", mapFile.Path)
    err := wm.AddFile(mapFile)
    if err != nil {
      return wm, err
    }
  }

  return wm, err
}

func run(args cli.Arguments) error {
  wm, err := getWordMap(args.MapFiles)
  if err != nil {
    return err
  }

  if args.Invert {
    wm.Invert()
  }

  inputFiles, err := fs.ReadFiles(args.InputFiles)
  if err != nil {
    return err
  }

  for _, file := range inputFiles {
    logger.Debugf("Processing '%s' as input file", file.Path)
    fixedFileData := replacer.ReplaceAll(file.Content, wm)

    if !args.DryRun {
      fs.WriteFile(file.Path, fixedFileData)
    } else {
      logger.Printf("Before:\n-------\n%s\n", file.Content)
      logger.Printf("After:\n------\n%s", fixedFileData)
    }
  }

  return nil
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

    err := run(args)
    if err != nil {
      logger.Error(err)
    }
  }
}
