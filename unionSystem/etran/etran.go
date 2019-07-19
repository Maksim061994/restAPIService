package etran

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

// PostSOAP - функция, которая выполняет POST запрос
func PostSOAP(url string, body string) string {
	httpClient := new(http.Client)
	resp, err := httpClient.Post(url, "text/xml; charset=utf-8", bytes.NewBufferString(body))
	if err != nil {
		fmt.Println("Ошибка при выполнении запроса")
		var msg string = "Ошибка при выполнении запроса"
		return msg
	}

	byteResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении данных из ЭТРАН")
	}
	return string(byteResp)
}

// GetQueryXML - функция, которая читает XML файл
func GetQueryXML(path string) string {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Ошибка чтения данных из %s", path)
	}
	return string(file)
}
