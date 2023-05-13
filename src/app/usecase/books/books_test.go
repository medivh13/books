package books

import (
	mockDTO "books/mocks/app/dto/books"
	mockInteg "books/mocks/infra/integration/books"
	mockRedis "books/mocks/infra/redis"
	mockReponse "books/mocks/interface/response"
	dto "books/src/app/dto/books"
	"context"
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : books
 */

type MockBooksUseCase struct {
	mock.Mock
}

type BooksUseCaseList struct {
	suite.Suite
	resp      *mockReponse.MockResponse
	mockDTO   *mockDTO.MockBooksDTO
	mockRedis *mockRedis.MockRedisUC
	useCase   BooksUCInterface
	mockInteg *mockInteg.MockInteg

	dtoGet  *dto.BookReqDTO
	dtoResp *dto.GetBooksRespDTO
}

func (suite *BooksUseCaseList) SetupTest() {
	suite.resp = new(mockReponse.MockResponse)
	suite.mockDTO = new(mockDTO.MockBooksDTO)
	suite.mockRedis = new(mockRedis.MockRedisUC)
	suite.mockInteg = new(mockInteg.MockInteg)
	suite.useCase = NewBooksUseCase(suite.mockRedis, suite.mockInteg)

	suite.dtoGet = &dto.BookReqDTO{
		Subject: "love",
	}

	suite.dtoResp = &dto.GetBooksRespDTO{
		Name: "test",
	}

}

func (u *BooksUseCaseList) TestGetBySubjectSuccess() {
	u.mockRedis.Mock.On("GetData", context.Background(), "love").Return("", errors.New(mock.Anything))
	u.mockInteg.Mock.On("GetBooksBySubject", "love").Return(mock.Anything, nil)
	u.mockRedis.Mock.On("SetData", context.Background(), "love", mock.Anything, time.Duration(2)*time.Minute).Return(nil)
	_, err := u.useCase.GetBooksBySubject(u.dtoGet)
	u.Equal(nil, err)
}

func (u *BooksUseCaseList) TestGetBySubjectFromRedisSuccess() {
	dataresp, _ := json.Marshal(u.dtoResp)
	u.mockRedis.Mock.On("GetData", context.Background(), "love").Return(string(dataresp), nil)
	_, err := u.useCase.GetBooksBySubject(u.dtoGet)
	u.Equal(nil, err)
}

func (u *BooksUseCaseList) TestGetBySubjectFail() {
	u.mockRedis.Mock.On("GetData", context.Background(), "love").Return("", errors.New(mock.Anything))
	u.mockInteg.Mock.On("GetBooksBySubject", "love").Return(nil, errors.New(mock.Anything))
	_, err := u.useCase.GetBooksBySubject(u.dtoGet)
	u.Equal(errors.New(mock.Anything), err)
}

func (u *BooksUseCaseList) TestGetBySubjectSetDataRedisFail() {
	dataresp, _ := json.Marshal(u.dtoResp)
	u.mockRedis.Mock.On("GetData", context.Background(), "love").Return("", errors.New(mock.Anything))
	u.mockInteg.Mock.On("GetBooksBySubject", "love").Return(u.dtoResp, nil)
	u.mockRedis.Mock.On("SetData", context.Background(), "love", dataresp, time.Duration(2)*time.Minute).Return(errors.New(mock.Anything))
	_, err := u.useCase.GetBooksBySubject(u.dtoGet)
	u.Equal(errors.New(mock.Anything), err)
}

func TestUsecase(t *testing.T) {
	suite.Run(t, new(BooksUseCaseList))
}
