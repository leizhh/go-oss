package main

import (
	"go-oss/dataServer/heartbeat"
	"go-oss/dataServer/locate"
	"go-oss/dataServer/objects"
	"go-oss/dataServer/temp"
	"log"
	"net/http"
	"os"
)

func main() {
	locate.CollectObjects()
	go heartbeat.StartHeartbeat()
	go locate.StartLocate()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/temp/", temp.Handler)
	log.Println("DATA Server is running")
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
