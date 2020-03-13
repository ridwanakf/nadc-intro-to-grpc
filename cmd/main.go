package main

import (
	"log"

	"github.com/ridwanakf/nadc-intro-to-grpc/internal/app"
	"github.com/ridwanakf/nadc-intro-to-grpc/internal/delivery/grpc"
)

func main() {
	// init app
	nadcGrpcApp, err := app.NewNadcGrpc()
	if err != nil {
		log.Fatalf("marshal error %+v", err)
	}
	defer func() {
		errs := nadcGrpcApp.Close()
		for e := range errs {
			log.Println(e)
		}
	}()

	grpc.Start(nadcGrpcApp)
}
