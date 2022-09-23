package output

import "tdg/internal/config"

type Constructor func(items []config.Item, params map[string]interface{}) (Writer, error)

type Writer interface {
	Write(rowNo int, values []string, last bool)
}
