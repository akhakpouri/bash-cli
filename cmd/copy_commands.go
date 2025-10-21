package cmd

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const (
	red    = "\033[0;31m"
	green  = "\033[0;32m"
	yellow = "\033[0;33m"
	none   = "\033[0m" // Reset to default color
)

func RecusiveCopy(dir string, source string, destination string) {
	found := false
	err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && d.Name() == source {
			found = true
			parent := filepath.Dir(path)
			dest := filepath.Join(parent, destination)
			err := copy(path, dest)
			if err != nil {
				fmt.Printf("%s[ERROR]%s: %s", red, none, err)
			}
		}
		return nil
	})
	if !found {
		fmt.Printf("%s[WARN]: %s Filed were not found.\n", yellow, none)
	}

	if err != nil {
		panic(err)
	}
}

func copy(src, dst string) error {
	if _, err := os.Stat(dst); err == nil {
		fmt.Printf("%s [WARN]: %s %s alredy exists.\n\n", yellow, none, dst)
	} else if errors.Is(err, os.ErrNotExist) {
		fmt.Printf("%s[INFO]: %s Copying %s to %s\n", green, none, src, dst)
		source, err := os.Open(src)

		if err != nil {
			return printError("Failed to open the source file", err)
		}
		defer source.Close()

		destination, err := os.Create(dst)
		if err != nil {
			return printError("Failed to open the destination file", err)
		}
		defer destination.Close()

		_, err = io.Copy(destination, source)
		if err != nil {
			return printError("Failed to copy file content", err)
		}

		sourceInfo, err := source.Stat()
		if err != nil {
			return fmt.Errorf("failed to get source file's info %w\n\n", err)
		}

		err = os.Chmod(dst, sourceInfo.Mode())
		if err != nil {
			return printError("Failed to set destination file permissions", err)
		}
	} else {
		panic(err)
	}

	return nil
}

func printError(msg string, err error) error {
	return fmt.Errorf("%s[ERR]: %s %s. %w\n\n", red, none, msg, err)
}
