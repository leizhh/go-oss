package main

import (
	heartbeat "go-oss/dataServer/heartbeat"
	locate "go-oss/dataServer/locate"
	objects "go-oss/dataServer/objects"
	"log"
	"net/http"
	"os"
)

func main() {
	go heartbeat.StartHeartbeat()
	go locate.StartLocate()
	http.HandleFunc("/objects/", objects.Handler)
	log.Println("DATA Server is running")
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
