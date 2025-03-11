package plugin

import (
	"github.com/golangci/plugin-module-register/register"
	"github.com/please-close/go-linters/linters/detectany"
	"golang.org/x/tools/go/analysis"
)

func init() {
	register.Plugin("detectany", New)
}

type Plugin struct {
}

func New(settings any) (register.LinterPlugin, error) {
	return &Plugin{}, nil
}

func (p *Plugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{detectany.Analyzer}, nil
}

func (f *Plugin) GetLoadMode() string {
	return register.LoadModeSyntax
}
