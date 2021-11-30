package pipeline

import (
	"errors"
	"strconv"
)

type CovFilter struct{}

func (cf *CovFilter) Process(req Request) (Response, error) {
	data, ok := req.([]string)
	if !ok {
		return nil, errors.New("input data should be string slice")
	}
	dataIntSlice := []int{}

	for _, value := range data {
		intValue, err := strconv.Atoi(value)
		if err != nil {
			continue
		}
		dataIntSlice = append(dataIntSlice, intValue)
	}
	return dataIntSlice, nil
}
