package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
)

func main() {
	root := "./" // потом поменяем

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// пропускаем директории
		if info.IsDir() {
			return nil
		}

		ext := filepath.Ext(path)

		switch ext {

		case ".go":
			fmt.Println("[GO]   ", path)
			parseGoFile(path)

		case ".ts", ".tsx", ".js":
			fmt.Println("[JS]   ", path)

		case ".vue":
			fmt.Println("[VUE]  ", path)

		case ".py":
			fmt.Println("[PY]   ", path)

		case ".rb":
			fmt.Println("[RUBY] ", path)
		}

		return nil
	})

	if err != nil {
		panic(err)
	}
}

func parseGoFile(path string) {
	fset := token.NewFileSet()

	node, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		fmt.Println("parse error:", err)
		return
	}

	for _, decl := range node.Decls {

		fn, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}

		fmt.Println("  → FUNC:", fn.Name.Name)
	}
}
