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
	Rating      int
}

func (t Track) Serialize() (encoded []byte, err error) {
	return json.Marshal(t)
}

func (t Track) AsOneLine() string {
	return fmt.Sprintf("%s (%d): \t%s", t.Title, t.Rating, t.URL)
}

func (t Track) IsValid() bool {
	return "" != t.URL
}
