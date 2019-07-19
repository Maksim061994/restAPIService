package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

// URLS - структура с urls
type Config struct {
	PathToSoap string "json:\"pathToSoap\""
	UrlEtran string "json:\"urlEtran\""
	FromDate string "json:\"fromDate\""
	ToDate string "json:\"toDate\""
	PatternFind string "json:\"patternFind\""
}

// ReadDataFromJSON - функция обработки json файла
func ReadDataFromJSON(path string) Config {
	jsonFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	var conf Config
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &conf)
	return conf
}

// GetFromToDate - возвращает дату начала и окончания поиска из XML файла
func GetFromToDate(inputString, maskFromDate, maskToDate  string) (string, string) {
	// reFromDate, reToDate, reQuot - шаблоны для поиска fromDate и toDate
	reFromDate := regexp.MustCompile(fmt.Sprintf("\\&lt;%s(.*?)\\;/&gt", maskFromDate))
	reToDate := regexp.MustCompile(fmt.Sprintf("\\&lt;%s(.*?)\\;/&gt", maskToDate))
	reQuot := regexp.MustCompile("\\&quot;(.*?)\\&quot")

	// regfromDate, regToDate - получение данных на первом шаге парсинга
	regfromDate := reFromDate.FindStringSubmatch(inputString)
	regToDate := reToDate.FindStringSubmatch(inputString)
	
	// TODO: дописать обработчик ошибок
	fromDate := reQuot.FindStringSubmatch(string(regfromDate[1]))[1]
	toDate := reQuot.FindStringSubmatch(string(regToDate[1]))[1]
	return fromDate, toDate
}
