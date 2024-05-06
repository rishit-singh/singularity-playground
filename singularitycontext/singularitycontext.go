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
