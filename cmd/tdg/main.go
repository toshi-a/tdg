package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"tdg/internal/config"
	"tdg/internal/generator"
	"tdg/internal/output"
)

func main() {
	maxNumRow := 99999

	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "\nusage: %s definition_json_path output_rows\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  definition_json_path: path of definition json file.\n")
		fmt.Fprintf(os.Stderr, "           output_rows: number of output rows.\n")
		os.Exit(1)
	}

	definitionJSONPath := os.Args[1]
	if f, err := os.Stat(definitionJSONPath); os.IsNotExist(err) || f.IsDir() {
		fmt.Fprintf(os.Stderr, "\nDefinition file not found: %s \n", definitionJSONPath)
		os.Exit(2)
	}
	numOutputRows, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "\nNumber of output rows is not natural number: %s \n", os.Args[2])
		os.Exit(3)
	}
	if numOutputRows <= 0 || numOutputRows > maxNumRow {
		fmt.Fprintf(os.Stderr, "\nNumber of output rows is not 1 to %d natural number: %s \n", maxNumRow, os.Args[2])
		os.Exit(4)
	}

	b, err := os.ReadFile(definitionJSONPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "\ndefinition file read error: %s\n", err)
		os.Exit(5)
	}
	configStr := string(b)
	config := config.Config{}
	err = json.Unmarshal([]byte(configStr), &config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "\ndefinition file read error: %s\n", err)
		os.Exit(6)
	}

	gm := generator.NewManager()

	// Set generators for per items(columns)
	gs := make([]generator.Generator, 0)
	for _, item := range config.Items {
		ig, _ := gm.New(item.Generator, item.Params)
		gs = append(gs, ig)
	}

	om := output.NewManager()
	w, _ := om.New(config.Output.Writer, config.Items, config.Output.Params)

	// Items write loop
	for i := 0; i < numOutputRows; i++ {
		vs := make([]string, 0, len(gs))
		for _, g := range gs {
			v, _ := g.Generate()
			vs = append(vs, v)
		}
		w.Write(i, vs, i == numOutputRows-1)
	}
}
