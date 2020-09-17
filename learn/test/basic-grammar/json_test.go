package golearning_test

import (
	"encoding/json"
	"testing"
)

type User struct {
	Name    string
	Age     int
	Address Address  `json:"home_address"`
	Hobby   []string `json:"hobbies"`
}

type Address struct {
	City     string
	Country  string `json:",omitempty"`
	Province string `json:"Continent"`
}

func TestJsonBasic(t *testing.T) {
	// marshal
	user1 := User{Name: "Jimmy", Age: 42, Address: Address{Country: "CN", Province: "Beijing", City: "Beijing"}}
	data, err := json.Marshal(user1)
	if err != nil {
		t.Errorf("marshal error: %v", err)
	}
	t.Logf("data: %s", data)
	// marshal index
	data, _ = json.MarshalIndent(user1, "", "  ")
	t.Logf("data: \n%s", data)

	// unmarshal
	var user2 = &User{}
	json.Unmarshal(data, user2)
	t.Logf("user: \n%#v", *user2)
}
