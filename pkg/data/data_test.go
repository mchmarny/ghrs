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

func TestData(t *testing.T) {
	deleteDB()
	defer deleteDB()

	s, err := New(testDir)
	assert.NoError(t, err)
	assert.NotNil(t, s)

	val, err := s.Get("test")
	assert.NoError(t, err)
	assert.Equal(t, int64(0), val)

	err = s.Upsert("test", 1)
	assert.NoError(t, err)

	val, err = s.Get("test")
	assert.NoError(t, err)
	assert.Equal(t, int64(1), val)

	err = s.Upsert("test", int64(2))
	assert.NoError(t, err)

	val, err = s.Get("test")
	assert.NoError(t, err)
	assert.Equal(t, int64(2), val)

	err = s.Close()
	assert.NoError(t, err)
}
