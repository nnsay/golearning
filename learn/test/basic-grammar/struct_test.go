package golearning_test

import "testing"

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
