package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
)

// Define a service with a method
type Greeter struct{}

// Define the request and response structures 
type HelloRequest struct {
	Name string
}

type HelloResponse struct {
	Message string
}

// RPC Method (must be exported and have two arguments: request and response)
func (g *Greeter) SayHello(req HelloRequest, res *HelloResponse) error {
	if req.Name == "" {
		return errors.New("name cannot be empty")
	}
	res.Message = "Hello, " + req.Name + "!"
	return nil
}

func main() {
	// Create an RPC server and register the service 
	server := rpc.NewServer()
	greeter := new(Greeter)
	server.Register(greeter)

	// Listen for RPC calls 
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return	
	}
	fmt.Println("Server listening on port 1234")
}