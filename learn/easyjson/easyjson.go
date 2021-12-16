package easyjson

type Address struct {
	City     string
	Country  string `json:",omitempty"`
	Province string `json:"Continent"`
}
