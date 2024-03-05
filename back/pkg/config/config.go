package config

import (
	"encoding/json"
	"log"
	"os"
	"strings"
	"sync"
)

type Config struct {
	HttpHost string
	// HttpPort string
	MqHost string
	MqUser string
	MqPass string
	// PgHost   string
	// PgPort   string
	// PgDb     string
	// PgUser   string
	// PgPass   string
	TgToken  string
	TgChat   string
	Units    []string
	MqEnable bool
}

var (
	config Config
	once   sync.Once
)

func Get() *Config {

	once.Do(func() {
		config.HttpHost = os.Getenv("HTTP_HOST")
		config.MqHost = os.Getenv("MQTT_HOST")
		config.MqUser = "**"
		config.MqPass = "**"
		// config.PgHost = os.Getenv("PG_HOST")
		// config.PgPort = os.Getenv("PG_PORT")
		// config.PgDb = os.Getenv("PG_DB")
		config.TgToken = "*"
		config.TgChat = "*"
		s := os.Getenv("UNITS")
		config.Units = strings.Split(s, ",")
		mqEn := os.Getenv("MQTT_ENABLE")
		config.MqEnable = true
		if mqEn == "0" {
			config.MqEnable = false
		}

		b, err := json.MarshalIndent(config, "", "")
		if err != nil {
			log.Printf("json cfg err: %s", err.Error())
		}
		log.Printf("cfg: %s", string(b))
		config.MqUser = os.Getenv("MQTT_USER")
		config.MqPass = os.Getenv("MQTT_PASS")
		config.TgToken = os.Getenv("TOKEN")
		config.TgChat = os.Getenv("CHAT")
	})
	return &config
}
