// Command exhaustive is a command line interface for the exhaustive
// package at github.com/strongdm/exhaustive.
//
// # Usage
//
//	exhaustive [flags] [packages]
package main

import (
	"github.com/strongdm/exhaustive"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(exhaustive.Analyzer) }
