package main

import (
	"fmt"
	"net/http"
	"zhenai-spider/frontend/controller"
)

func main() {

	http.Handle("/search", controller.NewSearchResultHandler("view/index.html", "http://10.196.102.145:9200"))

	err := http.ListenAndServe(":9300", nil)
	if err != nil {
		fmt.Println("Failed to start HTTP server, error:", err)
	}
}
