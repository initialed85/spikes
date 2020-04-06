package main

import (
	"flag"
	"github.com/google/uuid"
	"github.com/koron/go-ssdp"
	"log"
	"time"
)

func server() {
	ad, err := ssdp.Advertise(
		"my:device",
		uuid.New().String(),
		"http://192.168.0.1:57086/foo.xml",
		"go-ssdp sample",
		1800)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 16; i++ {
		err = ad.Alive()
		if err != nil {
			log.Fatal(err)
		}

		time.Sleep(time.Second)
	}

	err = ad.Bye()
	if err != nil {
		log.Fatal(err)
	}

	err = ad.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func client() {
	onAlive := func(m *ssdp.AliveMessage) {
		log.Printf("onAlive: %+v", m)
	}

	onBye := func(m *ssdp.ByeMessage) {
		log.Printf("onAlive: %+v", m)
	}

	m := ssdp.Monitor{
		Alive: onAlive,
		Bye:   onBye,
	}

	err := m.Start()
	if err != nil {
		log.Fatal(err)
	}

	select {}
}

func main() {
	isServer := flag.Bool("server", false, "server mode")

	flag.Parse()

	if *isServer {
		log.Printf("server mode")
		server()
	} else {
		log.Printf("client mode")
		client()
	}
}
