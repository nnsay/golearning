package reflect_test

import (
	"fmt"
	"reflect"
	"testing"
)

func checkType(obj interface{}) {
	t := reflect.TypeOf(obj)
	switch t.Kind() {
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Println("Int")
	case reflect.Float32, reflect.Float64:
		fmt.Println("Float")
	case reflect.String:
		fmt.Println("String")
	case reflect.Func:
		fmt.Println("Funcion")
	default:
		fmt.Println("Unknow type: ", t)
	}
}

// 使用反射获取类型
func TestBasicUse(t *testing.T) {
	checkType(12)
	checkType(12.12)
	checkType(checkType)
	checkType("checkType")
	checkType([...]int{1})
}

// 使用反射类型和值
func TestTypeAndValue(t *testing.T) {
	var pi float64 = 3.141592653
	t.Logf("pi type is: %v and value is: %v", reflect.TypeOf(pi), reflect.ValueOf(pi))
	// 获取对象值的类型
	t.Log(reflect.ValueOf(pi).Type())
}

// 使用relect写万能代码, 通过字符串调用对象方法和属性
type Employee struct {
	EmployeeID string
	Name       string `format:"normal"`
	Age        int
}

func (e *Employee) UpdateAge(age int) {
	e.Age = age
}

func TestInvokeByName(t *testing.T) {
	e := &Employee{"1234", "jimmy wang", 33}
	// 使用Name获取属性: FieldByName不能用指针, 需要对象值
	t.Logf("Name: value(%[1]v), type(%[1]T)", reflect.ValueOf(*e).FieldByName("Name"))
	t.Logf("Name: value(%[1]v), type(%[1]T)", reflect.ValueOf(e).Elem().FieldByName("Name"))
	// 获取属性类型
	// nameType, ok := reflect.TypeOf(*e).FieldByName("Name")
	nameType, ok := reflect.TypeOf(e).Elem().FieldByName("Name")
	if !ok {
		t.Logf("Name is not the filed in the object")
		return
	}
	// 获取属性Tag
	t.Log("Tag format: ", nameType.Tag.Get("format"))
	// 调用对象方法: MethodByName不能用对象值, 需要用指针
	newAge := reflect.ValueOf((1))
	reflect.ValueOf(e).MethodByName("UpdateAge").Call([]reflect.Value{newAge})
	t.Logf("latest object: %#v", *e)
}
