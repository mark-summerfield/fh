// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"fmt"
	"os"
	"path"
	"strings"
	"unicode/utf8"

	"github.com/mark-summerfield/clip"
)

func showHelp() {
	fmt.Printf("usage: %s [SUBCOMMAND] ...\n\n%s\n",
		clip.Strong(strings.TrimSuffix(path.Base(os.Args[0]), ".exe")),
		clip.Emph("subcommands:"))
	subs := []string{
		clip.Strong("e") + ", " + clip.Strong(extractName) +
			" [-s SID] <FILE>",
		clip.Strong("i") + ", " + clip.Strong(ignoreName) +
			" <FILE1> [FILE2 ...]",
		clip.Strong("l") + ", " + clip.Strong(listName),
		clip.Strong("m") + ", " + clip.Strong(monitorName) +
			" <FILE1> [FILE2 ...]",
		clip.Strong("s") + ", " + clip.Strong(saveName),
		clip.Strong("u") + ", " + clip.Strong(unmonitorName) +
			" <FILE1> [FILE2 ...]",
		clip.Strong("-h") + ", " + clip.Strong(helpName),
		clip.Strong("-v") + ", " + clip.Strong(versionName)}
	leftWidths := make([]int, 0, len(subs))
	descWidths := make([]int, 0, len(descForSubcmd))
	maxLeft := 0
	maxDesc := 0
	descs := descsSlice()
	for i, sub := range subs {
		left := utf8.RuneCountInString(sub)
		if left > maxLeft {
			maxLeft = left
		}
		leftWidths = append(leftWidths, left)
		left = utf8.RuneCountInString(descs[i])
		if left > maxDesc {
			maxDesc = left
		}
		descWidths = append(descWidths, left)
	}
	argWidth := maxLeft
	width := clip.GetWidth()
	for i := 0; i < len(subs); i++ {
		if leftWidths[i]+descWidths[i]+4 <= width {
			fmt.Printf("  %s  %s\n", subs[i], descs[i])
		} else {
			fmt.Printf("  %s  ", subs[i])
			fmt.Print(clip.ArgHelp(argWidth, width, descs[i]))
		}
	}
	os.Exit(0)
}

func descsSlice() []string {
	descs := make([]string, 0, len(descForSubcmd))
	for _, key := range []string{extractName, ignoreName, listName,
		monitorName, saveName, unmonitorName, helpName, versionName} {
		descs = append(descs, descForSubcmd[key])
	}
	return descs
}
