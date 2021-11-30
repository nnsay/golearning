package pipeline

import "errors"

type SumFilter struct{}

func (sf *SumFilter) Process(req Request) (Response, error) {
	data, ok := req.([]int)
	if !ok {
		return nil, errors.New("input data should be int slice")
	}

	var sum int
	for _, value := range data {
		sum += value
	}

	return sum, nil
}
