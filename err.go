package genenv

import "errors"

// ErrNoVar is the error reported when there is nothing to
// generate.
var ErrNoVar = errors.New("nothing to generate")

// ErrModDir is the error reported trying to get Go module info
// out of a module directory.
var ErrModDir = errors.New("out of module directory")
