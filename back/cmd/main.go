package main

import (
	"back/internal/http_mq"
	"back/internal/mq_mq"
	"back/internal/unit"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Printf("Starting...")

	mq := mq_mq.Mq{}
	defer mq.Disconnect()

	mq.InitClient(os.Getenv("MQTT_HOST"), os.Getenv("MQTT_USER"), os.Getenv("MQTT_PASS"))
	if err := mq.Connect(); err != nil {
		log.Printf("mqtt connect err: %s ", err)
	}
	log.Printf("mq -- ")

	us := unit.GetUnits(&mq)

	us.AddUnit("0802")
	us.AddUnit("0803")
	us.AddUnit("0804")

	h := http_mq.GetHttpServer(us)
	log.Fatal(h.StartHttp("127.0.0.1:3100"))

}
