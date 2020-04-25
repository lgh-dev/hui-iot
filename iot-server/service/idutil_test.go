package service

import (
	assert2 "github.com/stretchr/testify/assert"
	"testing"
)

func TestIDTools_NextId(t *testing.T) {
	assert2.True(t, IDTool.NextId() > 0, "")
}

func TestIDTools_NextIdStr(t *testing.T) {
	assert2.True(t, IDTool.NextIdStr() != "", "")
}
func BenchmarkNextId(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IDTool.NextId()
	}
}

func BenchmarkNextIdStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IDTool.NextIdStr()
	}
}
