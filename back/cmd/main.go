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
	"log"

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

	hglob := hglob.NewHglob()

	tg := tgbot.GetTgbot()
	us := unit.Get(tg, l)

	service := service.NewUnitsService(us, hglob)

	mq := mq_mq.Get(l, us, hglob)
	defer mq.Disconnect()

	hub := ws.NewHub(service, hglob)

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
