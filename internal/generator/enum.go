package generator

import (
	"errors"
	"math/rand"
	"strings"
	"time"
)

type Enum struct {
	list  []string
	lists [][]string
}

func NewEnum(params map[string]interface{}) (g Generator, err error) {
	s := &Enum{}
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

	list, ok := params["list"]
	if ok {
		s.list = make([]string, len(list.([]interface{})))
		for i, v := range list.([]interface{}) {
			s.list[i] = v.(string)
		}
	}

	lists, ok := params["lists"]
	if ok {
		s.lists = make([][]string, len(lists.([]interface{})))
		for oi, ol := range lists.([]interface{}) {
			s.lists[oi] = make([]string, len(ol.([]interface{})))
			for ii, v := range ol.([]interface{}) {
				s.lists[oi][ii] = v.(string)
			}
		}
	}

	return
}

func (s *Enum) randomChoice(list []string) string {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(list))
	return list[i]
}

func (s *Enum) Generate() (string, error) {
	if s.lists != nil {
		vs := make([]string, 0, len(s.lists))
		for _, list := range s.lists {
			vs = append(vs, s.randomChoice(list))
		}
		return strings.Join(vs, " "), nil
	} else if s.list != nil {
		return s.randomChoice(s.list), nil
	} else {
		return "", nil
	}
	return "", nil
}
