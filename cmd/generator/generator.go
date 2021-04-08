package generator

import (
	"embed"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

//go:embed templates/*
var templates embed.FS

type Language int

const (
	Typescript = iota
	Go
	Python
	Csharp
)

func (l Language) String() string {
	return [...]string{"typescript", "go", "python", "csharp"}[l]
}

type Generator struct {
	language    Language
	packageName string
	dir         string
}

func NewGenerator(lang string, name string, dir string) (*Generator, error) {
	var l Language
	switch lang {
	case "typescript":
		l = Typescript
	case "go":
		l = Go
	case "csharp":
		return nil, errors.New("csharp is not yet supported")
	case "python":
		return nil, errors.New("python is not yet supported")
	default:
		return nil, errors.New(fmt.Sprintf("unknown language input: %s", lang))
	}
	g := &Generator{
		language:    l,
		packageName: name,
		dir:         dir,
	}
	return g, nil
}

func (g *Generator) Generate() error {
	dest := filepath.Join(g.dir, g.packageName)
	err := os.Mkdir(dest, 0755)
	if err != nil {
		return err
	}
	var templateDir string
	switch g.language {
	case Typescript:
		templateDir = "typescript"
	case Go:
		templateDir = "go"
	default:
		return fmt.Errorf("unsupported language: %s", g.language)
	}

	templateRoot := filepath.Join(".", "templates", templateDir)

	templateDirEntries, err := templates.ReadDir(templateRoot)
	if err != nil {
		return err
	}

	for _, entry := range templateDirEntries {
		err = copyAndReplace(entry, dest, templateRoot, "", g.packageName)
		if err != nil {
			fmt.Println("there was an error")
			return err
		}
	}

	return nil
}

func copyAndReplace(entry os.DirEntry, destRoot, templateRoot, subpath, packageName string) error {
	subpath = filepath.Join(subpath, entry.Name())
	dest := filepath.Join(destRoot, subpath)
	dest = strings.ReplaceAll(dest, "xyz", packageName)
	// work around https://github.com/golang/go/issues/45197
	dest = strings.ReplaceAll(dest, "not_a_go.mod", "go.mod")
	src := filepath.Join(templateRoot, subpath)

	if entry.IsDir() {
		err := os.Mkdir(dest, 0755)
		if err != nil {
			return err
		}
		entries, err := templates.ReadDir(src)
		if err != nil {
			return err
		}
		for _, e := range entries {
			err = copyAndReplace(e, destRoot, templateRoot, subpath, packageName)
			if err != nil {
				return err
			}
		}
	} else {
		b, err := templates.ReadFile(src)
		contents := string(b)
		contents = strings.ReplaceAll(contents, "xyz", packageName)

		if err != nil {
			return err
		}
		err = ioutil.WriteFile(dest, []byte(contents), 0755)
		if err != nil {
			return err
		}
	}

	return nil
}
