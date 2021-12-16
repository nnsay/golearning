package easyjson

import "testing"

func TestEasyjson(t *testing.T) {
	jsonstr := `{"City": "Beijing","Country": "China"}`
	p := Address{}
	p.UnmarshalJSON([]byte(jsonstr))
	t.Logf("%#v", p)

	str, err := p.MarshalJSON()
	if err != nil {
		t.Log("error:", err)
	}
	t.Logf("MarshalJSON result: %s", str)
}
