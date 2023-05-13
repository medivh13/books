package main

import (
	"context"

	usecases "books/src/app/usecase"
	bookUC "books/src/app/usecase/books"
	pickUpUC "books/src/app/usecase/pickup"
	"books/src/infra/config"

	"books/src/interface/rest"

	bookInteg "books/src/infra/integration/books"

	ms_log "books/src/infra/log"

	"books/src/infra/broker/nats"
	pickUpNats "books/src/infra/broker/nats/consumer/pickup"
	natsPublisher "books/src/infra/broker/nats/publisher"
	redis "books/src/infra/persistence/redis"
	redisUC "books/src/infra/persistence/redis/usecase"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// init context
	ctx := context.Background()

	// read the server environment variables
	conf := config.Make()

	// check is in production mode
	isProd := false
	if conf.App.Environment == "PRODUCTION" {
		isProd = true
	}

	// logger setup
	m := make(map[string]interface{})
	m["env"] = conf.App.Environment
	m["service"] = conf.App.Name
	logger := ms_log.NewLogInstance(
		ms_log.LogName(conf.Log.Name),
		ms_log.IsProduction(isProd),
		ms_log.LogAdditionalFields(m))

	redisClient, err := redis.NewRedisClient(conf.Redis, logger)
	bookIntegration := bookInteg.NewIntegOpenLibrary()

	bookRedis := redisUC.NewsRedis(redisClient)
	Nats := nats.NewNats(conf.Nats, logger)
	publisher := natsPublisher.NewPushWorker(Nats)

	// HTTP Handler
	// the server already implements a graceful shutdown.

	allUC := usecases.AllUseCases{
		BookUC:   bookUC.NewBooksUseCase(bookRedis, bookIntegration),
		PickUpUC: pickUpUC.NewPickUpUseCase(publisher),
	}

	pickUpNats.NewPickUpWorker(Nats, allUC.PickUpUC)

	httpServer, err := rest.New(
		conf.Http,
		isProd,
		logger,
		allUC,
	)
	if err != nil {
		panic(err)
	}
	httpServer.Start(ctx)

}
