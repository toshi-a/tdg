package config

type Item struct {
	Name      string                 `json:"name"`
	Generator string                 `json:"generator"`
	Params    map[string]interface{} `json:"params"`
	Quote     bool                   `json:"quote"`
}

type Output struct {
	Writer string                 `json:"writer"`
	Params map[string]interface{} `json:"params"`
}

type Config struct {
	Items  []Item `json:"items"`
	Output Output `json:"output"`
}
