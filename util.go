package genenv

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"

	"gopkg.in/yaml.v3"
)

type ModInfo struct {
	// Path holds the go module path.
	//   E.g. "github.com/marcozac/genenv"
	Path string `json:"Path"`

	// Dir holds the go module directory.
	//   E.g. "/my/module/directory"
	Dir string `json:"Dir"`

	// GoMod holds the go.mod file path.
	//   E.g. "/my/module/directory/go.mod"
	GoMod string `json:"GoMod"`
}

func GetModInfo() (*ModInfo, error) {
	cmd := exec.Command("go", "list", "-m", "-json")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("piping cmd stdout: %w", err)
	}

	err = cmd.Start()
	if err != nil {
		return nil, fmt.Errorf("starting cmd: %w", err)
	}

	var v ModInfo
	err = json.NewDecoder(stdout).Decode(&v)
	if err != nil {
		return nil, fmt.Errorf("decoding: %w", err)
	}

	err = cmd.Wait()
	if err != nil {
		return nil, fmt.Errorf("cmd exited with error: %w", err)
	}

	if v.GoMod == "" {
		return nil, errors.New("out of module directory")
	}

	return &v, nil
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
		return nil, fmt.Errorf("decoding: %w", err)
	}

	return &v, nil
}
