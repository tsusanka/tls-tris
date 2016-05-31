//+build ignore

package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/FiloSottile/tls-tris"
)

func main() {
	cert, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
	if err != nil {
		log.Fatal(err)
	}

	l, err := tls.Listen("tcp", os.Args[1], &tls.Config{
		MinVersion:   tls.VersionTLS13,
		MaxVersion:   tls.VersionTLS13,
		Certificates: []tls.Certificate{cert},
	})
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		if _, err := conn.Write([]byte("Do you want to play a game?")); err != nil {
			log.Println(err)
		}
		if answer, err := ioutil.ReadAll(conn); err != nil {
			log.Println(err)
		} else {
			log.Printf("%q\n", answer)
		}
		if err := conn.Close(); err != nil {
			log.Println(err)
		}
	}
}
