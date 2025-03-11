package cmd

import (
	"github.com/please-close/go-linters/linters/detectany"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() {
	unitchecker.Main(detectany.Analyzer)
}
