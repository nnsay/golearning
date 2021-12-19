// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package performance

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

func easyjson6a975c40DecodeGoleaningPerformance(in *jlexer.Lexer, out *Response) {
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
		case "transaction_id":
			out.TransactionID = string(in.String())
		case "exp":
			out.Expression = string(in.String())
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
func easyjson6a975c40EncodeGoleaningPerformance(out *jwriter.Writer, in Response) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"transaction_id\":"
		out.RawString(prefix[1:])
		out.String(string(in.TransactionID))
	}
	{
		const prefix string = ",\"exp\":"
		out.RawString(prefix)
		out.String(string(in.Expression))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Response) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6a975c40EncodeGoleaningPerformance(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Response) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6a975c40EncodeGoleaningPerformance(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Response) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6a975c40DecodeGoleaningPerformance(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Response) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6a975c40DecodeGoleaningPerformance(l, v)
}
func easyjson6a975c40DecodeGoleaningPerformance1(in *jlexer.Lexer, out *Request) {
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
		case "transaction_id":
			out.TransactionID = string(in.String())
		case "playload":
			if in.IsNull() {
				in.Skip()
				out.PlayLoad = nil
			} else {
				in.Delim('[')
				if out.PlayLoad == nil {
					if !in.IsDelim(']') {
						out.PlayLoad = make([]int, 0, 8)
					} else {
						out.PlayLoad = []int{}
					}
				} else {
					out.PlayLoad = (out.PlayLoad)[:0]
				}
				for !in.IsDelim(']') {
					var v1 int
					v1 = int(in.Int())
					out.PlayLoad = append(out.PlayLoad, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjson6a975c40EncodeGoleaningPerformance1(out *jwriter.Writer, in Request) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"transaction_id\":"
		out.RawString(prefix[1:])
		out.String(string(in.TransactionID))
	}
	{
		const prefix string = ",\"playload\":"
		out.RawString(prefix)
		if in.PlayLoad == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.PlayLoad {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.Int(int(v3))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Request) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6a975c40EncodeGoleaningPerformance1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Request) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6a975c40EncodeGoleaningPerformance1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Request) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6a975c40DecodeGoleaningPerformance1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Request) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6a975c40DecodeGoleaningPerformance1(l, v)
}
