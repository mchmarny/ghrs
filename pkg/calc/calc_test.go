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
	input[StateArg] = testDataPath
	input[KeyArg] = "cal"
	input[OperationArg] = "add"
	input[ValueArg] = "1"

	results, err := Calculate(input)
	assert.NoError(t, err)
	assert.NotNil(t, results)
	assert.Equal(t, 1, len(results))
	assert.Equal(t, "1", results[ResultArg])

	input[ValueArg] = "2"

	results, err = Calculate(input)
	assert.NoError(t, err)
	assert.NotNil(t, results)
	assert.Equal(t, 1, len(results))
	assert.Equal(t, "3", results[ResultArg])
}
