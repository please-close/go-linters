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
		// 检测结构体类型声明中的 any
		if typeSpec, ok := n.(*ast.TypeSpec); ok {
			// if strcutType, ok := typeSpec.Type.(*ast.StructType); ok {
			// 	for _, field := range strcutType.Fields.List {
			// 		if ident, ok := field.Type.(*ast.Ident); ok && ident.Name == "any" {
			// 			if ignoreAnyLint(field.Doc) {
			// 				continue
			// 			}
			// 			pass.Reportf(ident.Pos(), "no declaration of \"any\" type")
			// 		}
			// 	}
			// }
			checkAnyType(typeSpec.Type, pass)
		}
		// 检测函数参数/返回值中的 any
		if funcDecl, ok := n.(*ast.FuncDecl); ok {
			if ignoreAnyLint(funcDecl.Doc) {
				return true
			}
			if funcDecl.Type.Params != nil {
				for _, field := range funcDecl.Type.Params.List {
					// if ident, ok := field.Type.(*ast.Ident); ok && ident.Name == "any" {
					// 	pass.Reportf(ident.Pos(), "no declaration of \"any\" type")
					// }
					checkAnyType(field.Type, pass)
				}
			}
			if funcDecl.Type.Results != nil {
				for _, field := range funcDecl.Type.Results.List {
					// if ident, ok := field.Type.(*ast.Ident); ok && ident.Name == "any" {
					// 	pass.Reportf(ident.Pos(), "no declaration of \"any\" type")
					// }
					checkAnyType(field.Type, pass)
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

func checkAnyType(t ast.Expr, pass *analysis.Pass) {
	switch v := t.(type) {
	case *ast.Ident:
		if v.Name == "any" {
			if !ignoreAnyLint(v.Obj.Decl.(*ast.TypeSpec).Doc) {
				pass.Reportf(t.Pos(), "no declaration of \"any\" type")
			}
		}
	case *ast.StarExpr:
		checkAnyType(v.X, pass)
	case *ast.ArrayType:
		checkAnyType(v.Elt, pass)
	case *ast.MapType:
		checkAnyType(v.Key, pass)
		checkAnyType(v.Value, pass)
	case *ast.ChanType:
		checkAnyType(v.Value, pass)
	case *ast.StructType:
		for _, field := range v.Fields.List {
			checkAnyType(field.Type, pass)
		}
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
