package golearning_test

import (
	"bufio"
	"os"
	"testing"
	"text/template"
)

func ageToMonth(age int) int {
	return age * 12
}

func TestBasic(t *testing.T) {
	// marshal
	temstr := `{{.Name}}'s Work Report
	Age: {{.Age | ageToMonth}}
	Address: 
		Provice: {{.Address.Province}}
		City: {{.Address.City}}
	All hobby:
		{{range .Hobby}}{{.}} {{end}}
	`
	user1 := User{
		Name:    "Jimmy",
		Age:     30,
		Address: Address{Country: "CN", Province: "Beijing", City: "Beijing"},
		Hobby:   []string{"Codeing", "Swimming", "Singing"},
	}
	tmp := template.Must(template.New("test").Funcs(template.FuncMap{"ageToMonth": ageToMonth}).Parse(temstr))
	// standard out
	var rst = bufio.NewReader(os.Stdout)
	tmp.Execute(os.Stdout, user1)
	str, _ := rst.ReadBytes('\n')
	t.Logf(">%s<", str)
	// string: not work
	// var tpl bytes.Buffer
	// tmp.Execute(tpl, user1)
	// t.Logf(">> %s <<", tpl.String())
}
