package data

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testDir = "test.db"
)

func deleteDB() {
	os.Remove(testDir)
}

var (
	testDate = map[string]int64{
		"test1":    0,
		"2test":    0,
		"bad-data": 0,
	}
)

func TestData(t *testing.T) {
	deleteDB()
	defer deleteDB()

	s, err := New(testDir)
	assert.NoError(t, err)
	assert.NotNil(t, s)

	err = s.SaveAll(testDate)
	assert.NoError(t, err)

	ids, err := s.Query("test")
	assert.NoError(t, err)
	assert.Equal(t, len(ids), 2)

	ok, err := s.Update("2test", 1)
	assert.NoError(t, err)
	assert.True(t, ok)

	val, err := s.Get("2test")
	assert.NoError(t, err)
	assert.Equal(t, val, int64(1))

	err = s.Close()
	assert.NoError(t, err)
}
