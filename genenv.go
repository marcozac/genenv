package genenv

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Spec struct {
	// Allow holds a list of allowed variable values.
	// No entries means "any value".
	Allow []string `json:"allow" yaml:"allow"`

	// Deny holds a list of denied variable values.
	// No entries means "nothing".
	Deny []string `json:"deny" yaml:"deny"`

	// @TODO
	// Convert value to type.
	// Type string
}

type Genenv struct {
	// Package holds the name of the generated package.
	//
	// NOTE
	// Package value is not parsed and must ben not a path.
	Package string `json:"package" yaml:"package"`

	// Target holds the generated package directory path. Relative paths
	// must be relative to the module root directory.
	Target string `json:"target" yaml:"target"`

	Variables map[string]Spec `json:"variables" yaml:"variables"`

	mod *ModInfo
}

func (g *Genenv) init() error {
	if len(g.Variables) == 0 {
		return ErrNothing
	}

	var err error
	if g.mod == nil {
		g.mod, err = GetModInfo()
		if err != nil {
			return fmt.Errorf("getting module info: %w", err)
		}
	}

	switch {
	case g.Target == "" && g.Package == "":
		g.Package = "env"
		fallthrough // assign Target

	case g.Package != "" && g.Target == "": // no target
		g.Target = filepath.Join(g.mod.Dir, "internal", g.Package)

	case g.Target != "" && !filepath.IsAbs(g.Target): // resolve to root
		g.Target = filepath.Join(g.mod.Dir, g.Target)
	}

	var name string
	var found bool
	name, found, err = TargetPackage(g.Target)
	if err != nil {
		return fmt.Errorf("getting target package: %w", err)
	}

	if g.Package == "" {
		g.Package = name
		return nil
	}

	if name != g.Package && found {
		return fmt.Errorf("inconsistent package name %s: found %s in %s", g.Package, name, g.Target)
	}

	return nil
}

func ReadConfig(p string) (*Genenv, error) {
	f, err := os.Open(p)
	if err != nil {
		return nil, fmt.Errorf("opening %s: %w", p, err)
	}
	defer f.Close()

	dec := yaml.NewDecoder(f)
	dec.KnownFields(true)

	var v Genenv
	err = dec.Decode(&v)
	if err != nil {
		return nil, fmt.Errorf("decoding %s: %w", p, err)
	}

	return &v, nil
}
