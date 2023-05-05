// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"fmt"
	"os"
)

func doVersion(appName, desc string) {
	fmt.Printf("%s v%s\n", appName, Version[:len(Version)-1])
	os.Exit(0)
}

func doExtract(appName string, args []string, desc string) {
	fmt.Println("doExtract", args, desc)
	os.Exit(0)
}

func doIgnore(appName string, args []string, desc string) {
	fmt.Println("doIgnore", args, desc)
	os.Exit(0)
}

func doList(appName string, args []string, desc string) {
	fmt.Println("doList", args, desc)
	os.Exit(0)
}

func doMonitor(appName string, args []string, desc string) {
	fmt.Println("doMonitor", args, desc)
	os.Exit(0)
}

func doSave(appName string, args []string, desc string) {
	fmt.Println("doSave", args, desc)
	os.Exit(0)
}

func doUnmonitor(appName string, args []string, desc string) {
	fmt.Println("doUnmonitor", args, desc)
	os.Exit(0)
}
