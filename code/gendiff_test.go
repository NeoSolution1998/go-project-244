package code

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenDiffFlatJSON(t *testing.T) {
	path1 := filepath.Join("..", "testdata", "fixtures", "file1.json")
	path2 := filepath.Join("..", "testdata", "fixtures", "file2.json")
	expectedPath := filepath.Join("..", "testdata", "fixtures", "expected_stylish.txt")

	expectedBytes, err := os.ReadFile(expectedPath)
	require.NoError(t, err)

	result, err := GenDiff(path1, path2)
	require.NoError(t, err)

	assert.Equal(t, string(expectedBytes), result)
}
func TestGenDiffYAML(t *testing.T) {
	path1 := filepath.Join("..", "testdata", "fixtures", "file1.yml")
	path2 := filepath.Join("..", "testdata", "fixtures", "file2.yml")
	expectedPath := filepath.Join("..", "testdata", "fixture", "expected_stylish.txt")

	expected, err := os.ReadFile(expectedPath)
	require.NoError(t, err)

	result, err := GenDiff(path1, path2)
	require.NoError(t, err)

	assert.Equal(t, string(expected), result)
}

func TestBuildDiffFlat(t *testing.T) {
	data1 := map[string]any{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	data2 := map[string]any{
		"b": 2,
		"c": 4,
		"d": 5,
	}

	diff := buildDiff(data1, data2)

	expected := []DiffNode{
		{Key: "a", Status: StatusRemoved, OldVal: 1},
		{Key: "b", Status: StatusUnchanged, OldVal: 2},
		{Key: "c", Status: StatusUpdated, OldVal: 3, NewVal: 4},
		{Key: "d", Status: StatusAdded, NewVal: 5},
	}

	assert.Equal(t, expected, diff)
}

func TestFormatStylish(t *testing.T) {
	nodes := []DiffNode{
		{Key: "a", Status: StatusRemoved, OldVal: 1},
		{Key: "b", Status: StatusUnchanged, OldVal: 2},
		{Key: "c", Status: StatusUpdated, OldVal: 3, NewVal: 4},
	}

	result := formatStylish(nodes)

	expected := `{
  - a: 1
    b: 2
  - c: 3
  + c: 4
}`

	assert.Equal(t, expected, result)
}
