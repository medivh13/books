package article

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : books
 */

import (
	"net/http"

	dto "books/src/app/dto/books"
	usecases "books/src/app/usecase/books"
	common_error "books/src/infra/errors"
	"books/src/interface/rest/response"
)

type BooksHandlerInterface interface {
	// Create(w http.ResponseWriter, r *http.Request)
	GetBySubject(w http.ResponseWriter, r *http.Request)
}

type booksHandler struct {
	response response.IResponseClient
	usecase  usecases.BooksUCInterface
}

func NewBooksHandler(r response.IResponseClient, h usecases.BooksUCInterface) BooksHandlerInterface {
	return &booksHandler{
		response: r,
		usecase:  h,
	}
}

// func (h *booksHandler) Create(w http.ResponseWriter, r *http.Request) {

// 	postDTO := dto.ArticleReqDTO{}
// 	err := json.NewDecoder(r.Body).Decode(&postDTO)
// 	if err != nil {
// 		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
// 		return
// 	}
// 	err = postDTO.Validate()
// 	if err != nil {
// 		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
// 		return
// 	}

// 	err = h.usecase.Create(&postDTO)
// 	if err != nil {
// 		h.response.HttpError(w, common_error.NewError(common_error.FAILED_CREATE_DATA, err))
// 		return
// 	}

// 	h.response.JSON(
// 		w,
// 		"Successful Adding New Article",
// 		nil,
// 		nil,
// 	)
// }

func (h *booksHandler) GetBySubject(w http.ResponseWriter, r *http.Request) {

	getDTO := dto.BookReqDTO{}

	if r.URL.Query().Get("subject") != "" {
		getDTO.Subject = r.URL.Query().Get("subject")
	}

	err := getDTO.Validate()
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}

	data, err := h.usecase.GetBooksBySubject(&getDTO)
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.FAILED_RETRIEVE_DATA, err))
		return
	}

	h.response.JSON(
		w,
		"Successful Get Books",
		data,
		nil,
	)
}
