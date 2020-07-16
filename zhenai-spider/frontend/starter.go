package main

import (
	"fmt"
	"net/http"
	"zhenai-spider/frontend/controller"
)

func main() {
	http.Handle("/res/", http.StripPrefix("/res/", http.FileServer(http.Dir("./view/res"))))
	http.Handle("/", controller.NewSearchResultHandler("view/index.html", "http://10.196.102.145:9200"))

	err := http.ListenAndServe(":9527", nil)
	if err != nil {
		fmt.Println("Failed to start HTTP server, error:", err)
	}
}
