package main

import (
	"fmt"
	_playerRepository "game_server/player/player_repository/player_postgres"
	"log"
	"net/url"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		viper.GetString("db.user"),
		viper.GetString("db.pass"),
		viper.GetString("db.host"),
		viper.GetString("db.port"),
		viper.GetString("db.name"))
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		panic(err)
	}
	_playerRepository.NewPlayerRepository(db)
}
