package main

import (
	"time"

	"github.com/mchmarny/ghstore/pkg/calc"
	gha "github.com/sethvargo/go-githubactions"
)

var (
	version = "v0.0.1-default"
)

func main() {
	a := gha.WithFieldsMap(map[string]string{
		"version": version,
		"run_at":  time.Now().UTC().Format(time.RFC3339),
	})

	a.Infof("starting action")
	defer a.Infof("action completed")

	input := calc.GetArgs()
	for k := range input {
		v := a.GetInput(k)
		a.Debugf("input %s: %s", k, v)
		if v == "" {
			a.Fatalf("input %s is required", k)
		}
		input[k] = v
	}

	results, err := calc.Calculate(input)
	if err != nil {
		a.Fatalf("error calculating: %v", err)
	}

	for k, v := range results {
		a.Debugf("output %s: %s", k, v)
		a.SetOutput(k, v)
	}
}
