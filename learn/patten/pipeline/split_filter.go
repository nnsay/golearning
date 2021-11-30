package pipeline

import (
	"errors"
	"strings"
)

type SplitFilter struct{}

func (sf *SplitFilter) Process(req Request) (Response, error) {
	data, ok := req.(string)
	if !ok {
		return nil, errors.New("input data should be string")
	}
	dataSlice := strings.Split(data, ",")
	return dataSlice, nil
}
