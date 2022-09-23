package generator

type Constructor func(params map[string]interface{}) (Generator, error)

type Generator interface {
	Generate() (string, error)
}
