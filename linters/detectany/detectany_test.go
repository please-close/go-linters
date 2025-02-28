package detectany_test

import (
	"testing"

	"github.com/please-close/go-linters/linters/detectany"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, detectany.Analyzer, "a")
}
