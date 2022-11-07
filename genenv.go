package genenv

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

type gen struct {
	*Config
	Spec
	Key, Name string
}

func Generate(cfg *Config) error {
	err := cfg.init()
	if err != nil {
		return err
	}

	g := gen{Config: cfg}
	tmpl := template.Must(template.New("env.tmpl").
		ParseGlob(filepath.Join("templates", "*.tmpl")))

	var f *os.File
	sc := NewStrConv()
	for k, s := range cfg.Variables {
		g.Spec = s
		g.Key = k
		g.Name = sc.ToPascal(g.Key)
		if g.Doc == "" {
			g.Doc = fmt.Sprintf("// %s returns the value of the environment variable %q.", g.Name, g.Key)
		}

		f, err = os.Create(filepath.Join(g.Target, sc.ToLower(g.Name)+".go"))
		if err != nil {
			return err
		}
		defer f.Close()

		err = tmpl.Execute(f, g)
		if err != nil {
			return err
		}
	}

	return nil
}
