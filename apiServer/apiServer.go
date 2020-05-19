package main

import (
	"go-oss/apiServer/heartbeat"
	"go-oss/apiServer/locate"
	"go-oss/apiServer/objects"
	"go-oss/apiServer/versions"
	"log"
	"html/template"
	"net/http"
	"os"
)

func main() {
	go heartbeat.ListenHeartbeat()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/locate/", locate.Handler)
	http.HandleFunc("/versions/", versions.Handler)
	http.HandleFunc("/demo", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("demo.html")
		if err != nil {
			log.Println("err:",err)
			return
		}
		t.Execute(w, nil)
	})

	log.Println("API Server is running")
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
