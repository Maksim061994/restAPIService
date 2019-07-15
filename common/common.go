package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// URLS - структура с urls
type URLS struct {
	EtranURL string "json:\"urlEtran\""
}

// ReadDataFromJSON - функция обработки json файла
func ReadDataFromJSON(path string) URLS {
	jsonFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	var urls URLS
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &urls)
	return urls
}
