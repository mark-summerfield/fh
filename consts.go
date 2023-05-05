// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import _ "embed"

//go:embed Version.dat
var Version string

const (
	extractName   = "extract"
	ignoreName    = "ignore"
	listName      = "list"
	monitorName   = "monitor"
	saveName      = "save"
	unmonitorName = "unmonitor"
	helpName      = "help"
	versionName   = "version"
)

var descForSubcmd = map[string]string{
	extractName: "Extract the latest save of the given filename or the " +
		"save specified by the -s --sid.",
	ignoreName: "Ignore the specified file(s). If any are already " +
		"monitored they will be set to be unmonitored instead.",
	listName:    "List all the available save IDs (SIDs).",
	monitorName: "Monitor the specified file(s).",
	saveName: "Save all the monitored files that have changed since the " +
		"last save.",
	unmonitorName: "Unmonitor the specified file(s).",
	helpName:      "Show this help text and quit.",
	versionName:   "Show version and quit.",
}
