package output

import (
	"errors"
	"tdg/internal/config"
)

type Manager struct {
	writersByName map[string]Constructor
}

func NewManager() *Manager {
	s := &Manager{
		writersByName: make(map[string]Constructor),
	}
	s.registerOutput("sql", NewSQL)
	return s
}

func (s *Manager) registerOutput(name string, constructor Constructor) {
	s.writersByName[name] = constructor
}

func (s *Manager) New(name string, items []config.Item, params map[string]interface{}) (Writer, error) {
	wc, ok := s.writersByName[name]
	if !ok {
		return nil, errors.New("the specified writer does not exist")
	}
	return wc(items, params)
}
