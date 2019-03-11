package main

import (
	"fmt"
	"github.com/spf13/viper"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"time"

	"./commandControl"
	"./musicData"
)

func main() {
	// READ CONFIG
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// CONNECT TO BOT API
	b, err := tb.NewBot(tb.Settings{
		Token:  viper.GetString("api.token"),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	mdb := musicData.Init("MCBot.db")

	// REGISTER COMMANDS
	cc := commandControl.CommandControl{Bot: b, MDB: mdb}
	b.Handle("/music", cc.Music)
	b.Handle("/list", cc.List)
	b.Handle(tb.OnText, cc.HandleVote)

	fmt.Println("Init done - waiting for messages")
	defer cc.MDB.Close()
	b.Start()
}
