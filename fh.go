// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"
)

var AppName = strings.TrimSuffix(path.Base(os.Args[0]), ".exe")

func main() {
	log.SetFlags(0)
	args := os.Args[1:]
	maybeShowHelp(args) // May not return
	subcmd := args[0]
	args = args[1:]
	switch subcmd { // None of these return
	case "e", extractName:
		doExtract(args, descForSubcmd[extractName])
	case "d", "dump":
		doDump()
	case "i", ignoreName:
		doIgnore(args, descForSubcmd[ignoreName])
	case "l", listName:
		doList(args, descForSubcmd[listName])
	case "m", monitorName:
		doMonitor(args, descForSubcmd[monitorName])
	case "s", saveName:
		doSave(args, descForSubcmd[saveName])
	case "u", unmonitorName:
		doUnmonitor(args, descForSubcmd[unmonitorName])
	case "-v", "--version", versionName:
		doVersion(descForSubcmd[versionName])
	default:
		fmt.Printf("error: invalid subcommand %q: use -h or help\n", subcmd)
		os.Exit(2)
	}
}

func maybeShowHelp(args []string) {
	if len(args) == 0 || args[0] == "-h" || args[0] == "--help" ||
		args[0] == "help" {
		if len(args) == 2 { // Adds support for: -h subcommand
			args[0], args[1] = args[1], args[0]
		} else {
			showHelp() // doesn't return
		}
	}
}
