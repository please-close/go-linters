package main

import (
	"github.com/please-close/go-linters/linters/detectany"

	"golang.org/x/tools/go/analysis"
)

// nolint:detectany
func New(conf any) ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{detectany.Analyzer}, nil
}
