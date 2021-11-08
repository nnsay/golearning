package golearning_test

import (
	"math"
	"reflect"
	"sort"
	"testing"
)

func TestBasicUsage(t *testing.T) {
	// 声明一个map
	m1 := make(map[int]string)
	// 修改值
	m1[1] = "hello"
	m1[2] = "world"
	m1[3] = "!"
	t.Log(m1, len(m1))

	// 获取值
	if v, ok := m1[2]; ok {
		t.Logf("the m1 has the key: 2 and the value is %v", v)
	}

	// 遍历
	for key, value := range m1 {
		t.Logf("key: %v, value: %v", key, value)
	}

	// 删除
	delete(m1, 3)
	t.Log("deleted value:", m1)

	// 初始化声明
	m2 := map[string]string{"name": "Jimmy", "city": "Beijing"}
	t.Log("m2:", m2)
}

func TestLoopMap(t *testing.T) {
	me := make(map[string]string)
	me["weight"] = "75KG"
	me["age"] = "30 years old"
	me["name"] = "Jimmy Wang"

	for key, value := range me {
		t.Logf("%s: %s", key, value)
	}
}

func TestDeclearAndSetValue(t *testing.T) {
	m := map[string]string{"name": "Jimmy"}
	t.Log(m)
}

func TestFunValueMap(t *testing.T) {
	m := make(map[int]func(op int) int)
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return int(math.Pow(float64(op), 3)) }

	op := 2
	t.Logf("Orinal Number: %d", op)
	t.Logf("1 function result: %d", m[1](op))
	t.Logf("2 function result: %d", m[2](op))
	t.Logf("3 function result: %d", m[3](op))
}

func typeof(v interface{}) string {
	return reflect.TypeOf(v).String()
}

func TestMockSetWithMap(t *testing.T) {
	// set的类型是key的类型,value是bool类型
	s := map[int]bool{}
	t.Logf("s's type is: %s", typeof(s))
	addItem := func(s map[int]bool, v int) {
		s[v] = true
	}
	checkItem := func(s map[int]bool, v int) {
		if s[v] {
			t.Logf("%d is existing", v)
			return
		}
		t.Logf("%d is not existing", v)
	}
	deleteItem := func(s map[int]bool, v int) {
		delete(s, v)
	}
	// 添加元素
	s[1] = true
	addItem(s, 3)
	// 检查元素
	checkItem(s, 1)
	checkItem(s, 2)
	// 获取长度
	t.Logf("length: %d", len(s))

	// 删除元素
	deleteItem(s, 1)
	checkItem(s, 1)
	deleteItem(s, 1) // 已经删除的元素也可安全删除

	// 打印目前的值
	addItem(s, 1)
	addItem(s, 2)
	addItem(s, 3)
	keys := make([]int, 0, len(s))
	for k := range s {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	t.Logf("value: %#v", keys)
}
