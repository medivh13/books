package usecases

import (
	bookUC "books/src/app/usecase/books"
	pickUpUC "books/src/app/usecase/pickup"
)

type AllUseCases struct {
	BookUC   bookUC.BooksUCInterface
	PickUpUC pickUpUC.PickUpUCInterface
}
