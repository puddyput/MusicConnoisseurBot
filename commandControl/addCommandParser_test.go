package commandControl

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestAddCommandParser_GetHashtags(t *testing.T) {
	input := "no #yes #yes2#yes3\n#yes4\n#nospecial!characters"
	expected := []string{"#yes", "#yes2", "#yes3", "#yes4", "#nospecial"}
	p := AddCommandParser{input}
	assert.Equal(t, p.GetHashtags(), expected)

	input = "nothing"
	p = AddCommandParser{input}
	assert.Equal(t, len(p.GetHashtags()), 0)
}

func TestAddCommandParser_GetURL(t *testing.T) {
	input := "#music\ncomment\nhttps://www.youtube.com/watch?v=hPGeSVgvgzQ"
	expected := "https://www.youtube.com/watch?v=hPGeSVgvgzQ"
	p := AddCommandParser{input}
	assert.Equal(t, p.GetURL(), expected)

	input = "https://www.youtube.com/watch?v=hPGeSVgvgzQ "
	p = AddCommandParser{input}
	assert.Equal(t, p.GetURL(), expected)

	input = "nothing"
	p = AddCommandParser{input}
	assert.Equal(t, p.GetURL(), "")
}

func TestAddCommandParser_GetComment(t *testing.T) {
	input := "#music\ncomment is here #hashtag included\nhttps://www.youtube.com/watch?v=hPGeSVgvgzQ"
	expected := "#music\ncomment is here #hashtag included\nhttps://www.youtube.com/watch?v=hPGeSVgvgzQ"
	p := AddCommandParser{input}
	assert.Equal(t, p.GetComment(), expected)
}
