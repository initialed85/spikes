package main

import (
	"flag"
	"fmt"
	zmq "github.com/pebbe/zmq4"

	"log"
)

func server() {
	context, err := zmq.NewContext()
	if err != nil {
		log.Fatal(err)
	}
	defer context.Term()

	socket, err := context.NewSocket(zmq.REP)
	if err != nil {
		log.Fatal(err)
	}
	defer socket.Close()

	err = socket.Bind("tcp://*:5555")
	if err != nil {
		log.Fatal(err)
	}

	for {
		msg, err := socket.Recv(0)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("received: %+v", msg)

		data := fmt.Sprintf("you said: %v", msg)
		log.Printf("sending: %v", data)
		_, err = socket.Send(data, 0)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func client() {
	context, err := zmq.NewContext()
	if err != nil {
		log.Fatal(err)
	}
	defer context.Term()

	socket, err := context.NewSocket(zmq.REQ)
	if err != nil {
		log.Fatal(err)
	}
	defer socket.Close()

	err = socket.Connect("tcp://localhost:5555")
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 1024; i++ {
		data := fmt.Sprintf("iteration %v", i)
		log.Printf("sending: %v", data)
		_, err = socket.Send(data, 0)
		if err != nil {
			log.Fatal(err)
		}

		data, err = socket.Recv(0)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("received: %v", data)
	}
}

func main() {
	isServer := flag.Bool("server", false, "server mode")

	flag.Parse()

	if *isServer {
		log.Println("server mode")
		server()
	} else {
		log.Println("client mode")
		client()
	}
}
