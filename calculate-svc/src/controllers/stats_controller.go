package controllers

import (
	"calc/services"
	"encoding/json"
	"net/http"
	"strconv"
)

func GetBasicStatistics(resp http.ResponseWriter, req *http.Request) {
	start_timestamp, err := strconv.ParseInt(req.URL.Query().Get("from"), 10, 64)
	if err != nil {
		// TODO: Return proper Http error
		panic(err)
	}
	end_timestamp, err := strconv.ParseInt(req.URL.Query().Get("to"), 10, 64)
	if err != nil {
		// TODO: Return proper Http error
		panic(err)
	}
	stats, err := services.StatisticsService.GetBasicStatistics(start_timestamp, end_timestamp)
	if err != nil {
		// TODO: Return proper Http error
		panic(err)
	}
	jsonValue, _ := json.Marshal(stats)
	resp.Write(jsonValue)
}
