package service

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	jwt, err := GenerateToken("lgh")
	if err != nil {
		log.Println("err")
	}
	assert.True(t, err == nil, "GenerateToken err")
	log.Printf("jwt Generate succ: %s\n", jwt)
}

func TestParseToken(t *testing.T) {
	jwt, err := GenerateToken("lgh")
	if err != nil {
		log.Println("err")
	}
	claims, err := ParseToken(jwt)
	if err != nil {
		log.Println("err")
	}
	assert.True(t, err == nil, "GenerateToken err")
	assert.True(t, claims.AppKey == "lgh", "Parse jwt err")
	json, _ := json.Marshal(claims)
	log.Printf("jwt Parse succ: %s\n", json)
}

func BenchmarkGenerateToken(b *testing.B) {
	b.SetParallelism(4)
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			GenerateToken("lgh")
			//log.Printf("get jwt %s\n",jwt)
		}
	})
}
