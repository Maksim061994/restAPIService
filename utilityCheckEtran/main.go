package main

import (
	"fmt"
	"flag"
	"time"
	"strings"
	"./common"
	"./etran"
)


const layoutTime string = "02.01.2006 15:04:05"

func main() {
	// pathToConfig - путь до config задается с помощью флага -config="PATH-TO-CONF"
	pathToConfig := flag.String("config", ".\\config.json", "a string")
	flag.Parse()

	// config - current config
	config := common.ReadDataFromJSON(*pathToConfig)

	// etranBodyQuery - xml запрос для поиска изменений в заявках ЭТРАН
	etranBodyQuery := etran.GetQueryXML(config.PathToSoap)

	// maskFromDate, maskToDate - маски шаблонов для regex
	var maskFromDate string = config.FromDate
	var maskToDate string = config.ToDate

	// fromDate, toDate - получение даты началы и даты окончания сканирования
	fromDate, toDate := common.GetFromToDate(etranBodyQuery, maskFromDate, maskToDate)

	// start, end - датs началы и даты окончания после Parse
	start, err := time.Parse(layoutTime, fromDate)
	if err != nil {
		fmt.Println(err)
	}
	end, err := time.Parse(layoutTime, toDate)
	if err != nil {
		fmt.Println(err)
	}

	for dt := start; dt.Before(end); dt = dt.Add(1*time.Hour){
		dtFormat := dt.Format(layoutTime)
		fmt.Println(dtFormat)
		etranBodyQuery = strings.Replace(etranBodyQuery, fromDate, dtFormat, -1)
		// response - запрос к ЭТРАНу
		response := etran.PostSOAP(config.UrlEtran, etranBodyQuery)
		// Проверка response на наличие шаблона
		if strings.Contains(response, config.PatternFind){
			fmt.Println(strings.Repeat("-", 50))
			fmt.Println(response)
			fmt.Println(strings.Repeat("-", 50))
		}
		fromDate = dtFormat
	}

}
