// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"fmt"
	"os"
)

// See ~/app/go/clip/eg/subcommands/subcommands.go

func doVersion(desc string) {
	fmt.Printf("%s v%s\n", AppName, Version[:len(Version)-1])
	os.Exit(0)
}

func doDump() {
	fh := getFhd()
	fmt.Println("dump", fh)
	_ = fh.Dump(os.Stdout)
	os.Exit(0)
}

func doExtract(args []string, desc string) {
	fmt.Println("doExtract", args, desc)
	os.Exit(0)
}

func doIgnore(args []string, desc string) {
	fmt.Println("doIgnore", args, desc)
	os.Exit(0)
}

func doList(args []string, desc string) {
	fmt.Println("doList", args, desc)
	os.Exit(0)
}

func doMonitor(args []string, desc string) {
	fmt.Println("doMonitor", args, desc)
	os.Exit(0)
}

func doSave(args []string, desc string) {
	fmt.Println("doSave", args, desc)
	os.Exit(0)
}

func doUnmonitor(args []string, desc string) {
	fmt.Println("doUnmonitor", args, desc)
	os.Exit(0)
}
