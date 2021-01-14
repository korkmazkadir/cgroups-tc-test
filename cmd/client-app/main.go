package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"

	"../../app"
)

const fileName = "Test-1MB"

func main() {

	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) != 2 {
		panic(fmt.Errorf("you should provide an IP Address and a port number as arguments"))
	}

	ipAddress := argsWithoutProg[0]
	portNumber, err := strconv.Atoi(argsWithoutProg[1])
	if err != nil {
		panic(err)
	}

	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {

			client, err := app.NewClient(ipAddress, portNumber)
			if err != nil {
				panic(err)
			}

			getFileFromServer(client)
			wg.Done()
		}()
	}

	/*

		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {

				client, err := app.NewClient(ipAddress, portNumber)
				if err != nil {
					panic(err)
				}

				file := createRandomFile(1000500)
				putRandomFile(client, file)

				wg.Done()
			}()
		}

	*/

	wg.Wait()

}

func getFileFromServer(client *app.Client) {

	//rand.Seed(time.Now().UnixNano())
	//sleepTime := time.Duration(rand.Intn(2000))
	//time.Sleep(sleepTime * time.Millisecond)

	start := time.Now()
	file, err := client.GetFile(fileName)
	if err != nil {
		panic(err)
	}

	log.Printf("[GET]\t%d\t%d\t%s\n", time.Since(start).Milliseconds(), len(file.Data), file.Hash())

}

func putRandomFile(client *app.Client, file *app.File) {

	start := time.Now()
	err := client.PutFile(file)
	if err != nil {
		panic(err)
	}

	log.Printf("[PUT]\t%d\t%d\t%s\n", time.Since(start).Milliseconds(), len(file.Data), file.Hash())
}

func createRandomFile(fileSize int) *app.File {

	file := app.File{}
	file.Data = make([]byte, fileSize)

	rand.Seed(time.Now().UnixNano())
	size, err := rand.Read(file.Data)
	if err != nil {
		panic(err)
	}

	if size != fileSize {
		panic(fmt.Errorf("file size %d is not equal to expected file size %d", size, fileSize))
	}

	file.Name = file.Hash()

	return &file
}
