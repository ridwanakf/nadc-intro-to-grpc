package grpc

import (
	"log"

	"github.com/ridwanakf/nadc-intro-to-grpc/internal/app"
	"github.com/ridwanakf/nadc-intro-to-grpc/internal/delivery/grpc/service"
	"github.com/ridwanakf/nadc-intro-to-grpc/protos"
)

func Start(app *app.NadcGrpc) {
	svc := service.GetServices(app)

	grpcServer := NewServer(app.Cfg.GRPC.Address)

	protos.RegisterBookServiceServer(grpcServer.Server(), svc)

	if err := grpcServer.Run(); err != nil {
		log.Fatal(err)
	}
}
