package mock_redis

import (
	"time"

	"golang.org/x/net/context"

	redisUC "books/src/infra/persistence/redis/usecase"

	"github.com/stretchr/testify/mock"
)

var _ redisUC.RedisInt = &MockRedisUC{}

type MockRedisUC struct {
	mock.Mock
}

func (o *MockRedisUC) SetData(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	args := o.Called(ctx, key, value, ttl)

	var (
		err error
	)

	if n, ok := args.Get(0).(error); ok {
		err = n
	}

	return err
}

func (o *MockRedisUC) GetData(ctx context.Context, key string) (string, error) {
	args := o.Called(ctx, key)

	var (
		res string
		err error
	)

	if n, ok := args.Get(0).(string); ok {
		res = n
	}

	if n, ok := args.Get(1).(error); ok {
		err = n
	}

	return res, err
}
