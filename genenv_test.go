package genenv

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestSuite(t *testing.T) {
	suite.Run(t, &GenenvTestSuite{})
}

type GenenvTestSuite struct {
	suite.Suite
	modInfo *ModInfo
	f       string // configuration file path
	d       string // tempdir path
}

func (suite *GenenvTestSuite) SetupSuite() {
	var err error
	suite.modInfo, err = GetModInfo()
	if err != nil {
		suite.FailNow("getting module info", err)
	}

	suite.f = filepath.Join("testdata", ".genenv.yml")
	suite.d = suite.T().TempDir()
}

// resets g Target and Package
func (suite *GenenvTestSuite) resetTP(g *Genenv) {
	g.Target = ""
	g.Package = ""
}

func (suite *GenenvTestSuite) TestGenenvInit() {
	d, err := os.MkdirTemp(suite.d, "*")
	suite.Require().NoError(err)
	defer os.RemoveAll(d)

	g := &Genenv{}
	suite.ErrorIs(g.init(), ErrNothing)

	// Cover module loading and mock it.
	g.Variables = map[string]Spec{"foo": {}}
	g.Package = "okpkg"
	g.Target = filepath.Join("testdata", "target", "ok")
	suite.NoError(g.init())
	g.mod = &ModInfo{
		Path: "github.com/marcozac/fakemod",
		Dir:  d,
	}

	// No Package or Target.
	suite.resetTP(g)
	target := filepath.Join(g.mod.Dir, "internal", "env")
	suite.NoDirExists(target)
	suite.NoError(g.init())
	suite.DirExists(target)
	suite.Equal(target, g.Target)
	suite.Equal("env", g.Package)

	// No Target => assigned from package name.
	suite.resetTP(g)
	g.Package = "foo"
	target = filepath.Join(filepath.Dir(target), "foo")
	suite.NoDirExists(target)
	suite.NoError(g.init())
	suite.DirExists(target)
	suite.Equal(target, g.Target)
	suite.Equal("foo", g.Package)

	// No Package => assigned from target base.
	g.Package = ""
	suite.NoError(g.init())
	suite.Equal(target, g.Target)
	suite.Equal("foo", g.Package)

	// Aliasing error.
	//
	// g.Package: foo
	// g.Target: <temp>/internal/foo
	// package in *.go: bar
	var f *os.File
	f, err = os.CreateTemp(target, "*.go")
	suite.Require().NoError(err)
	_, err = f.WriteString("package bar\n")
	f.Close()
	suite.Require().NoError(err)
	suite.ErrorContains(g.init(), "inconsistent package name")

	// Aliasing ok.
	// Same target above.
	// g.Target: <temp>/internal/foo
	g.Package = "bar"
	suite.NoError(g.init())
	suite.Equal(target, g.Target)
	suite.Equal("bar", g.Package)

	// No Package => assigned from target (aliased).
	// Same target above.
	// g.Target: <temp>/internal/foo
	g.Package = ""
	suite.NoError(g.init())
	suite.Equal("bar", g.Package)

	// Mkdir Target (aliased).
	g.Package = "hello"
	g.Target = filepath.Join(d, "world")
	suite.NoDirExists(g.Target)
	suite.NoError(g.init())
	suite.DirExists(g.Target)
	suite.Equal("hello", g.Package)

	// Target package error.
	g.Package = "x"
	g.Target = filepath.Join(suite.modInfo.Dir, "testdata", "target", "err")
	suite.Error(g.init())
	suite.resetTP(g)

	// ModInfo error.
	// Out of module directory.
	wd, err := os.Getwd()
	suite.Require().NoError(err)
	suite.Require().NoError(os.Chdir(suite.T().TempDir())) // change wd
	defer func() {
		suite.Require().NoError(os.Chdir(wd))
	}()
	suite.ErrorContains((&Genenv{Variables: map[string]Spec{"foo": {}}}).init(), "out of module directory")
}
