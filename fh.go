// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mark-summerfield/fhd"
	"github.com/mark-summerfield/gong"
)

//go:embed Version.dat
var Version string

func main() {
	// TODO subcommands, e.g.,: s|save
	// TODO hidden subcommand: d|dump
	fmt.Printf("fh v%s using fhd v%s\n", strings.TrimSpace(Version),
		strings.TrimSpace(fhd.Version))
	filename := filepath.Join(os.TempDir(), "fh-test.fhd")
	_ = os.Remove(filename)
	fhd, err := fhd.New(filename)
	gong.CheckError("unexpected error", err)
	fmt.Println("filename:", filename)
	fmt.Println("fhd:", fhd)
	fmt.Println("fhd.Dump()")
	_ = fhd.Dump(os.Stdout)
}
