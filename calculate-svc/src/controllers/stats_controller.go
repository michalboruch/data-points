package controllers

import (
	"calc/services"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func GetBasicStatistics(resp http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")
	start_timestamp, err := strconv.ParseInt(req.URL.Query().Get("from"), 10, 64)
	if err != nil {
		// TODO: Return proper Http error
		log.Fatal("Error while parsing querystring. err: ", err)
	}
	end_timestamp, err := strconv.ParseInt(req.URL.Query().Get("to"), 10, 64)
	if err != nil {
		// TODO: Return proper Http error
		log.Fatal("Error while parsing querystring. err: ", err)
	}
	stats, err := services.StatisticsService.GetBasicStatistics(name, start_timestamp, end_timestamp)
	if err != nil {
		// TODO: Return proper Http error
		panic(err)
	}
	jsonValue, _ := json.Marshal(stats)
	resp.Write(jsonValue)
}
