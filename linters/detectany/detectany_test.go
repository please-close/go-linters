package detectany_test

import (
	"testing"

	"go-linters/linters/detectany"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, detectany.Analyzer, "a")
}
