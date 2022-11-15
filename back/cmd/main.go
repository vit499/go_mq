package main

import (
	"back/internal/http_mq"
	"back/internal/mq_mq"
	"back/internal/unit"
	"log"
)

func main() {

	log.Printf("Starting...")

	// up := make([]*unit.Unit, 10)

	// u1 := unit.Unit{}
	// u1.Init("0802")
	// up[0] = &u1

	// u2 := unit.Unit{}
	// u2.Init("0803")
	// up[1] = &u2

	// u3 := unit.Unit{}
	// u3.Init("0804")
	// up[2] = &u3

	// h := http_mq.HttpServer{}
	// log.Fatal(h.StartHttp(":3100"))

	mq := mq_mq.Mq{}
	mq.Init("tcp://vit496.ru:2083", "ab@m.ru", "1111")
	mq.InitClient()
	if err := mq.Connect(); err != nil {
		log.Printf("mqtt connect err: %s ", err)
	}
	log.Printf("mq -- ")

	us := unit.GetUnits(&mq)

	us.AddUnit("0802")

	//u3.PrintUnit()

	h := http_mq.HttpServer{}
	log.Fatal(h.StartHttp("127.0.0.1:3100"))
}
