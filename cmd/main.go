package main

import (
	"archi/config"
	"archi/logger"
	"archi/service"
	"archi/storage"
	"archi/transport"
	"archi/transport/handlers"
	"context"
	"fmt"
	"go.uber.org/zap"
	"log"
	"os"
	"os/signal"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {

	conf := config.NewConfig()

	l, err := logger.Init(conf)
	if err != nil {
		return fmt.Errorf("cannot init logger: %w", err)
	}
	defer func(l *zap.Logger) {
		err = l.Sync()
		if err != nil {
			log.Fatalln(err)
		}
	}(l)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	gracefulShutdown(cancel, l)

	storage, err := storage.NewStorage(l, ctx, conf)
	if err != nil {
		l.Info("Error storage initialization")
		return err
	}
	l.Info("Storage initialized")

	service, err := service.NewService(l, storage)
	if err != nil {
		l.Info("Error service initialization")
		return err
	}
	l.Info("Service initialized")

	//mid := middleware.NewJWTAuth(conf)
	//l.Info("Middleware initialized")

	handler := handlers.NewHandlers(l, conf, storage, service)
	l.Info("Handler initialized")

	l.Info("Start server")
	srv := transport.NewServer(handler, conf)
	err = srv.Run(ctx)
	if err != nil {
		l.Info("Error with running server")
		return err
	}

	return nil
}

func gracefulShutdown(ctx context.CancelFunc, l *zap.Logger) {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)
	go func() {
		log.Println(<-done)
		l.Info("Gracefully shutdown")
		ctx()
	}()
}
