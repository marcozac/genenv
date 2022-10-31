package genenv

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/suite"
)

type GenenvTestSuite struct {
	suite.Suite
	modInfo *ModInfo
	f       string // configuration file path
}

func TestSuite(t *testing.T) {
	suite.Run(t, &GenenvTestSuite{})
}

func (suite *GenenvTestSuite) SetupSuite() {
	var err error
	suite.modInfo, err = GetModInfo()
	if err != nil {
		suite.FailNow("getting module info", err)
	}

	suite.f = filepath.Join("testdata", ".genenv.yml")
}
