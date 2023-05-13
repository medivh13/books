package books

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : books
 */

import (
	dto "books/src/app/dto/books"
	"context"
	"encoding/json"
	"log"
	"time"

	integ "books/src/infra/integration/books"
	redis "books/src/infra/persistence/redis/usecase"
)

type BooksUCInterface interface {
	GetBooksBySubject(req *dto.BookReqDTO) (*dto.GetBooksRespDTO, error)
}

type booksUseCase struct {
	BooksRedis redis.RedisInt
	BooksInteg integ.OpenLibraryServices
}

func NewBooksUseCase(r redis.RedisInt, i integ.OpenLibraryServices) *booksUseCase {
	return &booksUseCase{
		BooksRedis: r,
		BooksInteg: i,
	}
}

func (uc *booksUseCase) GetBooksBySubject(req *dto.BookReqDTO) (*dto.GetBooksRespDTO, error) {

	var resp *dto.GetBooksRespDTO

	dataRedis, err := uc.BooksRedis.GetData(context.Background(), req.Subject)
	if err != nil {
		log.Printf("unable to GET data from redis. error: %v", err)
	}

	if dataRedis != "" {
		// get data from redis if is there
		_ = json.Unmarshal([]byte(dataRedis), &resp)

		log.Println("data from redis")

	} else {
		resp, err = uc.BooksInteg.GetBooksBySubject(req.Subject)

		if err != nil {
			log.Println(err)
			return nil, err
		}

		redisData, _ := json.Marshal(resp)
		ttl := time.Duration(2) * time.Minute

		// set data to redis
		rds := uc.BooksRedis.SetData(context.Background(), req.Subject, redisData, ttl)
		if err := rds; err != nil {
			log.Printf("unable to SET data. error: %v", err)
			return nil, err
		}
	}

	return resp, nil
}
