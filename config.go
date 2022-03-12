package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	environment string
)

func init() {
	_ = godotenv.Load()
	initVar := func(k string, req bool) string {
		val, ok := os.LookupEnv(k)
		if !ok && req {
			panic(fmt.Sprintf("envvar %s must be set", k))
		}
		return val
	}

	environment = initVar("ENV", true)
}

func Development() bool {
	return environment == "development"
}
