package service

import (
	"github.com/ridwanakf/nadc-intro-to-grpc/internal/app"
)

// Services list all service struct
type Services struct {
	*BookService
}

// GetServices list all grpc service
func GetServices(app *app.NadcGrpc) *Services {
	return &Services{
		BookService: NewBookService(app.UseCases.BookUC),
	}
}
