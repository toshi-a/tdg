package generator_test

import (
	"fmt"
	"os"
	"strconv"
	"tdg/internal/generator"
	"testing"
)

var s *generator.Integer

func TestMain(m *testing.M) {
	status := m.Run()
	os.Exit(status)
}

func TestInteger_Generate(t *testing.T) {
	t.Setenv("APP_ENV", "test")

	min := 10
	max := 11
	params := make(map[string]interface{})
	params["min"] = float64(min)
	params["max"] = float64(max)
	s, _ := generator.NewInteger(params)

	for i := 0; i < 100; i++ {
		v, _ := s.Generate()
		n, _ := strconv.Atoi(v)
		if n < min || n > max {
			t.Error(fmt.Sprintf("generate error (min: %d, max: %d, generaetd: %d)", min, max, n))
		}
	}
}
