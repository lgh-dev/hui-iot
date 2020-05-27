package db

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//
func TestTDEngineDB_CheckDBExists(t *testing.T) {
	db := *NewTDengine()
	assert.True(t, db.CheckDBExists(), "TestTDEngineDB_CheckTableExists err")
}
