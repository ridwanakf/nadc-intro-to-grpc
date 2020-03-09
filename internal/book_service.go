package internal

import (
	"context"
	"database/sql"
	"log"

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

func (s *BookService) GetBookByID(ctx context.Context, req *protos.BookIDRequest) (*protos.Book, error) {
	//get BookId
	bookID := req.GetBookId()

	//prepare sql statement
	stmt, err := s.db.Prepare(SQLGetBookByID)
	if err != nil{
		log.Printf("[internal][GetBookByID] invalid prepare statement :%+v\n", err)
		return &protos.Book{}, err
	}
	defer stmt.Close()

	//execute query
	rows, err := stmt.Query(bookID)
	var book Book
	for rows.Next() {
		err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
			&book.Rate,
			&book.Category,
		)
		if err != nil {
			log.Printf("[internal][GetBookByID] fail to scan :%+v\n", err)
			continue
		}
	}

	//convert book entity to protos
	res := &protos.Book{
		BookId:               book.ID,
		BookTitle:            book.Title,
		BookAuthorName:       book.Author,
		BookCategory:         book.Category,
		BookRate:             book.Rate,
	}

	return res, nil
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
