package pipeline

type StraightPipeline struct {
	Name    string
	Filters *[]Filter
}

func NewStraightPipeline(name string, filters ...Filter) *StraightPipeline {
	return &StraightPipeline{
		Name:    name,
		Filters: &filters,
	}
}

func (sp *StraightPipeline) Process(data Request) (Response, error) {
	var ret interface{}
	var err error
	for _, f := range *sp.Filters {
		ret, err = f.Process(data)
		if err != nil {
			return ret, err
		}
		data = ret
	}
	return ret, err
}
