package genenv

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type ModInfo struct {
	// Path holds the go module path.
	// E.g. "github.com/marcozac/genenv"
	Path string `json:"Path"`

	// Dir holds the go module directory.
	// E.g. "/my/module/directory"
	Dir string `json:"Dir"`

	// GoMod holds the go.mod file path.
	// E.g. "/my/module/directory/go.mod"
	GoMod string `json:"GoMod"`
}

// GetModInfo runs "go list -m -json" and returns a [*ModInfo] from its output.
// If there is an error and the command is running out of a module directory,
// the error reported is [ErrModDir].
//
//	info, err := GetModInfo()
//	info = &ModInfo{
//		Path:  "github.com/marcozac/genenv",
//		Dir:   "/my/module/directory",
//		GoMod: "/my/module/directory/go.mod",
//	}
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
		return nil, ErrModDir
	}

	return &v, nil
}
