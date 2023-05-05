// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	_ "embed"
	"fmt"
	"os"
	"path"
	"strings"
)

//go:embed Version.dat
var Version string

func main() {
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
	switch subcmd { // None of these return
	case "e", extractName:
		doExtract(appName, args, descForSubcmd[extractName])
	case "i", ignoreName:
		doIgnore(appName, args, descForSubcmd[ignoreName])
	case "l", listName:
		doList(appName, args, descForSubcmd[listName])
	case "m", monitorName:
		doMonitor(appName, args, descForSubcmd[monitorName])
	case "s", saveName:
		doSave(appName, args, descForSubcmd[saveName])
	case "u", unmonitorName:
		doUnmonitor(appName, args, descForSubcmd[unmonitorName])
	case "-v", "--version", versionName:
		doVersion(appName, descForSubcmd[versionName])
	default:
		fmt.Printf("error: invalid subcommand %q: use -h or --help\n",
			subcmd)
		os.Exit(2)
	}
}
