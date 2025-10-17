package cmd

import (
	"testing"
)

func TestRec(t *testing.T) {
	dir := "~/code/go/bash-cli/tmp"
	src := "foo.txt"
	dst := "bar.txt"

	RecusiveCopy(dir, src, dst)
}
