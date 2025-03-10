package detectany

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "detectany",
	Doc:  "no declaration of \"any\" type",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := func(n ast.Node) bool {
		// 重置类型中的泛型
		if typeSpec, ok := n.(*ast.TypeSpec); ok {
			typeSpec.TypeParams = nil
		}
		// 重置函数泛型
		if funcType, ok := n.(*ast.FuncType); ok {
			funcType.TypeParams = nil
		}

		if ident, ok := n.(*ast.Ident); ok && ident.Name == "any" {
			pass.Reportf(ident.Pos(), "no declaration of \"any\" type")
		}

		return true
	}

	for _, file := range pass.Files {
		ast.Inspect(file, inspect)
	}
	return nil, nil
}

func checkAnyType(t ast.Expr, pass *analysis.Pass) {
	switch v := t.(type) {
	case *ast.Ident:
		if v.Name == "any" {
			pass.Reportf(t.Pos(), "no declaration of \"any\" type")
		}
	case *ast.StarExpr:
		checkAnyType(v.X, pass)
	case *ast.ArrayType:
		checkAnyType(v.Elt, pass)
	case *ast.MapType:
		checkAnyType(v.Key, pass)
		checkAnyType(v.Value, pass)
	case *ast.StructType:
		for _, field := range v.Fields.List {
			if !ignoreAnyLint(field.Doc) {
				checkAnyType(field.Type, pass)
			}
		}
	case *ast.IndexExpr:
		checkAnyType(v.Index, pass)
	case *ast.IndexListExpr:
		for _, indice := range v.Indices {
			checkAnyType(indice, pass)
		}
	case *ast.TypeAssertExpr:
		checkAnyType(v.Type, pass)
	default:
		//
	}
}

func ignoreAnyLint(commentGroup *ast.CommentGroup) bool {
	// pos := fset.Position(node.Pos())
	if commentGroup == nil {
		return false
	}
	for _, comment := range commentGroup.List {
		if strings.Contains(comment.Text, "nolint:detectany") {
			return true
		}
	}
	return false
}
