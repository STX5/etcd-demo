package main

import (
	"etcd-demo/discovery"
	"log"
	"time"
)

func main() {
	var endpoints = []string{"localhost:2379"}
	ser := discovery.NewServiceDiscovery(endpoints)
	defer ser.Close()
	ser.WatchService("/web/")
	ser.WatchService("/gRPC/")
	for range time.Tick(10 * time.Second) {
		log.Println(ser.GetServices())

	}
}
