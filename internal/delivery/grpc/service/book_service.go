package service

import (
	"context"
	"github.com/ridwanakf/nadc-intro-to-grpc/internal"
	"github.com/ridwanakf/nadc-intro-to-grpc/internal/converter"
	"github.com/ridwanakf/nadc-intro-to-grpc/internal/entity"
	"github.com/ridwanakf/nadc-intro-to-grpc/protos"
)

type BookService struct {
	bookUC internal.BookUC
}

func NewBookService(bookUC internal.BookUC) *BookService {
	return &BookService{
		bookUC: bookUC,
	}
}

func (s *BookService) GetBookByID(ctx context.Context, req *protos.BookIDRequest) (*protos.Book, error) {
	bookID := req.GetBookId()

	res, err := s.bookUC.GetBookByID(bookID)
	if err != nil {
		return &protos.Book{}, err
	}

	return converter.ConverterBookToProto(res), nil
}

func (s *BookService) InsertNewBook(ctx context.Context, req *protos.Book) (*protos.BoolResponse, error) {
	book := entity.Book{
		ID:       req.GetBookId(),
		Title:    req.GetBookTitle(),
		Author:   req.GetBookAuthorName(),
		Category: req.GetBookCategory(),
		Rate:     req.GetBookRate(),
	}

	res, err := s.bookUC.InsertNewBook(book)
	if err != nil || !res {
		return &protos.BoolResponse{
			Response: res,
		}, err
	}

	return &protos.BoolResponse{
		Response: res,
	}, nil
}

func (s *BookService) UpdateBookRating(ctx context.Context, req *protos.BookRateRequest) (*protos.BoolResponse, error) {
	bookID := req.GetBookId()
	bookRate := req.GetBookRate()

	res, err := s.bookUC.UpdateBookRating(bookID, bookRate)
	if err != nil || !res {
		return &protos.BoolResponse{
			Response: res,
		}, err
	}

	return &protos.BoolResponse{
		Response: res,
	}, nil
}

func (s *BookService) DeleteBookByID(ctx context.Context, req *protos.BookIDRequest) (*protos.BoolResponse, error) {
	bookID := req.GetBookId()

	res, err := s.bookUC.DeleteBookByID(bookID)
	if err != nil || !res {
		return &protos.BoolResponse{
			Response: res,
		}, err
	}

	return &protos.BoolResponse{
		Response: res,
	}, nil
}

func (s *BookService) SearchBookByName(ctx context.Context, req *protos.BookSlugRequest) (*protos.BooksResponse, error) {
	bookSlug := req.GetBookName()

	res, err := s.bookUC.SearchBookByName(bookSlug)
	if err != nil {
		return &protos.BooksResponse{}, nil
	}

	booksProto := make([]*protos.Book, len(res))
	for index := range res {
		booksProto[index] = converter.ConverterBookToProto(res[index])
	}

	return &protos.BooksResponse{
		Books: booksProto,
	}, nil
}
