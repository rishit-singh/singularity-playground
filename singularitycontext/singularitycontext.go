package singularitycontext

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"playground/util"
)

type APIConfig struct {
	Username string
	Password string
	URL      string
}

type DatasetMetadata struct {
	URL             string
	MD5             string
	DataCollection  string
	DataType        string
	AnalysisGroup   string
	Sample          string
	Population      string
	DataReusePolicy string
}

type SingularityContext struct {
	Config APIConfig
}

type StorageRequest struct {
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

func (ctx *SingularityContext) CreateStorage(name string, path string, metadata DatasetMetadata) (any, error) {
	data, _ := util.ToJson(metadata)

	request, err := http.NewRequest("POST", fmt.Sprintf("%s%s", ctx.Config.URL, "preparation"), bytes.NewBuffer([]byte(data)))

	request.Header.Set("Content-Type", "application/json")
	request.SetBasicAuth(ctx.Config.Username, ctx.Config.Password)

	response, err := http.DefaultClient.Do(request)

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

// func NewSingularityContext() *SingularityContext {
// 	return
// }
