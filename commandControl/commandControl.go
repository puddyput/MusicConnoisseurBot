package commandControl

import (
	"../musicData"
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type CommandControl struct {
	Bot *tb.Bot
	MDB *musicData.Database
}

// COMMANDS

// Music
func (cc CommandControl) Music(m *tb.Message) {

	// parse message
	p := AddCommandParser{m.Text}

	track := musicData.Track{
		URL:         p.GetURL(),
		Title:       p.getTitle(),
		Description: p.GetComment(),
		Hashtags:    p.GetHashtags(),
		MessageId:   m.ID,
	}
	// is valid?
	if !track.IsValid() {
		cc.Bot.Send(m.Chat, "Brudi die URL fehlt 🙄")
		return
	} else {
		// save to database
		cc.MDB.PutTrack(track.URL, track)

		// post vote link
		cc.Bot.Send(m.Chat, fmt.Sprintf("If you like it, send /like_%d", track.MessageId))
	}
}

// List
func (cc CommandControl) List(m *tb.Message) {
	re := regexp.MustCompile("/list\\s(.*)")
	matches := re.FindStringSubmatch(m.Text)
	search := ""
	if len(matches) > 1 {
		search = matches[1]
	}
	tracks := cc.MDB.Find(search)
	reply := "Tracks matching your criteria:\n"
	for _, t := range tracks {
		reply += t.ShortDescription() + "\n"
	}

	_, err := cc.Bot.Send(m.Chat, reply, tb.NoPreview)
	if nil != err {
		log.Print("could not send response to LIST command")
	}
}

// votes
func (cc CommandControl) HandleVote(m *tb.Message) {
	if strings.HasPrefix(m.Text, "/like_") {
		re := regexp.MustCompile("/like_(.*)@|$")
		matches := re.FindStringSubmatch(m.Text)
		id, _ := strconv.Atoi(matches[1])
		// check if id is a music post
		t := cc.MDB.FindByMessageId(id)
		if &t != nil {
			// update vote count
			t.Votes++
			cc.MDB.PutTrack(t.URL, t)
		}
	}
}
