package unittest

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConcatString(t *testing.T) {
	assert := assert.New(t)
	elems := [...]string{"1", "2", "3", "4", "5"}
	var ret string
	for _, e := range elems {
		ret += e
	}
	assert.Equal("12345", ret)
}

func TestConcatStringByButter(t *testing.T) {
	assert := assert.New(t)
	elems := [...]string{"1", "2", "3", "4", "5"}
	var ret bytes.Buffer
	for _, e := range elems {
		ret.WriteString(e)
	}
	assert.Equal("12345", ret.String())
}

// 改写上面两个普通测试为benchmark测试
func BenchmarkConcatString(b *testing.B) {
	elems := [...]string{"1", "2", "3", "4", "5"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var ret string
		for _, e := range elems {
			ret += e
		}
	}
	b.StopTimer()
}

func BenchmarkConcatStringByButter(b *testing.B) {
	elems := [...]string{"1", "2", "3", "4", "5"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var ret bytes.Buffer
		for _, e := range elems {
			ret.WriteString(e)
		}
	}
	b.StopTimer()
}
