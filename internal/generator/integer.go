package generator

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Integer struct {
	min int32
	max int32
}

func NewInteger(params map[string]interface{}) (g Generator, err error) {
	s := &Integer{}
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

	min, ok := params["min"]
	if !ok {
		return nil, errors.New("min is not set in the parameter")
	}
	s.min = int32(min.(float64))

	max, ok := params["max"]
	if !ok {
		return nil, errors.New("max is not set in the parameter")
	}
	s.max = int32(max.(float64))

	return
}

func (s *Integer) Generate() (string, error) {
	rand.Seed(time.Now().UnixNano())
	mv := int64(s.max - s.min)
	return fmt.Sprintf("%d", rand.Int63n(mv)+int64(s.min)), nil
}
