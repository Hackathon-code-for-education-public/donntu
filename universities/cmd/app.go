package main

import (
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
	"universities/internal/app"
	"universities/internal/logger"
)

func main() {
	srv := grpc.NewServer()
	log := logger.New()
	a, cleanup, err := app.InitApp(srv, log)
	if err != nil {
		log.Error(err.Error())
		panic(err)
	}

	log.Debug("init app")
	l, err := net.Listen("tcp", ":"+a.Cfg.GRPC.Port)
	if err != nil {
		log.Error(err.Error())
		panic(err)
	}
	defer func() {
		if err1 := l.Close(); err != nil {
			log.Error(err.Error())
			panic(err1)
		}
	}()

	log.Debug("grpc server start")
	if err := srv.Serve(l); err != nil {
		log.Error(err.Error())
		panic(err)
	}

	log.Debug("waiting for shutdown signal")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case <-quit:
		cleanup()
	}
}
