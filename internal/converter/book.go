package converter

import (
	"github.com/ridwanakf/nadc-intro-to-grpc/internal/entity"
	"github.com/ridwanakf/nadc-intro-to-grpc/protos"
)

func ConverterBookToProto(book entity.Book) *protos.Book{
	return 	&protos.Book{
		BookId:         book.ID,
		BookTitle:      book.Title,
		BookAuthorName: book.Author,
		BookCategory:   book.Category,
		BookRate:       book.Rate,
	}
}
