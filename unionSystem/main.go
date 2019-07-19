package main

import (
	"fmt"

	"./common"
	"./etran"
)

// UrlsCommon - json файл с urls для систем
const UrlsCommon string = "./config/url.json"

// ClaimsReq - xml запрос для поиска изменений в заявках ЭТРАН
const ClaimsReq string = "./config/etranChangesReq.xml"

var urls = common.ReadDataFromJSON(UrlsCommon)
var etranBodyQuery = etran.GetQueryXML(ClaimsReq)

func main() {
	response := etran.PostSOAP(urls.EtranURL, etranBodyQuery)
	// var xmlData = []byte(response)

	// // KYCValue struct

	// type createBody struct {
	// 	Body string `xml:"Body"`
	// }

	// type CreateSoapEnvelope struct {
	// 	CreateBody createBody `xml:"Envelope"`
	// }

	// var createEnv CreateSoapEnvelope

	// err := xml.Unmarshal(xmlData, &createEnv)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	fmt.Println(response)
}
