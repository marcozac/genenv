package genenv

import (
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
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
		return nil, err
	}

	err = cmd.Start()
	if err != nil {
		return nil, fmt.Errorf("starting cmd: %w", err)
	}

	var v ModInfo
	err = json.NewDecoder(stdout).Decode(&v)
	if err != nil {
		return nil, fmt.Errorf("unmarshaling mod info: %w", err)
	}

	err = cmd.Wait()
	if err != nil {
		return nil, fmt.Errorf("cmd exited with error: %w", err)
	}

	if v.GoMod == "" {
		return nil, errors.New("running out of a module directory")
	}

	return &v, nil
}
