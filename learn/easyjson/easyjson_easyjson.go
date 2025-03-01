// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package easyjson

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson8d5c760DecodeGoleaningLearnEasyjson(in *jlexer.Lexer, out *Address) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "City":
			out.City = string(in.String())
		case "Country":
			out.Country = string(in.String())
		case "Continent":
			out.Province = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson8d5c760EncodeGoleaningLearnEasyjson(out *jwriter.Writer, in Address) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"City\":"
		out.RawString(prefix[1:])
		out.String(string(in.City))
	}
	if in.Country != "" {
		const prefix string = ",\"Country\":"
		out.RawString(prefix)
		out.String(string(in.Country))
	}
	{
		const prefix string = ",\"Continent\":"
		out.RawString(prefix)
		out.String(string(in.Province))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Address) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson8d5c760EncodeGoleaningLearnEasyjson(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Address) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson8d5c760EncodeGoleaningLearnEasyjson(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Address) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson8d5c760DecodeGoleaningLearnEasyjson(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Address) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson8d5c760DecodeGoleaningLearnEasyjson(l, v)
}
