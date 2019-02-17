package commandControl

import (
	"../musicData"
	"encoding/json"
	"github.com/boltdb/bolt"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
)

type CommandControl struct {
	Bot *tb.Bot
	DB  *bolt.DB
}

// COMMANDS
func (cc CommandControl) Music(m *tb.Message) {
	// parse message
	p := AddCommandParser{m.Text}
	track := musicData.Track{
		URL:         p.GetURL(),
		Description: p.GetComment(),
		Hashtags:    p.GetHashtags(),
	}
	// is valid?
	if !track.IsValid() {
		cc.Bot.Send(m.Chat, "Brudi die URL fehlt 🙄")
		return
	}

	encoded, err := track.Serialize()
	if err != nil {
		log.Fatal(err)
		return
	}

	// save to database
	err = cc.DB.Update(func(tx *bolt.Tx) error {
		key := []byte(track.URL)
		tx.Bucket([]byte("Tracks")).Put(key, encoded)
		return nil
	})
	if err != nil {
		log.Fatal(err)
		return
	}
}

func (cc CommandControl) List(m *tb.Message) {
	err := cc.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Tracks"))
		c := b.Cursor()

		response := "Everything that matches your criteria:\n"
		for k, v := c.First(); k != nil; k, v = c.Next() {
			var t musicData.Track
			err := json.Unmarshal(v, &t)
			if err != nil {
				log.Print(err)
			}
			response += t.AsOneLine() + "\n"
		}

		cc.Bot.Send(m.Chat, response)
		return nil
	})
	if err != nil {
		log.Fatal(err)
		return
	}
}
