package commandControl

import (
	"fmt"
	"github.com/boltdb/bolt"
	tb "gopkg.in/tucnak/telebot.v2"
)

type CommandControl struct {
	Bot *tb.Bot
	DB  *bolt.DB
}

// COMMANDS
func (c CommandControl) Music(m *tb.Message) {
	// parse message

}

func (c CommandControl) List(m *tb.Message) {
	fmt.Printf("/list")
}
