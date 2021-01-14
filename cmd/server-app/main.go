package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/rpc"
	"os"
	"time"

	"../../app"
)

const fileName = "Test-1MB"

func main() {

	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) == 0 {
		panic(fmt.Errorf("you should provide a port number"))
	}

	portNumber := argsWithoutProg[0]

	server := app.NewServer()
	putRandomFile(server, fileName, 1000500)

	rpc.Register(server)

	// resolves tcp address
	tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf(":%s", portNumber))
	if err != nil {
		panic(err)
	}

	// tcp network listener
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		panic(err)
	}

	log.Printf("Listening on %s \n", portNumber)

	for {
		conn, err := listener.Accept()
		if err == nil {
			go rpc.ServeConn(conn)
		}
	}

}

func putRandomFile(server *app.Server, fileName string, fileSize int) {

	file := app.File{}
	file.Name = fileName
	file.Data = make([]byte, fileSize)

	rand.Seed(time.Now().UnixNano())
	size, err := rand.Read(file.Data)
	if err != nil {
		panic(err)
	}

	if size != fileSize {
		panic(fmt.Errorf("file size %d is not equal to expected file size %d", size, fileSize))
	}

	err = server.PutFile(file, nil)

	if err != nil {
		panic(err)
	}
}
