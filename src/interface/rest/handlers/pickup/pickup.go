package article

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : books
 */

import (
	"encoding/json"
	"net/http"

	dto "books/src/app/dto/pickup"
	usecases "books/src/app/usecase/pickup"
	common_error "books/src/infra/errors"
	"books/src/interface/rest/response"
)

type BooksHandlerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetDataPickUp(w http.ResponseWriter, r *http.Request)
}

type booksHandler struct {
	response response.IResponseClient
	usecase  usecases.PickUpUCInterface
}

func NewBooksHandler(r response.IResponseClient, h usecases.PickUpUCInterface) BooksHandlerInterface {
	return &booksHandler{
		response: r,
		usecase:  h,
	}
}

func (h *booksHandler) Create(w http.ResponseWriter, r *http.Request) {

	postDTO := dto.ReqPickupDTO{}
	err := json.NewDecoder(r.Body).Decode(&postDTO)
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}
	err = postDTO.Validate()
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}

	err = h.usecase.Create(&postDTO)
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.FAILED_CREATE_DATA, err))
		return
	}

	h.response.JSON(
		w,
		"Successful Adding New PickUp Schedule",
		nil,
		nil,
	)
}

func (h *booksHandler) GetDataPickUp(w http.ResponseWriter, r *http.Request) {

	data := h.usecase.GetDataPickUp()

	h.response.JSON(
		w,
		"Successful Get Books",
		data,
		nil,
	)
}
