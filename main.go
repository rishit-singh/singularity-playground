package main

import (
	"encoding/json"
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
	"FTP_USERNAME",
	"FTP_PASSWORD",
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

func LoadConfig() (singularitycontext.AuthConfig, singularitycontext.AuthConfig) {
	LoadEnvs()

	variables := make([]string, 6)

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
		}, singularitycontext.AuthConfig{
			URL:      "",
			Username: variables[3],
			Password: variables[4],
		}
}

func main() {
	config, ftpConfig := LoadConfig()

	ftpConfig.URL = os.Args[4] // os.Args[3]
	// print(ftpConfig)

	singularity := singularitycontext.SingularityContext{
		Config: config,
	}

	// response, err := singularity.GetStorages()

	metaJson, err := os.ReadFile(os.Args[1])

	if err != nil {
		print(err)
		return
	}

	var metadata []singularitycontext.DatasetMetadata

	json.Unmarshal([]byte(metaJson), &metadata)

	// fmt.Printf("Meta size: %d", len(metadata))

	response, err := singularity.CreateStorage(os.Args[2], os.Args[3], metadata[0], singularitycontext.AuthConfig{
		URL:      ftpConfig.URL,
		Username: ftpConfig.Username,
		Password: ftpConfig.Password,
	}, singularitycontext.StorageConfig{
		Concurrency: 0,
		IdleTimeout: 300,
	})

	// if err != nil {
	// 	print("Failed to get create storage")
	// 	return
	// }

	str, _ := util.ToJson(response)

	fmt.Println(str)
}
