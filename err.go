package genenv

import "errors"

// ErrNothing is the error reported when there is nothing to
// generate.
var ErrNothing = errors.New("nothing to generage")

// ErrModDir is the error reported trying to get Go module info
// out of a module directory.
var ErrModDir = errors.New("out of module directory")
