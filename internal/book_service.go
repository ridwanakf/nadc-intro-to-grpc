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

func (s *BookService) InitDB() error {
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
	// get bookID
	bookID := req.GetBookId()

	// prepare sql statement
	stmt, err := s.db.Prepare(SQLGetBookByID)
	if err != nil {
		log.Printf("[internal][GetBookByID] invalid prepare statement :%+v\n", err)
		return &protos.Book{}, err
	}
	defer stmt.Close()

	// execute query
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

	// convert book entity to protos
	res := &protos.Book{
		BookId:         book.ID,
		BookTitle:      book.Title,
		BookAuthorName: book.Author,
		BookCategory:   book.Category,
		BookRate:       book.Rate,
	}

	return res, nil
}

func (s *BookService) InsertNewBook(ctx context.Context, req *protos.Book) (*protos.BoolResponse, error) {
	// get Book
	book := Book{
		ID:       req.GetBookId(),
		Title:    req.GetBookTitle(),
		Author:   req.GetBookAuthorName(),
		Category: req.GetBookCategory(),
		Rate:     req.GetBookRate(),
	}

	// prepare sql statement
	stmt, err := s.db.Prepare(SQLInsertNewBook)
	if err != nil {
		log.Printf("[internal][InsertNewBook] invalid prepare statement :%+v\n", err)
		return &protos.BoolResponse{
			Response: false,
		}, err
	}
	defer stmt.Close()

	// execute query
	_, err = stmt.Query(book.Title, book.Author, book.Rate, book.Category)
	if err != nil {
		log.Printf("[internal][InsertNewBook] error inserting new book :%+v\n", err)
		return &protos.BoolResponse{
			Response: false,
		}, err
	}

	return &protos.BoolResponse{
		Response: true,
	}, nil
}

func (s *BookService) UpdateBookRating(ctx context.Context, req *protos.BookRateRequest) (*protos.BoolResponse, error) {
	// get bookID and bookRate
	bookID := req.GetBookId()
	bookRate := req.GetBookRate()

	// prepare sql statement
	stmt, err := s.db.Prepare(SQLUpdateBookRating)
	if err != nil {
		log.Printf("[internal][UpdateBookRating] invalid prepare statement :%+v\n", err)
		return &protos.BoolResponse{
			Response: false,
		}, err
	}
	defer stmt.Close()

	// execute query
	_, err = stmt.Query(bookRate, bookID)
	if err != nil {
		log.Printf("[internal][UpdateBookRating] error updating rating :%+v\n", err)
		return &protos.BoolResponse{
			Response: false,
		}, err
	}

	return &protos.BoolResponse{
		Response: true,
	}, nil
}

func (s *BookService) DeleteBookByID(ctx context.Context, req *protos.BookIDRequest) (*protos.BoolResponse, error) {
	// get bookID
	bookID := req.GetBookId()

	// prepare sql statement
	stmt, err := s.db.Prepare(SQLDeleteBookByID)
	if err != nil {
		log.Printf("[internal][DeleteBookByID] invalid prepare statement :%+v\n", err)
		return &protos.BoolResponse{
			Response: false,
		}, err
	}
	defer stmt.Close()

	// execute query
	_, err = stmt.Query(bookID)
	if err != nil {
		log.Printf("[internal][DeleteBookByID] error inserting new book :%+v\n", err)
		return &protos.BoolResponse{
			Response: false,
		}, err
	}

	return &protos.BoolResponse{
		Response: true,
	}, err
}

func (s *BookService) SearchBookByName(ctx context.Context, req *protos.BookSlugRequest) (*protos.BooksResponse, error) {
	// get bookSlug
	bookSlug := req.GetBookName()

	// prepare sql statement
	stmt, err := s.db.Prepare(SQLSearchBookByName)
	if err != nil {
		log.Printf("[internal][SearchBookByName] invalid prepare statement :%+v\n", err)
		return &protos.BooksResponse{}, err
	}
	defer stmt.Close()

	// execute query
	rows, err := stmt.Query("%" + bookSlug + "%")
	var books []Book
	for rows.Next() {
		book := Book{}
		err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
			&book.Rate,
			&book.Category,
		)
		if err != nil {
			log.Printf("[internal][SearchBookByName] fail to scan :%+v\n", err)
			continue
		}
		books = append(books, book)
	}

	// convert all books entity to protos
	booksProto := make([]*protos.Book, len(books))
	for index := range books {
		booksProto[index] = &protos.Book{
			BookId:         books[index].ID,
			BookTitle:      books[index].Title,
			BookAuthorName: books[index].Author,
			BookCategory:   books[index].Category,
			BookRate:       books[index].Rate,
		}
	}

	return &protos.BooksResponse{
		Books: booksProto,
	}, nil
}
