package musicData

import (
	"encoding/json"
	"fmt"
)

type Track struct {
	Title       string
	URL         string
	Hashtags    []string
	Description string
	MessageId   int
	Votes       int
}

func (t Track) Serialize() (encoded []byte, err error) {
	return json.Marshal(t)
}

func (t Track) ShortDescription() string {
	return fmt.Sprintf("%s (%d 👍 /like_%d):\n%s\n", t.Title, t.Votes, t.MessageId, t.URL)
}

func (t Track) IsValid() bool {
	return "" != t.URL
}
