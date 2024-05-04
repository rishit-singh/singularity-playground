package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
	"playground/database"
	"github.com/joho/godotenv"
)

type APIConfig struct {
	Username string
	Password string
	URL      string
}

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

func LoadConfig() APIConfig {
	LoadEnvs()

	variables := make([]string, 3)

	for index, val := range RequiredEnvVars {
		variable, exists := os.LookupEnv(val)

		if !exists {
			panic(fmt.Sprintf("Required environment variable %s not set.", val))
		}

		variables[index] = variable
	}

	return APIConfig{
		URL:      variables[0],
		Username: variables[1],
		Password: variables[2],
	}
}

func createSchedule(config ScheduleConfig) {
	// request := schedule.CreateRequest{
	// 	Preparation:          config.Preparation,
	// 	Provider:             config.Provider,
	// 	HTTPHeaders:          config.HTTPHeaders,
	// 	URLTemplate:          config.URLTemplate,
	// 	PricePerGBEpoch:      config.PricePerGBEpoch,
	// 	PricePerGB:           config.PricePerGB,
	// 	PricePerDeal:         config.PricePerDeal,
	// 	Verified:             config.Verified,
	// 	IPNI:                 config.IPNI,
	// 	KeepUnsealed:         config.KeepUnsealed,
	// 	ScheduleCron:         config.ScheduleCron,
	// 	StartDelay:           config.StartDelay,
	// 	Duration:             config.Duration,
	// 	ScheduleDealNumber:   config.ScheduleDealNumber,
	// 	TotalDealNumber:      config.TotalDealNumber,
	// 	ScheduleDealSize:     config.ScheduleDealSize,
	// 	TotalDealSize:        config.TotalDealSize,
	// 	Notes:                config.Notes,
	// 	MaxPendingDealSize:   config.MaxPendingDealSize,
	// 	MaxPendingDealNumber: config.MaxPendingDealNumber,
	// 	AllowedPieceCIDs:     config.AllowedPieceCIDs,
	// 	Force:                config.Force,
	// }

	// c := context.TODO()

	// lotusClient := util.NewLotusClient("lotus-apu", "lotus-token")

	// schedule, err := schedule.Default.CreateHandler(c, db, lotusClient, request)

	// if err != nil {
	// 	return errors.WithStack(err)
	// }
}

func GetPreparations(config APIConfig) (string, error) {
	request, err := http.NewRequest("GET", fmt.Sprintf("%s%s", config.URL, "preparation"), nil)

	request.Header.Add("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", config.Username, config.Password)))))

	// fmt.Println(fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", config.Username, config.Password)))))
	// fmt.Println(fmt.Sprintf("%s:%s", config.Username, config.Password))

	response, err := http.DefaultClient.Do(request)
		
	if err != nil {
		fmt.Print(err)
		return "", err
	}

	body, err := io.ReadAll(response.Body)

	// body := response.Status

	if err != nil {
		fmt.Printf("Failed to read response body: %s\n", err)

		return "", err
	}

	return string(body), nil

	// var preparations any

	// err = json.Unmarshal(body, &preparations)

	// if err != nil {
	// 	fmt.Printf("Failed to parse response body: %s\n", err)

	// 	return nil, err
	// }

	// return preparations, nil
}

func main() {
	config := LoadConfig()

	preparations, err := GetPreparations(config)

	if err != nil {
		print("Failed to get preparations")
		return
	}

	fmt.Println(preparations)
	// for _, val := range preparations {
	// 	print(val)
	// }
}
