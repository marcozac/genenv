package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestMain(m *testing.M) {
	cfgFile = filepath.Join("..", "..", "testdata", ".genenv.yml")
	os.Exit(m.Run())
}

func TestCmd(t *testing.T) {
	err := rootCmd.Execute()
	if err != nil {
		t.Error(err)
	}
}
