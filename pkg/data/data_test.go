package data

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testDir = "data/test.db"
)

func deleteDB() {
	os.RemoveAll(filepath.Dir(testDir))
}

func TestData(t *testing.T) {
	deleteDB()
	defer deleteDB()

	testKey := "test"

	s, err := New(testDir)
	assert.NoError(t, err)
	assert.NotNil(t, s)

	val, err := s.Get(testKey)
	assert.NoError(t, err)
	assert.Equal(t, int64(0), val)

	err = s.Upsert(testKey, 1)
	assert.NoError(t, err)

	err = s.Update(testKey, 1)
	assert.NoError(t, err)

	val, err = s.Get(testKey)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), val)

	err = s.Upsert(testKey, int64(2))
	assert.NoError(t, err)

	val, err = s.Get(testKey)
	assert.NoError(t, err)
	assert.Equal(t, int64(2), val)

	results, err := s.Query("%%es%%")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(results))

	list := map[string]int64{
		"test1": 1,
		"test2": 2,
		"test3": 3,
	}
	err = s.SaveAll(list)
	assert.NoError(t, err)

	results, err = s.Query("%%es%%")
	assert.NoError(t, err)
	assert.Equal(t, len(list)+1, len(results))

	err = s.Delete(testKey)
	assert.NoError(t, err)

	err = s.Close()
	assert.NoError(t, err)
}
