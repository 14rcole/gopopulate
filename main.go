package main

import (
  "os"
)

var makeDirs, makeRegFiles bool

func PopulateDir(baseDir, types string, depth int = 0, maxFiles int = 4) error {
  err := parseTypes
  if err != nil {
    return err
  }

  err = popDirHelper(baseDir, depth, maxFiles)
  if err != nil {
    return err
  }

  return nil
}

func popDirHelper(baseDir string, depth, maxFiles int) error {
  var numFiles int = 0
  for numFiles < maxFiles; numFiles++ {
    type := pickType()
    newFile, err := genFile()
    if err != nil {
      return err
    }
    if type == 1 && depth > 0 {
      popDirHelper(strings.Concat(baseDir, newFile), depth - 1, maxFiles)
    }
  }
  return nil
}

func genDataForFile() string {
}

func parseTypes() error {
}

func pickType() string {
  if makeDirs && makeRegFiles {
  } else if makeRegFiles {
    return 0
  } else {
    return 1
  }
}

func genFile(type int) error {
  switch type {
    case 0:
      // Create a file and call genDataForFile to write to it
    case 1:
      os.Mkdir
  }
}
