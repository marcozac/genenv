package main

import (
	"log"
	"path/filepath"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/suite"
)

type CmdTestSuite struct {
	suite.Suite
	cmd *cobra.Command
	err error
}

func TestSuite(t *testing.T) {
	suite.Run(t, &CmdTestSuite{cmd: rootCmd()})
}

func (suite *CmdTestSuite) TestCmd() {
	suite.Error(suite.cmd.Execute(), "ReadConfig error not reported")

	cfgFile = filepath.Join("..", "..", "testdata", ".genenv.yml")
	suite.NoError(suite.cmd.Execute(), "ReadConfig error not reported")
}

func (suite *CmdTestSuite) TestMain() {
	fatal = suite.fatalMock
	defer func() {
		fatal = log.Fatal
	}()

	main()
	suite.Error(suite.err, "ReadConfig error not reported")
}

func (suite *CmdTestSuite) fatalMock(v ...interface{}) {
	suite.err = v[0].(error)
}
