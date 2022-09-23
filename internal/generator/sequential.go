package generator

import (
	"errors"
	"fmt"
)

type Sequential struct {
	digit   int32
	format  string
	initial int32
	current int32
}

func NewSequential(params map[string]interface{}) (g Generator, err error) {
	s := &Sequential{format: "%d", initial: 0, current: -1}
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

	digit, ok := params["digit"]
	if ok {
		s.digit = int32(digit.(float64))
		s.format = fmt.Sprintf("%%0%dd", s.digit)
	}

	initial, ok := params["initial"]
	if ok {
		s.initial = int32(initial.(float64))
		s.current = s.initial - 1
	}

	return
}

func (s *Sequential) Generate() (string, error) {
	s.current++
	return fmt.Sprintf(s.format, s.current), nil
}
