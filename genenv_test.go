package genenv

import (
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

func (suite *GenenvTestSuite) TestGenerate() {
	cfg := &Config{
		Package: "env",
		Target:  filepath.Join("testdata", "target", "gen"),
		Variables: map[string]Spec{
			"FOO": {Allow: []string{"foo"}},
			"BAR": {Deny: []string{"bar"}},
			"FOO_BAR": {
				Allow: []string{"foo"},
				Deny:  []string{"bar"},
			},
			"BAZ": {
				Doc: "// doc comment",
			},
			"FOO_BAZ": {
				Required: true,
			},
			"BAR_BAZ": {
				Required: true,
				Allow:    []string{"foo"},
				Deny:     []string{"bar"},
			},
		},
	}

	suite.NoError(Generate(cfg))
}
