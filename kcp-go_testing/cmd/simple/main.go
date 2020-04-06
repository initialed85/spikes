package main

import (
	"crypto/sha1"
	"flag"
	"fmt"
	"github.com/xtaci/kcp-go"
	"golang.org/x/crypto/pbkdf2"
	"io"
	"log"
	"time"
)

const (
	password = "P@$$w0rd123!@#"
	salt     = "$@l78@3"
)

var (
	key      = pbkdf2.Key([]byte(password), []byte(salt), 1024, 32, sha1.New)
	block, _ = kcp.NewAESBlockCrypt(key)
)

func server() {
	handle := func(conn *kcp.UDPSession) {
		for {
			buf := make([]byte, 4096)

			_, err := conn.Read(buf)
			if err != nil {
				log.Print(err)
				return
			}
			log.Printf("received: %v\n", string(buf))

			data := fmt.Sprintf("you said: %v", string(buf))
			log.Printf("sending: %v\n", data)
			_, err = conn.Write([]byte(data))
			if err != nil {
				log.Print(err)
				return
			}
		}
	}

	listener, err := kcp.ListenWithOptions("0.0.0.0:12345", block, 10, 3)
	if err != nil {
		log.Fatal(err)
	}

	for {
		s, err := listener.AcceptKCP()
		if err != nil {
			log.Fatal(err)
		}

		go handle(s)
	}
}

func client() {
	sess, err := kcp.DialWithOptions("0.0.0.0:12345", block, 10, 3)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 4096; i++ {
		data := fmt.Sprintf("the time is %v", time.Now().String())
		log.Printf("sending: %v\n", data)
		_, err := sess.Write([]byte(data))
		if err != nil {
			log.Fatal(err)
		}

		buf := make([]byte, 4096)
		_, err = io.ReadFull(sess, buf)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("received: %v\n", string(buf))
	}
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
