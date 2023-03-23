package main

import (
	"strconv"
	"time"

	"github.com/mchmarny/ghstore/pkg/data"
	gha "github.com/sethvargo/go-githubactions"
)

const (
	resultKey = "result"
)

var (
	version = "v0.0.1-default"

	input = map[string]string{
		"data":    "data.db",
		"counter": "counter",
		"action":  "add",
		"value":   "0",
	}
)

func main() {
	// init action with the version and build time
	a := gha.WithFieldsMap(map[string]string{
		"version": version,
		"build":   time.Now().UTC().Format(time.RFC3339),
	})

	// log start and end
	a.Infof("starting action")
	defer a.Infof("action completed")

	// check inputs
	for k := range input {
		v := a.GetInput(k)
		if v == "" {
			a.Fatalf("input %s is required", k)
		}
		input[k] = v
	}

	// init data store
	dataFile := input["data"]
	s, err := data.New(dataFile)
	if err != nil {
		a.Fatalf("error initializing data: %s - %s", dataFile, err)
	}
	defer s.Close()

	// get current counter value
	dataID := input["counter"]
	dataVal, err := s.Get(dataID)
	if err != nil {
		a.Fatalf("error getting data: %s - %s", dataID, err)
	}

	// parse input value
	newVal, err := strconv.ParseInt(input["value"], 10, 64)
	if err != nil {
		a.Fatalf("error parsing value: %s", err)
	}

	// perform action
	act := input["action"]
	switch act {
	case "add":
		dataVal += newVal
	case "sub":
		dataVal -= newVal
	case "set":
		dataVal = newVal
	default:
		a.Fatalf("invalid action: %s", act)
	}

	// save result
	if err := s.Upsert(dataID, dataVal); err != nil {
		a.Fatalf("error saving data: %s - %s", dataID, err)
	}

	// get result
	storedVal, err := s.Get(dataID)
	if err != nil {
		a.Fatalf("error getting updated data: %s - %s", dataID, err)
	}

	// validate result
	if storedVal != dataVal {
		a.Fatalf("error validating data: %s - %d != %d", dataID, storedVal, dataVal)
	}

	// set output
	a.SetOutput(resultKey, strconv.FormatInt(storedVal, 10))
}
