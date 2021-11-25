package main

import (
	"github.com/Rodge0/importas"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(importas.Analyzer)
}
