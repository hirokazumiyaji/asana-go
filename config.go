package asana

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/naoya/go-pit"
	"os"
)

const (
	apiVersion  = "1.0"
	apiEndpoint = "https://app.asana.com/api"
)

func GetEndpoint() string {
	return fmt.Sprintf("%s/%s", apiEndpoint, apiVersion)
}

func GetApiKey() (string, error) {
	apiKey := os.Getenv("ASANA_API_KEY")
	if apiKey == "" {
		conf, err := pit.Get("app.asana.com")
		if err == nil {
			apiKey = conf["apikey"]
		}
	}

	if apiKey == "" {
		return "", errors.New("Not Found API Key. Please Set ENV(ASANA_API_KEY) or Pit(app.asana.com:apikey).")
	}

	return base64.StdEncoding.EncodeToString([]byte(apiKey + ":")), nil
}
