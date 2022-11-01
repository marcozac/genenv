package genenv

import (
	"os"

	"golang.org/x/mod/modfile"
)

func (suite *GenenvTestSuite) TestModInfo() {
	var data []byte
	data, err := os.ReadFile(suite.modInfo.GoMod)
	suite.Require().NoError(err)

	// Get current module info.
	var mf *modfile.File
	mf, err = modfile.Parse(suite.modInfo.GoMod, data, nil)
	suite.Require().NoError(err)
	suite.Require().Equal("github.com/marcozac/genenv", mf.Module.Mod.String())
}
