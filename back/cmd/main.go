package main

import (
	"back/internal/http_mq"
	"back/internal/mq_mq"
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

	tg := tgbot.GetTgbot()

	mq := mq_mq.Get(l)
	defer mq.Disconnect()

	hub := ws.NewHub()

	us := unit.Get(mq, tg, hub, l)

	_, err = http_mq.GetHttpServer(us, l, hub)
	if err != nil {
		return err
	}

	// _, err = ws.GetWsServer(l)
	// if err != nil {
	// 	return err
	// }

	return nil
}

func main() {

	if err := run(); err != nil {
		log.Fatalf("err: %s", err.Error())
	}
}
