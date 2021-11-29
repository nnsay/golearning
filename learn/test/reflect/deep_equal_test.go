package reflect_test

import (
	"errors"
	"reflect"
	"testing"
)

type Worker struct {
	EmployeeID string
	Name       string `format:"normal"`
	Age        int
}

type Person struct {
	Name string
}

type Peoper2 struct {
	Name string
}

func TestDeepEqual(t *testing.T) {
	m1 := map[int]string{1: "one", 2: "two", 3: "three"}
	m2 := map[int]string{
		1: "one", 3: "three", 2: "two",
		// 4: "four",
	}

	// t.Logf("m1 == m2 result: %v", m1 == m2)
	// map比较值查看key和value,不在乎书写顺序
	t.Logf("m1 == m2 result: %v", reflect.DeepEqual(m1, m2))

	// 切片比较值查看顺序和值
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{1, 2, 3, 4, 5}
	s3 := []int{5, 2, 3, 4, 1}
	t.Logf("s1 == s2 result: %v", reflect.DeepEqual(s1, s2))
	t.Logf("s1 == s3 result: %v", reflect.DeepEqual(s1, s3))

	// 结构体比较: 类型/属性/值
	p1 := Person{"jimmy"}
	p2 := Person{"jimmy"}
	p3 := Peoper2{"jimmy"}
	t.Logf("p1 == p2 result: %v", reflect.DeepEqual(p1, p2))
	t.Logf("p1 == p3 result: %v", reflect.DeepEqual(p1, p3))
}

func fillValue(obj interface{}, data map[string]interface{}) error {
	if reflect.TypeOf(obj).Kind() != reflect.Ptr {
		return errors.New("first parameter must be point type")
	}
	if reflect.TypeOf(obj).Elem().Kind() != reflect.Struct {
		return errors.New("first parameter must be struct type")
	}
	if data == nil {
		return errors.New("second parameter is require")
	}

	for key, value := range data {
		filed, ok := reflect.TypeOf(obj).Elem().FieldByName(key)
		if !ok {
			continue
		}
		if filed.Type != reflect.TypeOf(value) {
			continue
		}
		reflect.ValueOf(obj).Elem().FieldByName(key).Set(reflect.ValueOf((value)))
	}

	return nil
}

func TestSetValue(t *testing.T) {
	e := &Worker{}
	p := &Person{}
	data := map[string]interface{}{"Name": "jimmy", "Age": 10}
	if ee := fillValue(e, data); ee != nil {
		t.Error(ee)
	}
	if pe := fillValue(p, data); pe != nil {
		t.Error(pe)
	}
	t.Logf("%#v", e)
	t.Logf("%#v", p)
}

/**
Knowledge
1. 反射可以增加复用性
2. 反射会带来性能问题
3. 反射相关代码可读性较差
*/
