package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
)

var filter *string

func init() {
	filter = flag.String("filter", "", `Regular expression to filter environment variables. E.g.: --filter=".*V.*R$"`)
}

func main() {
	flag.Parse()

	m := envVarsMap(*filter)

	for k := range m {
		fmt.Printf("%s=%s\n", k, "********")
	}
}

func envVarsMap(filter string) map[string]string {
	envMap := make(map[string]string)

	regexp, err := regexp.Compile(filter)
	if err != nil {
		panic(err)
	}

	envVars := os.Environ()
	for _, envVar := range envVars {
		varName, varValue := splitEnvVar(envVar)

		if filter != "" {
			if !regexp.MatchString(varName) {
				continue
			}

		}

		envMap[varName] = varValue
	}

	return envMap
}

func splitEnvVar(envVar string) (string, string) {
	name := ""
	value := ""

	for i, char := range envVar {
		if char == '=' {
			name = envVar[:i]
			value = envVar[i+1:]
			break
		}
	}

	return name, value
}
