package golearning_test

import (
	"math/rand"
	"testing"
	"time"
)

func TestRandPass(t *testing.T) {
	// t.Log(rand.Int63())
	// t.Log(rand.Int63())
	// t.Log(rand.Int63())
	// t.Logf("%#v", rand.NewSource(time.Now().UnixNano()+rand.Int63()))
	t.Log(rand.Intn(10))
	t.Log(rand.Intn(10))
	t.Log(rand.Intn(10))
	t.Log(rand.Intn(10))
	pass := genRandPass()
	t.Log(pass)
	var str string
	if str == "" {
		t.Log("empty")
	}
}

func genRandPass() string {
	baseStr := "abcdefghijkmnpqrstuvwxyzABCDEFGHJKMNPQRSTUVWXYZ234567890"
	passLength := 10
	r := rand.New(rand.NewSource(time.Now().UnixNano() + rand.Int63()))
	bytes := make([]byte, passLength)
	for i := 0; i < passLength; i++ {
		bytes[i] = baseStr[r.Intn(len(baseStr))]
	}
	return string(bytes)
}
