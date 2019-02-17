package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/spf13/viper"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"time"

	"./commandControl"
)

func main() {
	// READ CONFIG
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// INITIALIZE DATABASE
	db, err := bolt.Open("MCBot.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Verify all Top-Level-Buckets exist
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("Tracks"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
		return
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

	// REGISTER COMMANDS
	cc := commandControl.CommandControl{Bot: b, DB: db}

	b.Handle("/music", func(m *tb.Message) {
		cc.Music(m)
	})
	b.Handle("/list", func(m *tb.Message) {
		cc.List(m)
	})
	fmt.Println("Init done - waiting for messages")
	b.Start()
}
