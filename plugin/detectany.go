package main

import (
	"go-linters/linters/detectany"

	"golang.org/x/tools/go/analysis"
)

// golint:ignore:detectany
func New(conf any) ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{detectany.Analyzer}, nil
}
