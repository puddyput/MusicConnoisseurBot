package commandControl

import (
	"../musicData"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"regexp"
	"strings"
)

type CommandControl struct {
	Bot *tb.Bot
	MDB *musicData.Database
}

// COMMANDS
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
	}

	// save to database
	cc.MDB.PutTrack(track.URL, track)
}

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
		reply += t.AsOneLine() + "\n"
	}

	_, err := cc.Bot.Send(m.Chat, reply, tb.NoPreview)
	if nil != err {
		log.Print("could not send response to LIST command")
	}
}

func (cc CommandControl) HandleReply(m *tb.Message) {
	originalMessage := m.ReplyTo
	// check if original message is a music post
	t := cc.MDB.FindByMessageId(originalMessage.ID)
	if &t != nil {
		// is reply to a track
		if strings.Contains(m.Text, "👍") {
			// update vote count
			t.Votes++
			cc.MDB.PutTrack(t.URL, t)
		}
	}
}
