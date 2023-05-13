package usecase

import (
	"time"

	redis "github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

type solution_ch_part1Redis struct {
	Rdb *redis.Client
}

type RedisInt interface {
	SetData(ctx context.Context, key string, value interface{}, ttl time.Duration) error
	GetData(ctx context.Context, key string) (string, error)
}

func NewsRedis(rdb *redis.Client) *solution_ch_part1Redis {
	return &solution_ch_part1Redis{
		Rdb: rdb,
	}
}

func (p *solution_ch_part1Redis) SetData(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	rds := p.Rdb.Set(context.Background(), key, value, ttl)

	return rds.Err()
}

func (p *solution_ch_part1Redis) GetData(ctx context.Context, key string) (string, error) {
	dataRedis, err := p.Rdb.Get(context.Background(), key).Result()

	return dataRedis, err
}
