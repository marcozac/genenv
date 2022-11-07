package genenv

import (
	"os"
	"path/filepath"
	"text/template"
)

type gen struct {
	*Config
	Spec
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
	for _, s := range cfg.Variables {
		g.Spec = s

		f, err = os.Create(filepath.Join(g.Target, cfg.sc.ToLower(s.Name)+".go"))
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
