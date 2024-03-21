package constants

import (
	"fmt"
	"os"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

var (
	resourcePath = GetEnv("RESOURCE_PATH", "./web/resources")
)

func GetEnv(key string, defaultVal string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return defaultVal
}

/* starting path at build/web/resource/%s */
func GetResource(pathTofileString string) string {
	return fmt.Sprintf("%s/%s", resourcePath, pathTofileString)
}

/* using get resource */
func GetSVGString(pathTofileString string) string {
	path := GetResource(pathTofileString)
	if app.IsClient {
		path = strings.ReplaceAll(path, "./", "/")
		resp, err := resty.New().R().Get(path)
		if err != nil {
			panic(err)
		}
		return string(resp.Body())
	}

	bu, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(bu)
}
