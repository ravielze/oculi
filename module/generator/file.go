package generator

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"

	u "github.com/ravielze/oculi/module/generator/utils"
)

func ReadDeclaredFunctions(folderName, fileName string) []string {
	fset := token.NewFileSet()
	fileData := u.ReadFile(folderName, fileName)

	node, err := parser.ParseFile(fset, fmt.Sprintf("%s/%s", folderName, fileName), nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	var result []string
	for _, f := range node.Decls {
		fn, ok := f.(*ast.FuncDecl)
		if !ok {
			continue
		}

		params := ""
		for _, p := range fn.Type.Params.List {
			if len(params) > 0 {
				params += ", "
			}
			params += fmt.Sprintf("%s %s", p.Names[0].Name, fileData[p.Type.Pos()-1:p.Type.End()-1])
		}

		results := ""
		if fn.Type.Results != nil {
			for _, r := range fn.Type.Results.List {
				if len(results) > 0 {
					results += ", "
				}
				results += fmt.Sprintf("%s", r.Type)
			}
		}
		fmt.Println(fn.Name.Name, "|", params, "|", results)
		result = append(result, fn.Name.Name)
	}
	return result
}

func ReadDeclaredInterfaces(folderName, fileName string) map[string][]string {
	fset := token.NewFileSet()
	fileData := u.ReadFile(folderName, fileName)

	node, err := parser.ParseFile(fset, fmt.Sprintf("%s/%s", folderName, fileName), nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	result := map[string][]string{}
	for _, f := range node.Decls {
		fn, okf := f.(*ast.GenDecl)
		if !okf {
			continue
		}
		if fn.Tok != 84 {
			continue
		}
		for _, ent := range fn.Specs {
			ty, okty := ent.(*ast.TypeSpec)
			if !okty {
				continue
			}

			itf, okitf := ty.Type.(*ast.InterfaceType)
			if !okitf {
				continue
			}

			var methods []string
			for _, m := range itf.Methods.List {
				methods = append(methods, fileData[m.Pos()-1:m.End()-1])
				a, okok := m.Type.(*ast.FuncType)
				if !okok {
					continue
				}

				params := ""
				for _, p := range a.Params.List {
					if len(params) > 0 {
						params += ", "
					}
					params += fmt.Sprintf("%s %s", p.Names[0].Name, fileData[p.Type.Pos()-1:p.Type.End()-1])
				}

				results := ""
				if a.Results != nil {
					for _, r := range a.Results.List {
						if len(results) > 0 {
							results += ", "
						}
						results += fmt.Sprintf("%s", r.Type)
					}
				}
				fmt.Println(m.Names[0].Name, "|", params, "|", results)
				//fmt.Println(NewMethod(m.Names[0].Name, fileData[m.Pos()-1:m.End()-1]))
			}
			result[ty.Name.Name] = methods
		}
	}

	return result
}
