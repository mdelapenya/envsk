package main

import (
	"fmt"
	"os"
)

func main() {
	m := envVarsMap()

	for k, v := range m {
		fmt.Printf("%s=%s\n", k, v)
	}
}

func envVarsMap() map[string]string {
	envMap := make(map[string]string)

	envVars := os.Environ()
	for _, envVar := range envVars {
		varName, varValue := splitEnvVar(envVar)

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
