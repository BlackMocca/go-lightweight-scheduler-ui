package constants

import (
	"fmt"
	"os"
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
