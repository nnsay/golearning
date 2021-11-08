package golearning_test

import (
	"fmt"
	"testing"
	"unsafe"
)

type Salary struct {
	overtime float64
	basic    float64
}

type Employee struct {
	ID       int
	Name     string
	Position string
	salary   Salary
}

func (e *Employee) SetSalary(salary Salary) {
	e.salary = salary
}

func SetSalaryV2(e *Employee, basic float64) {
	e.salary.basic = basic
}

func TestTwoType(t *testing.T) {
	e := Employee{1, "jimmy", "Beijing", Salary{}}
	e1 := Employee{ID: 1, Name: "jimmy1", Position: "Beijing", salary: Salary{}}
	e2 := new(Employee)
	e2.ID = 1
	e2.Name = "jimmy2"
	e2.Position = "Shanghai"
	t.Log(e)
	t.Log(e1)
	t.Log(e2)
	t.Log(e1.Position, e1.salary.basic)
	t.Logf("e type: %T", e1) // 类型
	t.Logf("e type: %T", e2) // 指针
}

// 使用结构体定义的方法, 属性地址不一样, 说明发生了复制
// func (e Employee) String() string {
// 	fmt.Printf("address: %x \n", unsafe.Pointer(&e.Position))
// 	return fmt.Sprintf("ID:%d; Name:%s", e.ID, e.Name)
// }

// 使用指针定义的方法, 属性地址一样, 说明是引用传递
func (e *Employee) String() string {
	fmt.Printf("address: %x \n", unsafe.Pointer(&e.Position))
	return fmt.Sprintf("ID:%d; Name:%s", e.ID, e.Name)
}
func TestTwoTypeFunc(t *testing.T) {
	e := Employee{1, "jimmy", "Beijing", Salary{}}
	// e := &Employee{1, "jimmy", "Beijing", Salary{}}
	// e := new(Employee)
	// e.Position = "Beijing"
	fmt.Printf("address: %x \n", unsafe.Pointer(&e.Position))
	t.Log(e.String())
	// 总结:
	// 1. new(Employee), Employee, &Employee 三个方式在调用和对象使用上一致, 也说明三个方式的产生的对象都是一样即都是指针
	// 2. 使用结构体定义的方法会发生值复制传递, 而指针类型不会
	// 3. 推荐使用指针定义方法
}

func TestStructBasic(t *testing.T) {
	var e Employee
	t.Logf("e is %#v", e)
	var e2 = Employee{}
	t.Logf("e is %#v", e2)
	e3 := &Employee{}
	t.Logf("e is %#v", *e3)
	e4 := &Employee{}
	e4.SetSalary(Salary{basic: 0.0, overtime: 1999.0})
	t.Logf("e is %#v", e4)
	t.Logf("e basic salary is: %.2f", e4.salary.basic)
	SetSalaryV2(e4, 13.5)
	t.Logf("e basic salary is: %.2f", (*e4).salary.basic)
}

type Point struct {
	X, y int
}

type Circle struct {
	Center Point
	Radius int
}

type Wheel struct {
	Circle Circle
	Spokes int
}

func TestExportAbility(t *testing.T) {
	var p = Point{X: 1, y: 2}
	t.Logf("p:%#v", p)

	// anonymity member does not work
	var w Wheel
	w.Circle.Center.X = 8
	w.Circle.Center.y = 9
	w.Circle.Radius = 5
	// w.X = 8
	// w.y = 9
	// w.Radius = 5
	w.Spokes = 20
	t.Logf("%#v", w)
	t.Logf("%v", w)
}
