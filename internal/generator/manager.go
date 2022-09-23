package generator

import (
	"errors"
)

type Manager struct {
	generatorsByName map[string]Constructor
}

func NewManager() *Manager {
	s := &Manager{
		generatorsByName: make(map[string]Constructor),
	}
	s.registerGenerator("integer", NewInteger)
	s.registerGenerator("sequential", NewSequential)
	s.registerGenerator("enum", NewEnum)
	s.registerGenerator("static", NewStatic)
	s.registerGenerator("uuid", NewUUID)
	s.registerGenerator("date", NewDate)
	return s
}

func (s *Manager) registerGenerator(name string, generatorConstructor Constructor) {
	s.generatorsByName[name] = generatorConstructor
}

func (s *Manager) New(name string, params map[string]interface{}) (Generator, error) {
	gc, ok := s.generatorsByName[name]
	if !ok {
		return nil, errors.New("the specified generator does not exist")
	}
	return gc(params)
}
