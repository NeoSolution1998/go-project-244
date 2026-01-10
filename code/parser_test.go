package code

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseFileSuccess(t *testing.T) {
	path := filepath.Join("..", "testdata", "fixtures", "file1.json")

	data, err := parseFile(path)

	require.NoError(t, err)
	assert.Equal(t, map[string]any{
		"host":    "hexlet.io",
		"timeout": float64(50),
		"proxy":   "123.234.53.22",
		"follow":  false,
	}, data)
}

func TestParseFileInvalidJSON(t *testing.T) {
	path := filepath.Join("..", "testdata", "fixtures", "invalid.json")

	_, err := parseFile(path)

	assert.Error(t, err)
}

func TestParseFileNotExists(t *testing.T) {
	_, err := parseFile("no_such_file.json")

	assert.Error(t, err)
}
