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

type Config struct {
	// Package holds the name of the generated package.
	//
	// NOTE
	// Package value is not parsed and must ben not a path.
	Package string `json:"package" yaml:"package"`

	// Target holds the generated package directory path. Relative paths
	// must be relative to the module root directory.
	Target string `json:"target" yaml:"target"`

	// Header holds the generated files header.
	//
	// Default:
	//   // Code generated by genenv. DO NOT EDIT.
	Header string

	Variables map[string]Spec `json:"variables" yaml:"variables"`

	mod *ModInfo
}

func (c *Config) init() error {
	if len(c.Variables) == 0 {
		return ErrNoVar
	}

	if c.Header == "" {
		c.Header = "// Code generated by genenv. DO NOT EDIT."
	}

	var err error
	if c.mod == nil {
		c.mod, err = GetModInfo()
		if err != nil {
			return fmt.Errorf("getting module info: %w", err)
		}
	}

	switch {
	case c.Target == "" && c.Package == "":
		c.Package = "env"
		fallthrough // assign Target

	case c.Package != "" && c.Target == "": // no target
		c.Target = filepath.Join(c.mod.Dir, "internal", c.Package)

	case c.Target != "" && !filepath.IsAbs(c.Target): // resolve to root
		c.Target = filepath.Join(c.mod.Dir, c.Target)
	}

	var name string
	var found bool
	name, found, err = TargetPackage(c.Target)
	if err != nil {
		return fmt.Errorf("getting target package: %w", err)
	}

	if c.Package == "" {
		c.Package = name
		return nil
	}

	if name != c.Package && found {
		return fmt.Errorf("inconsistent package name %s: found %s in %s", c.Package, name, c.Target)
	}

	return nil
}

func ReadConfig(p string) (*Config, error) {
	f, err := os.Open(p)
	if err != nil {
		return nil, fmt.Errorf("opening %s: %w", p, err)
	}
	defer f.Close()

	dec := yaml.NewDecoder(f)
	dec.KnownFields(true)

	var cfg Config
	err = dec.Decode(&cfg)
	if err != nil {
		return nil, fmt.Errorf("decoding %s: %w", p, err)
	}

	return &cfg, nil
}