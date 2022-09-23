package generator

import (
	"errors"
)

type Static struct {
	value string
}

func NewStatic(params map[string]interface{}) (g Generator, err error) {
	s := &Static{}
	err = nil
	defer func() {
		catch := recover()
		if catch != nil {
			g = nil
			err = errors.New("param cast error")
		} else {
			g = s
		}
	}()

	v, ok := params["value"]
	if !ok {
		return nil, errors.New("value is not set in the parameter")
	}
	s.value = v.(string)

	return
}

func (s *Static) Generate() (string, error) {
	return s.value, nil
}
