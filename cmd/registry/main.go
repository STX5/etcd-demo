package main

import (
	"etcd-demo/registry"
	"log"
)

func main() {
	var endpoints = []string{"localhost:2379"}
	ser, err := registry.NewServiceRegister(endpoints, "/web/node1", "localhost:8000", 5)
	if err != nil {
		log.Fatalln(err)
	}
	//监听续租相应chan
	go ser.ListenLeaseRespChan()
	select {
	// case <-time.After(20 * time.Second):
	// 	ser.Close()
	}
}
