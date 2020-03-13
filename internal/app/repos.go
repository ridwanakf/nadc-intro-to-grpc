package app

import (
	"database/sql"

	"github.com/ridwanakf/nadc-intro-to-grpc/internal"
	db2 "github.com/ridwanakf/nadc-intro-to-grpc/internal/repo/db"
)

type Repos struct {
	bookRepo internal.BookRepo
}

func newRepos(db *sql.DB) (*Repos, error) {
	r := &Repos{
		bookRepo: db2.NewBookDB(db),
	}

	return r, nil
}

func (r *Repos) Close() []error {
	var errs []error
	return errs
}
