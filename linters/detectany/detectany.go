package detectany

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "detectany",
	Doc:  "no declaration of \"any\" type",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := func(n ast.Node) bool {
		// 检测结构体类型声明中的 any
		if typeSpec, ok := n.(*ast.TypeSpec); ok {
			if strcutType, ok := typeSpec.Type.(*ast.StructType); ok {
				for _, field := range strcutType.Fields.List {
					if ident, ok := field.Type.(*ast.Ident); ok && ident.Name == "any" {
						pass.Reportf(ident.Pos(), "no declaration of \"any\" type")
					}
				}
			}
		}
		// 检测函数参数/返回值中的 any
		if funcDecl, ok := n.(*ast.FuncDecl); ok {
			for _, field := range funcDecl.Type.Params.List {
				if ident, ok := field.Type.(*ast.Ident); ok && ident.Name == "any" {
					pass.Reportf(ident.Pos(), "no declaration of \"any\" type")
				}
			}
			if funcDecl.Type.Results != nil {
				for _, field := range funcDecl.Type.Results.List {
					if ident, ok := field.Type.(*ast.Ident); ok && ident.Name == "any" {
						pass.Reportf(ident.Pos(), "no declaration of \"any\" type")
					}
				}
			}
		}

		return true
	}

	for _, file := range pass.Files {
		ast.Inspect(file, inspect)
	}
	return nil, nil
}
