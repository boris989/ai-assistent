package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"

	"github.com/boris989/ai-assistent/internal/indexer"
)

func main() {
	root := "./test_project" // путь к тестовому проекту

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
			fmt.Println("[GO]", path)

			fileChunks := parseGoFile(path)

			for _, c := range fileChunks {
				fmt.Println("  →", c.Type, ":", c.Name, ":", c.Content)
			}
		}

		return nil
	})

	if err != nil {
		panic(err)
	}
}

func parseGoFile(path string) []indexer.Chunk {
	fset := token.NewFileSet()

	node, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		fmt.Println("parse error:", err)
		return nil
	}

	var chunks []indexer.Chunk

	for _, decl := range node.Decls {
		switch decl := decl.(type) {

		case *ast.FuncDecl:
			start := fset.Position(decl.Pos()).Offset
			end := fset.Position(decl.End()).Offset

			content, err := extractCode(path, start, end)
			if err != nil {
				continue
			}

			chunks = append(chunks, indexer.Chunk{
				ID:       indexer.GenerateID(content),
				FilePath: path,
				Language: indexer.LanguageGo,
				Name:     decl.Name.Name,
				Type:     indexer.ChunkTypeFunction,
				Content:  content,
			})

		case *ast.GenDecl:
			for _, spec := range decl.Specs {
				typeSpec, ok := spec.(*ast.TypeSpec)
				if !ok {
					continue
				}

				start := fset.Position(typeSpec.Pos()).Offset
				end := fset.Position(typeSpec.End()).Offset

				content, err := extractCode(path, start, end)
				if err != nil {
					continue
				}

				switch typeSpec.Type.(type) {

				case *ast.StructType:
					chunks = append(chunks, indexer.Chunk{
						ID:       indexer.GenerateID(content),
						FilePath: path,
						Language: indexer.LanguageGo,
						Name:     typeSpec.Name.Name,
						Type:     indexer.ChunkTypeStruct,
						Content:  content,
					})

				case *ast.InterfaceType:
					chunks = append(chunks, indexer.Chunk{
						ID:       indexer.GenerateID(content),
						FilePath: path,
						Language: indexer.LanguageGo,
						Name:     typeSpec.Name.Name,
						Type:     indexer.ChunkTypeInterface,
						Content:  content,
					})
				}
			}
		}
	}

	return chunks
}

func extractCode(path string, start, end int) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	if start >= len(data) || end > len(data) {
		return "", fmt.Errorf("invalid range")
	}

	return string(data[start:end]), nil
}
