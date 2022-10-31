package genenv

import (
	"fmt"
	"go/build"
	"io/fs"
	"os"
	"path/filepath"
)

// TargetPackage is an utility function to get the package name from its target
// directory.
//
// TargetPackage checks d existence, reporting a [*fs.PathError] on failure.
//
// If d does not exist, it is created and its base path is assumed as package name.
// The d base path is assumed as package name also if d exists but does not
// contain any *.go file.
//
// If d exists and contains any *.go file, found is reported as true and returns
// the package name or the error occurred by importing it.
func TargetPackage(d string) (name string, found bool, err error) {
	var info fs.FileInfo
	if info, err = os.Stat(d); err != nil {
		if !os.IsNotExist(err) {
			return "", false, &fs.PathError{
				Op:   "target directory info",
				Path: d,
				Err:  err,
			}
		}

		err = os.MkdirAll(d, 0o755)
		if err != nil {
			return "", false, fmt.Errorf("creating target directory: %w", err)
		}
		return filepath.Base(d), false, nil

	} else if !info.IsDir() {
		return "", false, fmt.Errorf("%s is not a directory", d)
	}

	var pkg *build.Package
	pkg, err = build.Default.ImportDir(d, 0)
	if err != nil {
		if _, ok := err.(*build.NoGoError); ok {
			return filepath.Base(d), false, nil
		}
		return "", true, fmt.Errorf("importing package: %w", err)
	}

	return pkg.Name, true, nil
}
