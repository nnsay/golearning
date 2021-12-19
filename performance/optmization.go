package performance

import (
	"encoding/json"
	"strconv"
	"strings"
)

func createRequest() string {
	payload := make([]int, 100, 100)
	for i := 0; i < len(payload); i++ {
		payload[i] = i
	}
	req := Request{"demo_request", payload}
	v, err := json.Marshal(&req)
	if err != nil {
		panic(err)
	}
	return string(v)
}

func processRequestOld(reqs []string) []string {
	reps := []string{}
	for _, req := range reqs {
		reqObj := &Request{}
		json.Unmarshal([]byte(req), reqObj)
		ret := ""
		for _, e := range reqObj.PlayLoad {
			ret += strconv.Itoa(e) + ","
		}
		resObj := &Response{reqObj.TransactionID, ret}
		resJson, err := json.Marshal(&resObj)
		if err != nil {
			panic(err)
		}
		reps = append(reps, string(resJson))
	}
	return reps
}

func processRequest(reqs []string) []string {
	reps := []string{}
	for _, req := range reqs {
		reqObj := &Request{}
		reqObj.UnmarshalJSON([]byte(req))
		// json.Unmarshal([]byte(req), reqObj)
		var buf strings.Builder
		for _, e := range reqObj.PlayLoad {
			buf.WriteString(strconv.Itoa(e))
			buf.WriteString(",")
		}
		resObj := &Response{reqObj.TransactionID, buf.String()}
		// resJson, err := json.Marshal(&resObj)
		resJson, err := resObj.MarshalJSON()
		if err != nil {
			panic(err)
		}
		reps = append(reps, string(resJson))
	}
	return reps
}
