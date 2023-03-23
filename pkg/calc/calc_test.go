package calc

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testDataPath = "calc/test.db"
)

func deleteDB() {
	os.RemoveAll(filepath.Dir(testDataPath))
}

func TestCalc(t *testing.T) {
	deleteDB()
	defer deleteDB()

	input := GetArgs()
	input[DataKey] = testDataPath
	input[CounterKey] = "cal"
	input[ActionKey] = "add"
	input[ValueKey] = "1"

	results, err := Calculate(input)
	assert.NoError(t, err)
	assert.NotNil(t, results)
	assert.Equal(t, 1, len(results))
	assert.Equal(t, "1", results[ResultKey])

	input[ValueKey] = "2"

	results, err = Calculate(input)
	assert.NoError(t, err)
	assert.NotNil(t, results)
	assert.Equal(t, 1, len(results))
	assert.Equal(t, "3", results[ResultKey])
}
