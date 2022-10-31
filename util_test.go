package genenv

import (
	"os"

	"golang.org/x/mod/modfile"
)

func (suite *GenenvTestSuite) TestGomod() {
	var data []byte
	data, err := os.ReadFile(suite.modInfo.GoMod)
	suite.Require().NoError(err)

	var mf *modfile.File
	mf, err = modfile.Parse(suite.modInfo.GoMod, data, nil)
	suite.Require().NoError(err)

	suite.Require().Equal("github.com/marcozac/genenv", mf.Module.Mod.String())
}

func (suite *GenenvTestSuite) TestReadConfig() {
	g, err := ReadConfig(suite.f)
	suite.NoError(err)

	suite.Require().NotNil(g)
	suite.Require().NotNil(g.Variables)

	suite.Require().Contains(g.Variables, "FOO")
	suite.Require().Contains(g.Variables, "BAR")

	suite.Contains(g.Variables["FOO"].Allow, "foo")
	suite.Contains(g.Variables["BAR"].Deny, "bar")
}
