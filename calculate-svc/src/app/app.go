package app

import (
	"calc/controllers"
	"net/http"
)

func StartApp() {
	http.HandleFunc("/basic_stats", controllers.GetBasicStatistics)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
