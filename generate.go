package GoGRF

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"text/template"
)

func Generate(definition *GrfDefinition) error {
	if err := copyResources(); err != nil {
		return err
	}
	return generateToFile(&definition.Cargos[0], "cargo", "cargo")
}

func copyResources() error {
	if err := copyDirectory("resources", "generated"); err != nil {
		return err
	}
	return nil
}

func generateToFile(data interface{}, template string, name string) error {
	template = template + ".go.tmpl"
	name = name + ".nml"

	tmplFile, err := templatePath(template)
	if err != nil {
		return err
	}

	tmpl, err := parseTemplate(template, tmplFile)
	if err != nil {
		return err
	}

	if err := ensureDirectory("generated"); err != nil {
		return err
	}

	if err := createFile("generated", name, tmpl, data); err != nil {
		return err
	}

	return nil
}

func templatePath(template string) (string, error) {
	filePath, err := filePath()
	if err != nil {
		return "", err // propagate error directly
	}
	return filepath.Join(filePath, "templates", template), nil
}

func parseTemplate(name string, tmplFile string) (*template.Template, error) {
	tmpl, err := template.New(name).ParseFiles(tmplFile)
	if err != nil {
		return nil, fmt.Errorf("error parsing template %s: %w", tmplFile, err)
	}
	return tmpl, nil
}

func createFile(dir, filename string, tmpl *template.Template, data interface{}) error {
	file, err := os.Create(filepath.Join(dir, filename))
	if err != nil {
		return fmt.Errorf("error creating file %s: %w", filename, err)
	}
	defer file.Close()

	if err := tmpl.Execute(file, data); err != nil {
		return fmt.Errorf("error executing template: %w", err)
	}

	return nil
}

func filePath() (string, error) {
	pc := make([]uintptr, 1)
	n := runtime.Callers(1, pc)
	if n == 0 {
		return "", fmt.Errorf("unable to get caller information")
	}

	fn := runtime.FuncForPC(pc[0])
	if fn == nil {
		return "", fmt.Errorf("unable to resolve function information")
	}

	file, _ := fn.FileLine(pc[0])

	absolutePath, err := filepath.Abs(filepath.Dir(file))
	if err != nil {
		return "", fmt.Errorf("error converting to absolute path: %v", err)
	}

	return absolutePath, nil
}
