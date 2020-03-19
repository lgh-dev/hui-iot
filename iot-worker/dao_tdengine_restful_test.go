package iot_worker

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTDengineRestful_CheckDBExists(t *testing.T) {
	db := *NewTDengineRestful()
	assert.True(t, db.CheckDBExists(), "TestTDengineRestful_CheckDBExists err")
}
