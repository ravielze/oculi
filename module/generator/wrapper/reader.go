package generator_wrapper

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"sort"

	u "github.com/ravielze/fuzzy-broccoli/module/generator/utils"
)

func GetMethodWrapper(folderName string) (InterfaceWrapper, InterfaceWrapper, InterfaceWrapper) {
	controllerImplemented := readDeclaredFunctions(folderName, "controller.go")
	usecaseImplemented := readDeclaredFunctions(folderName, "usecase.go")
	repositoryImplemented := readDeclaredFunctions(folderName, "repository.go")

	allDeclared := readDeclaredInterfaces(folderName, "entity.go")
	controllerDeclared := allDeclared["IController"]
	usecaseDeclared := allDeclared["IUsecase"]
	repositoryDeclared := allDeclared["IRepo"]

	return InterfaceWrapper{
			Implemented: controllerImplemented,
			Declared:    controllerDeclared,
		}, InterfaceWrapper{
			Implemented: usecaseImplemented,
			Declared:    usecaseDeclared,
		}, InterfaceWrapper{
			Implemented: repositoryImplemented,
			Declared:    repositoryDeclared,
		}
}

func readDeclaredFunctions(folderName, fileName string) []MethodWrapper {
	fset := token.NewFileSet()
	fileData := u.ReadFile(folderName, fileName)

	node, err := parser.ParseFile(fset, fmt.Sprintf("%s/%s", folderName, fileName), nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	var result []MethodWrapper
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
		result = append(result, MethodWrapper{
			Name:      fn.Name.Name,
			Return:    results,
			Parameter: params,
			Body:      fileData[fn.Body.Pos()+3 : fn.Body.End()-4],
		})
	}
	sort.Sort(MethodWrapperSorter(result))
	return result
}

func readDeclaredInterfaces(folderName, fileName string) map[string][]MethodWrapper {
	fset := token.NewFileSet()
	fileData := u.ReadFile(folderName, fileName)

	node, err := parser.ParseFile(fset, fmt.Sprintf("%s/%s", folderName, fileName), nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	result := map[string][]MethodWrapper{}
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

			var methods []MethodWrapper
			for _, m := range itf.Methods.List {
				methodName := m.Names[0].Name
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
				methods = append(methods, MethodWrapper{
					Name:      methodName,
					Return:    results,
					Parameter: params,
					Body:      "",
				})
			}
			sort.Sort(MethodWrapperSorter(methods))
			result[ty.Name.Name] = methods
		}
	}

	return result
}
