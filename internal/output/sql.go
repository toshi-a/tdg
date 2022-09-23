package output

import (
	"errors"
	"fmt"
	"strings"
	"tdg/internal/config"
)

type SQL struct {
	tableName string
	format    string
}

func NewSQL(items []config.Item, params map[string]interface{}) (w Writer, err error) {
	s := &SQL{}
	err = nil
	defer func() {
		catch := recover()
		if catch != nil {
			w = nil
			err = errors.New("param cast error")
		} else {
			w = s
		}
	}()

	tableName, ok := params["table_name"]
	if !ok {
		return nil, errors.New("table_name is not set in the parameter")
	}
	s.tableName = tableName.(string)

	cns := make([]string, 0, len(items))
	cfs := make([]string, 0, len(items))
	for _, item := range items {
		cns = append(cns, item.Name)
		if item.Quote {
			cfs = append(cfs, "'%s'")
		} else {
			cfs = append(cfs, "%s")
		}
	}
	s.format = "INSERT INTO " +
		s.tableName +
		" (" +
		strings.Join(cns, ", ") +
		") VALUES (" +
		strings.Join(cfs, ", ") +
		");"
	return
}

func (s *SQL) Write(rowNo int, values []string, last bool) {
	avs := make([]any, 0, len(values))
	for _, v := range values {
		avs = append(avs, v)
	}
	var nl string
	if last {
		nl = ""
	} else {
		nl = "\n"
	}
	fmt.Printf(s.format+nl, avs...)
}
