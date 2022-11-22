package tgbot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Сделать канал публичным. После получения id можно сделать частным.
// Ввести в браузере
// https://api.telegram.org/botBOT:TOKEN/getChat?chat_id=@имяканала
// https://api.telegram.org/botXXXX/getChat?chat_id=@xxxgroup

type Tgbot struct {
	token string     // := os.Getenv("TOKEN")
	chat_id string   // := os.Getenv("CHAT")
}

func GetTgbot(token string, chat_id string) *Tgbot {
	tgbot := Tgbot{token: token, chat_id: chat_id}

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
	// token := os.Getenv("TOKEN")
	// chat_id := os.Getenv("CHAT")
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", tg.token)

	body := cc1(tg.chat_id, src)
	//body := cc2(tg.chat_id, src)
	//body := cc3(tg.chat_id, src)
  log.Printf("send to tg: %s ", src)
	_, err := http.Post(url, "application/json", body)
	if err != nil {
		log.Printf("send bot err: %s", err.Error())
	}
	//log.Printf("re: %v", res)
}
