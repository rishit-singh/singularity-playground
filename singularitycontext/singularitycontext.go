package singularitycontext

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"playground/util"
)

type AuthConfig struct {
	Username string
	Password string
	URL      string
}

type DatasetMetadata struct {
	URL             string `json:"url"`
	MD5             string `json:"md5"`
	DataCollection  string `json:"data_collection"`
	DataType        string `json:"data_type"`
	AnalysisGroup   string `json:"analysis_group"`
	Sample          string `json:"sample"`
	Population      string `json:"population"`
	DataReusePolicy string `json:"data_reuse_policy"`
}

type StorageConfig struct {
	Concurrency uint
	IdleTimeout uint
}

type SingularityContext struct {
	Config AuthConfig
}

func (ctx *SingularityContext) GetStorages() ([]any, error) {
	request, err := http.NewRequest("GET", fmt.Sprintf("%s%s", ctx.Config.URL, "storage"), nil)

	authRequest := util.NewBasicAuthRequest(request, util.Credentials{Username: ctx.Config.Username, Password: ctx.Config.Password})

	response, err := http.DefaultClient.Do(authRequest.Request)

	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		fmt.Printf("Failed to read response body: %s\n", err)

		return nil, err
	}

	var storages []any

	err = json.Unmarshal(body, &storages)

	if err != nil {
		fmt.Printf("Failed to parse response body: %s\n", err)

		return nil, err
	}

	return storages, nil
}

func (ctx *SingularityContext) GetPreparations() ([]any, error) {
	request, err := http.NewRequest("GET", fmt.Sprintf("%s%s", ctx.Config.URL, "preparation"), nil)

	authRequest := util.NewBasicAuthRequest(request, util.Credentials{Username: ctx.Config.Username, Password: ctx.Config.Password})

	response, err := http.DefaultClient.Do(authRequest.Request)

	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		fmt.Printf("Failed to read response body: %s\n", err)

		return nil, err
	}

	var preparations []any

	err = json.Unmarshal(body, &preparations)

	if err != nil {
		fmt.Printf("Failed to parse response body: %s\n", err)

		return nil, err
	}

	return preparations, nil
}

// CreatePreparation creates a new preparation in the instance with the gien name and a list of source Storages
func (ctx *SingularityContext) CreatePreparation(name string, sourceStorages []string) (any, error) {
	options := map[string]any{}

	options["name"] = name
	options["sourceStorages"] = sourceStorages

	data, err := util.ToJson(options)

	url := fmt.Sprintf("%s%s", ctx.Config.URL, "preparation")

	request, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(data)))

	request.Header.Set("Content-Type", "application/json")
	request.SetBasicAuth(ctx.Config.Username, ctx.Config.Password)

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	bodyJson, err := io.ReadAll(response.Body)

	return string(bodyJson), nil
}

func (ctx *SingularityContext) CreateStorage(name string, path string, metadata DatasetMetadata, config AuthConfig, storageConfig StorageConfig) (any, error) {
	metadataJson, err := util.ToJson(metadata)

	println(metadataJson)

	dataMap := map[string]any{
		"name":     name,
		"path":     path,
		"metadata": metadataJson,
		"config": map[string]any{
			"host":        config.URL,
			"user":        config.Username,
			"pass":        config.Password,
			"concurrency": fmt.Sprintf("%d", storageConfig.Concurrency),
			"idleTimeout": fmt.Sprintf("%d", storageConfig.IdleTimeout),
		},
	}

	data, _ := util.ToJson(dataMap)

	println(fmt.Sprintf("%s%s", ctx.Config.URL, "storage/ftp"))

	request, err := http.NewRequest("POST", fmt.Sprintf("%s%s", ctx.Config.URL, "storage/ftp"), bytes.NewBuffer([]byte(data)))

	request.Header.Set("Content-Type", "application/json")
	request.SetBasicAuth(ctx.Config.Username, ctx.Config.Password)

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	fmt.Println(response.Status)

	body, err := io.ReadAll(response.Body)

	if err != nil {
		fmt.Printf("Failed to read response body: %s\n", err)

		return nil, err
	}

	var responseObject map[string]any

	err = json.Unmarshal(body, &responseObject)

	if err != nil {
		fmt.Printf("Failed to parse response body: %s\n", err)

		return nil, err
	}

	return responseObject, nil
}

// func NewSingularityContext() *SingularityContext {
// 	return
// }
