package test

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var rootDir string

var separator string

var jsonData map[string]any

const jsonFileName = "dir.json"

func loadJson() {
	separator = string(filepath.Separator)
	fmt.Println(separator)
	workDir, _ := os.Getwd()
	fmt.Println(workDir)
	rootDir = workDir[:strings.LastIndex(workDir, separator)]
	fmt.Println(rootDir)

	jsonBytes, _ := os.ReadFile(workDir + separator + jsonFileName)

	err := json.Unmarshal(jsonBytes, &jsonData)
	if err != nil {
		panic("Load Json Data Error: " + err.Error())
	}

}
func parseMap(mapData map[string]any, parentDir string) {
	for k, v := range mapData {
		switch v.(type) {
		case string:
			{
				path, _ := v.(string)
				if path == "" {
					continue
				}
				if parentDir != "" {
					path = parentDir + separator + path
					if k == "text" {
						parentDir = path
					}
				} else {
					parentDir = path
				}
				createDir(path)
			}
		case []any:
			{
				parseArray(v.([]any), parentDir)
			}
		}

	}
}

func parseArray(jsonData []any, parentDir string) {
	for _, v := range jsonData {
		mapV, _ := v.(map[string]any)
		parseMap(mapV, parentDir)
	}
}

func createDir(path string) {
	if path == "" {
		return
	}
	fmt.Println(path)
	err := os.MkdirAll(rootDir+separator+path, fs.ModePerm)
	if err != nil {
		panic("Create Dir Error: " + err.Error())
	}
}

func TestGenerateDir01(t *testing.T) {
	fmt.Println("testSet")
	fmt.Println("-------------------------")
	loadJson()
	parseMap(jsonData, "")
}
