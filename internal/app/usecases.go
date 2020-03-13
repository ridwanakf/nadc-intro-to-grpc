package app

import (
	"github.com/ridwanakf/nadc-intro-to-grpc/internal"
	"github.com/ridwanakf/nadc-intro-to-grpc/internal/usecase"
)

type Usecases struct {
	BookUC internal.BookUC //BookUC Usecase
}

func newUsecases(repos *Repos) *Usecases {
	return &Usecases{
		BookUC: usecase.NewBookUsecase(repos.bookRepo),
	}
}

func (*Usecases) Close() []error {
	var errs []error
	return errs
}
