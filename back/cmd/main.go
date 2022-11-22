package main

import (
	"back/internal/http_mq"
	"back/internal/mq_mq"
	"back/internal/unit"
	"back/pkg/tgbot"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file")
	}
	log.Printf("Starting...")

	mq := mq_mq.GetMq(os.Getenv("MQTT_HOST"), os.Getenv("MQTT_USER"), os.Getenv("MQTT_PASS"))
	defer mq.Disconnect()

	mq.InitClient()
	if err := mq.Connect(); err != nil {
		log.Printf("mqtt connect err: %s ", err)
	}
	log.Printf("mq -- ")

	tg := tgbot.GetTgbot(os.Getenv("TOKEN"), os.Getenv("CHAT"))

	us := unit.GetUnits(mq, tg)

	us.AddUnit("0802")
	us.AddUnit("0803")
	us.AddUnit("0804")

	h := http_mq.GetHttpServer(us)
	httpHost := os.Getenv("HTTP_HOST")
	log.Printf("Start HTTP %s ", httpHost)
	log.Fatal(h.StartHttp(httpHost))

}
