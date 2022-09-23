package generator

import (
	"errors"
	"math/rand"
	"time"
)

type Date struct {
	min    int64
	max    int64
	format string
}

func NewDate(params map[string]interface{}) (g Generator, err error) {
	s := &Date{format: "2006-01-02 15:04:05"}
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
	mid, err := time.Parse("2006-01-02T15:04:05 MST", min.(string))
	if err != nil {
		return nil, errors.New("min time parse error:" + err.Error())
	}
	s.min = mid.Unix()

	max, ok := params["max"]
	if !ok {
		return nil, errors.New("max is not set in the parameter")
	}
	mad, err := time.Parse("2006-01-02T15:04:05 MST", max.(string))
	if err != nil {
		return nil, errors.New("max time parse error:" + err.Error())
	}
	s.max = mad.Unix()

	format, ok := params["format"]
	if ok {
		s.format = format.(string)
	}

	return
}

func (s *Date) Generate() (string, error) {
	rand.Seed(time.Now().UnixNano())
	mv := int64(s.max - s.min)
	return time.Unix(rand.Int63n(mv)+int64(s.min), 0).Format(s.format), nil
}
