package main

import (
	"back/internal/hglob"
	"back/internal/http_mq"
	"back/internal/mq_mq"
	"back/internal/service/units_service"
	"back/internal/unit"
	"back/internal/ws"
	"back/pkg/config"
	"back/pkg/logger"
	"back/pkg/tgbot"
	"context"
	"log"
	"time"

	"github.com/joho/godotenv"
	//"github.com/pkg/errors"
)

func run() error {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file")
	}
	_ = config.Get()

	l := logger.Get()
	l.Info().Msg("Starting...")

	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		l.Info().Msg("Stopping...")
		cancel()
		time.Sleep(10 * time.Millisecond)
		l.Info().Msg("Stopped.")
	}()

	hglob := hglob.NewHglob()

	tg := tgbot.GetTgbot()
	us := unit.Get(ctx, tg, l)

	service := service.NewUnitsService(us, hglob)

	mq := mq_mq.Get(ctx, l, us, hglob)
	defer mq.Disconnect()

	hub := ws.NewHub(ctx, service, hglob)

	_, err = http_mq.GetHttpServer(service, l, hub)
	if err != nil {
		return err
	}

	return nil
}

func main() {

	if err := run(); err != nil {
		log.Fatalf("err: %s", err.Error())
	}
}
