// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	_ "embed"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/mark-summerfield/fhd"
	"github.com/mark-summerfield/gong"
)

//go:embed Version.dat
var Version string

// See ~/app/go/clip/eg/subcommands/subcommands.go

func main() {
	// TODO subcommands, e.g.,: s|save
	// TODO hidden subcommand: d|dump
	appName := strings.TrimSuffix(path.Base(os.Args[0]), ".exe")
	args := os.Args[1:]
	if len(args) == 0 || args[0] == "-h" || args[0] == "--help" ||
		args[0] == "help" {
		if len(args) == 2 { // Adds support for: -h subcommand
			args[0], args[1] = args[1], args[0]
		} else {
			showHelp() // doesn't return
		}
	}
	subcmd := args[0]
	args = args[1:]
	switch subcmd {
	case "-v", "--version", "v", "version":
		fmt.Printf("%s v%s\n", appName, Version[:len(Version)-1])
		os.Exit(0)
	case "s", "save":
		//return parseCompare(appName, args, descs[0])
		return
	case "d", "dump":
		dump() // does not return
	}
	fmt.Printf("error: invalid subcommand %q: use -h or --help\n",
		subcmd)
	os.Exit(2)
}

func dump() {
	filename := filepath.Join(os.TempDir(), "fh-test.fhd")
	_ = os.Remove(filename)
	fhd, err := fhd.New(filename)
	gong.CheckError("unexpected error", err)
	fmt.Println("filename:", filename)
	fmt.Println("fhd:", fhd)
	fmt.Println("fhd.Dump()")
	_ = fhd.Dump(os.Stdout)
	os.Exit(0)
}

func showHelp() {
	fmt.Println("showHelp TODO")
	os.Exit(0)
}
