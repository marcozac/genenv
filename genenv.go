package genenv

type Genenv struct {
	// ConfigFile holds the configuration file path.
	// Overrides any other configuration.
	ConfigFile string `json:"-" yaml:"-"`

	// Package holds the generated package path.
	//   E.g. "github.com/marcozac/genenv/internal/env"
	Package string `json:"package" yaml:"package"`

	// Target holds the generated package directory path, relative
	// to the module root directory.
	//   E.g. "internal/env"
	Target string `json:"target" yaml:"target"`

	Variables map[string]Spec `json:"variables" yaml:"variables"`
}

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
