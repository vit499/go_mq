package main

import (
	//"back/internal/http_mq"
	//"back/internal/mq_mq"
	"back/internal/unit"
	"log"
)


func main() {
	log.Printf("Starting...")

	up := make([]*unit.Unit, 10)

	u1 := unit.Unit{}
	u1.Init("0802")
	up[0] = &u1

	u2 := unit.Unit{}
	u2.Init("0803")
	up[1] = &u2

	u3 := unit.Unit{}
	u3.Init("0804")
	up[2] = &u3

	//h := http_mq.HttpServer{}
	//log.Fatal(h.StartHttp(":3100"))

	// mq := mq_mq.Mq{}
	// mq.Init(up, 3)
	// mq.Start()

	u3.PrintUnit()

}