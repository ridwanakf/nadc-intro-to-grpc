package internal

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/ridwanakf/nadc-intro-to-grpc/protos"
)

type BookService struct {
	db *sql.DB
}

func (s *BookService) InitDB() error{
	// Initialize SQL DB
	db, err := sql.Open("postgres", "postgres://username:password@host-address:port/database-name?sslmode=disable")
	if err != nil {
		return err
	}

	// Check if db connected
	if err = db.PingContext(context.Background()); err != nil {
		return err
	}

	s.db = db

	return nil
}

func (s *BookService) GetBookByID(context.Context, *protos.BookIDRequest) (*protos.Book, error) {
	panic("implement me")
}

func (s *BookService) InsertNewBook(context.Context, *protos.Book) (*protos.BoolResponse, error) {
	panic("implement me")
}

func (s *BookService) UpdateBookRating(context.Context, *protos.BookRateRequest) (*protos.BoolResponse, error) {
	panic("implement me")
}

func (s *BookService) DeleteBookByID(context.Context, *protos.BookIDRequest) (*protos.BoolResponse, error) {
	panic("implement me")
}

func (s *BookService) SearchBookByName(context.Context, *protos.BookSlugRequest) (*protos.Book, error) {
	panic("implement me")
}
