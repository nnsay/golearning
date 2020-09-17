package golearning_test

import (
	"bufio"
	"os"
	"testing"
	"text/template"
)

func TestBasic(t *testing.T) {
	// marshal
	temstr := `{{.Name}}'s Work Report
	Age: {{.Age}}
	Address: 
		Provice: {{.Address.Province}}
		City: {{.Address.City}}
	All hobby:
		{{range .Hobby}}{{.}} {{end}}
	`
	user1 := User{
		Name:    "Jimmy",
		Age:     42,
		Address: Address{Country: "CN", Province: "Beijing", City: "Beijing"},
		Hobby:   []string{"Codeing", "Swimming", "Singing"},
	}
	tmp := template.Must(template.New("test").Parse(temstr))
	var rst = bufio.NewReader(os.Stdout)
	tmp.Execute(os.Stdout, user1)
	str, _ := rst.ReadString('\n')
	t.Logf("%s", str)
}
