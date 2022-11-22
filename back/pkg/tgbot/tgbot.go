package tgbot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// Сделать канал публичным. После получения id можно сделать частным.
// Ввести в браузере
// https://api.telegram.org/botBOT:TOKEN/getChat?chat_id=@имяканала
// https://api.telegram.org/botXXXX/getChat?chat_id=@xxxgroup

type Tgbot struct {
}

func getTgbot() *Tgbot {
	tgbot := Tgbot{}

	return &tgbot
}

// cc1, cc2, cc3 три способа получить Reader (тело запроса) в виде json,
// чтобы передать его в Post запрос
func cc1(chat_id string, src string) io.Reader {
	type jsonBot struct {
		Chat string `json:"chat_id"`
		Text string `json:"text"`
	}
	jBot := jsonBot{Chat: chat_id, Text: src}
	j1, _ := json.Marshal(jBot) // j1 []byte
	body := bytes.NewBuffer(j1) // Reader
	return body
}
func cc2(chat_id string, src string) io.Reader {
	jBot := map[string]string{"chat_id": chat_id, "text": src}
	j1, _ := json.Marshal(jBot)
	body := bytes.NewBuffer(j1)
	return body
}
func cc3(chat_id string, src string) io.Reader {
	strJson := fmt.Sprintf("{\"chat_id\":\"%s\",\"text\":\"%s\"}", chat_id, src)
	//log.Printf("strJson: %v", strJson)
	j1 := []byte(strJson)
	body := bytes.NewReader(j1)
	return body
}

// для отправки в бот нужно сформировать url и body для POST запроса
// в url подставить токен бота
// в body подставить chat_id и text
func (tg *Tgbot) SendMes(src string) {
	token := os.Getenv("TOKEN")
	chat_id := os.Getenv("CHAT")
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)

	body := cc1(chat_id, src)
	//body := cc2(chat_id, src)
	//body := cc3(chat_id, src)

	_, err := http.Post(url, "application/json", body)
	if err != nil {
		log.Printf("send bot err: %s", err.Error())
	}
	//log.Printf("re: %v", res)
}
