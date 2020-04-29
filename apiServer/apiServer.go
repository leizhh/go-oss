package main

import (
	config "YJC-OSS/config"
	heartbeat "YJC-OSS/apiServer/heartbeat"
	locate "YJC-OSS/apiServer/locate"
	objects "YJC-OSS/apiServer/objects"
	"log"
	"net/http"
)

func main() {
	go heartbeat.ListenHeartbeat()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/locate/", locate.Handler)
	log.Println("API Server is running")
	log.Fatal(http.ListenAndServe(config.LISTEN_ADDRESS, nil))
}
