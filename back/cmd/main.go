package main

import (
	"back/internal/hglob"
	"back/internal/http_mq"
	"back/internal/mq_mq"
	"back/internal/sensor"
	"back/internal/service/sensor_service"
	"back/internal/service/units_service"
	"back/internal/unit"
	"back/internal/ws"
	"back/pkg/config"
	"back/pkg/logger"
	"back/pkg/tgbot"
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"
	//"github.com/pkg/errors"
)

func run() error {
	err := godotenv.Load()
	if err != nil {
		//log.Printf("Error loading .env file")
	}
	_ = config.Get()

	l := logger.Get()
	l.Info().Msg("Starting...")

	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		//l.Info().Msg("Stopping...")
		cancel()
		time.Sleep(50 * time.Millisecond)
		//time.Sleep(2 * time.Second)
		//l.Info().Msg("Stopped.")
	}()

	hglob := hglob.NewHglob()

	tg := tgbot.GetTgbot()
	us := unit.Get(ctx, tg, l)
	sens := sensor.NewDataSensor(tg, l)

	units_service := units_service.NewUnitsService(us, hglob)
	sensorService := sensor_service.NewSensorService(sens, l)

	err = mq_mq.Get(ctx, l, us, hglob)
	if err != nil {
		return err
	}

	hub := ws.NewHub(ctx, units_service, hglob)

	go func() {
		err := http_mq.GetHttpServer(ctx, units_service, sensorService, l, hub)
		if err != nil {
			l.Error().Msg(err.Error())
		}
	}()
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	<-sigint

	return nil
}

func main() {
	// ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	// defer stop()

	if err := run(); err != nil {
		log.Fatalf("err: %s", err.Error())
	}
}
