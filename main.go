package main

import (
	"fmt"
	"os"
	"playground/singularitycontext"
	"playground/util"

	"github.com/joho/godotenv"
)

var EnvLoaded bool = false

var RequiredEnvVars []string = []string{
	"SG_URL",
	"SG_USERNAME",
	"SG_PASSWORD",
}

func LoadEnvs() {
	if !EnvLoaded {
		err := godotenv.Load()
		if err != nil {
			panic("Failed to load environment variables")
		}

		EnvLoaded = true
	}
}

func LoadConfig() singularitycontext.APIConfig {
	LoadEnvs()

	variables := make([]string, 3)

	for index, val := range RequiredEnvVars {
		variable, exists := os.LookupEnv(val)

		if !exists {
			panic(fmt.Sprintf("Required environment variable %s not set.", val))
		}

		variables[index] = variable
	}

	return singularitycontext.APIConfig{
		URL:      variables[0],
		Username: variables[1],
		Password: variables[2],
	}
}

func main() {
	config := LoadConfig()

	singularity := singularitycontext.SingularityContext{
		Config: config,
	}

	preparations, err := singularity.GetStorages()

	if err != nil {
		print("Failed to get preparations")
		return
	}
	str, _ := util.ToJson(preparations)

	fmt.Println(str)
}
