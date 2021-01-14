package app

import (
	"fmt"
	"net/rpc"
)

// Client implemnt an app client
type Client struct {
	netAddress string
	rpcClient  *rpc.Client
}

const serviceGetFile = "Server.GetFile"
const servicePutFile = "Server.PutFile"

// NewClient creates a client for the app
func NewClient(IPAddress string, portNumber int) (*Client, error) {

	client := Client{}
	client.netAddress = fmt.Sprintf("%s:%d", IPAddress, portNumber)

	rpcClient, err := rpc.Dial("tcp", client.netAddress)
	if err != nil {
		return nil, err
	}

	client.rpcClient = rpcClient
	return &client, nil
}

// GetFile gets the file with name from the application server
func (c *Client) GetFile(fileName string) (*File, error) {

	file := new(File)
	err := c.rpcClient.Call(serviceGetFile, fileName, file)
	if err != nil {
		return nil, err
	}

	return file, nil
}

// PutFile puts a new file
func (c *Client) PutFile(file *File) error {

	err := c.rpcClient.Call(servicePutFile, file, nil)
	if err != nil {
		return err
	}

	return nil
}

// Close closes the connection to the server
func (c *Client) Close() error {
	return c.rpcClient.Close()
}
