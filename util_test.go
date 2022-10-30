package genenv

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/mod/modfile"
)

func TestGomod(t *testing.T) {
	var err error
	var info *ModInfo

	info, err = GetModInfo()
	require.NoError(t, err)

	var data []byte
	data, err = os.ReadFile(info.GoMod)
	require.NoError(t, err)

	var mf *modfile.File
	mf, err = modfile.Parse(info.GoMod, data, nil)
	require.NoError(t, err)

	assert.Equal(t, "github.com/marcozac/genenv", mf.Module.Mod.String())
}
