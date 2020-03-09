package main

import (
	"net"

	"github.com/ridwanakf/nadc-intro-to-grpc/internal"
	"github.com/ridwanakf/nadc-intro-to-grpc/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)


func main() {
	listener, err := net.Listen("tcp", "localhost:4040")
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()

	bookSvc := internal.BookService{}
	err = bookSvc.InitDB()

	if err != nil{
		panic(err)
	}

	protos.RegisterBookServiceServer(srv, &bookSvc)

	reflection.Register(srv)

	if err = srv.Serve(listener); err != nil {
		panic(err)
	}
}
