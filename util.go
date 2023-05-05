// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mark-summerfield/fhd"
	"github.com/mark-summerfield/gong"
)

func getFhd() *fhd.Fhd {
	entries, err := os.ReadDir(".")
	gong.CheckError("", err)
	files := make([]string, 0)
	for _, entry := range entries {
		if strings.ToUpper(filepath.Ext(entry.Name())) == ".FHD" {
			files = append(files, entry.Name())
		}
	}
	switch {
	case len(files) == 1:
		fh, err := fhd.New(files[0])
		gong.CheckError("", err)
		return fh
	case len(files) > 1:
		return chooseFhd(files)
	}
	fmt.Println("error: failed to find .fhd file", files)
	os.Exit(1)
	return nil
}

func chooseFhd(files []string) *fhd.Fhd {
	// TODO show list to user and ask them to choose, e.g.,
	// 1. abc.fhd
	// 2. def.fhd
	// Choose [1-2 or q for quit]:
	fmt.Println("chooseFhd unimplemented", files)
	os.Exit(1)
	return nil
}
