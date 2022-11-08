package genenv

import (
	"os"
	"path/filepath"
)

func (suite *GenenvTestSuite) TestTargetPackage() {
	// Package exists (aliased).
	name, found, err := TargetPackage(filepath.Join("testdata", "target", "ok"))
	suite.NoError(err)
	suite.True(found)
	suite.Equal(name, "okpkg")

	// Package exists (with errors).
	name, found, err = TargetPackage(filepath.Join("testdata", "target", "err"))
	suite.Error(err)
	suite.True(found)
	suite.Zero(name)

	// Package directory exists (no *.go files).
	name, found, err = TargetPackage(filepath.Join("testdata", "target", "nogo"))
	suite.NoError(err)
	suite.False(found)
	suite.Equal(name, "nogo")

	// Package directory does not exist.
	name, found, err = TargetPackage(filepath.Join(suite.d, "mkdir"))
	suite.NoError(err)
	suite.False(found)
	suite.Equal(name, "mkdir")

	// Not a directory error.
	var f *os.File
	f, err = os.CreateTemp(suite.d, "*")
	suite.Require().NoError(err)
	f.Close()
	name, found, err = TargetPackage(f.Name())
	suite.Error(err)
	suite.False(found)
	suite.Zero(name)

	// Stat error.
	_, _, err = TargetPackage("\x00")
	suite.Error(err)

	// Mkdir error.
	var ed string
	ed, err = os.MkdirTemp(suite.d, "*")
	suite.NoError(err)
	suite.NoError(os.Chmod(ed, 0o555))
	_, _, err = TargetPackage(filepath.Join(ed, "dir"))
	suite.Error(err)
}
