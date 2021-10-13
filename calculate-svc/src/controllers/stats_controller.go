package controllers

import "net/http"

func GetBasicStatistics(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("foo"))
}
