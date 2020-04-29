package main

import (
	heartbeat "YJC-OSS/dataServer/heartbeat"
	locate "YJC-OSS/dataServer/locate"
	objects "YJC-OSS/dataServer/objects"
	"log"
	"net/http"
	config "YJC-OSS/config"
)

func main() {
	go heartbeat.StartHeartbeat()
	go locate.StartLocate()
	http.HandleFunc("/objects/", objects.Handler)
	log.Println("DATA Server is running")
	log.Fatal(http.ListenAndServe(config.LISTEN_ADDRESS, nil))
}
