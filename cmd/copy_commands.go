package cmd

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
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
				fmt.Printf("exception occured. %s", err)
			}
		}
		return nil
	})
	if !found {
		fmt.Printf("no files with name %s were found.\n", source)
	}

	if err != nil {
		panic(err)
	}
}

func copy(src, dst string) error {
	if _, err := os.Stat(dst); err == nil {
		return fmt.Errorf("file exists %s. canceling copy.\n\n", dst)
	} else if errors.Is(err, os.ErrNotExist) {
		fmt.Printf("copying %s to %s\n", src, dst)
		source, err := os.Open(src)

		if err != nil {
			return fmt.Errorf("failed to open the source file %w\n\n", err)
		}
		defer source.Close()

		destination, err := os.Create(dst)
		if err != nil {
			return fmt.Errorf("failed to open the destination file %w\n\n", err)
		}
		defer destination.Close()

		_, err = io.Copy(destination, source)
		if err != nil {
			return fmt.Errorf("failed to copy file content %w\n\n", err)
		}

		sourceInfo, err := source.Stat()
		if err != nil {
			return fmt.Errorf("failed to get source file's info %w\n\n", err)
		}

		err = os.Chmod(dst, sourceInfo.Mode())
		if err != nil {
			return fmt.Errorf("failed to set destination file permissions: %w\n\n", err)
		}
	} else {
		panic(err)
	}

	return nil
}
