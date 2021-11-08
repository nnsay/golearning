package golearning_test

import (
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestStringUnionCode(t *testing.T) {
	// lengh取的字节byte
	c1 := "hello"
	t.Logf("lenght: %d", len(c1))
	c2 := "中国"
	t.Logf("lenght: %d", len(c2))
	c3 := []rune(c2)
	t.Logf("rune lenght: %d", len(c3))
	// 字符集值
	c4 := "中"
	cr4 := []rune(c4)
	t.Logf("unioncode: %x", c4)
	t.Logf("unioncode: %x, length: %d", cr4[0], len(cr4))
	t.Logf("中的 unicode %x", cr4[0])
	t.Logf("中的 utf8 %x", c4)
	// 字符遍历
	for _, r := range "中国人民共和国" {
		t.Logf("%[1]c %[1]x", r)
	}
}

func TestStringFn(t *testing.T) {
	// split
	s := "大,家,好,啊"
	parts := strings.Split(s, ",")
	for _, v := range parts {
		t.Logf("%s", v)
	}
	// join
	t.Log("join result:", strings.Join(parts, " "))
}

func TestCovert(t *testing.T) {
	s := "10"
	if v, err := strconv.Atoi(s); err == nil {
		t.Logf("字符串转整型: %d, %s", v, reflect.TypeOf(v).String())
		t.Logf("整型转字符串: %#v", strconv.Itoa(v))
	}
}
