package main

import (
	"encoding/json"
	"fmt"
	"os"
	"playground/singularitycontext"

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

func LoadConfig() singularitycontext.AuthConfig {
	LoadEnvs()

	variables := make([]string, 3)

	for index, val := range RequiredEnvVars {
		variable, exists := os.LookupEnv(val)

		if !exists {
			panic(fmt.Sprintf("Required environment variable %s not set.", val))
		}

		variables[index] = variable
	}

	return singularitycontext.AuthConfig{
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

	// preparations, err := singularity.GetStorages()

	metaJson, err := os.ReadFile(os.Args[1])

	if err != nil {
		print(err)
		return
	}

	var metadata []singularitycontext.DatasetMetadata

	json.Unmarshal([]byte(metaJson), &metadata)

	response, err := singularity.CreateStorage(os.Args[2], os.Args[3], metadata[0], singularitycontext.AuthConfig{
		URL:      "",
		Username: "",
		Password: "",
	}, singularitycontext.StorageConfig{
		Concurrency: 0,
		IdleTimeout: 300,
	})

	if err != nil {
		print("Failed to get preparations")
		return
	}

	// str, _ := util.ToJson(response)

	fmt.Println(response)
}
