package commandControl

import (
	"mvdan.cc/xurls"
	"regexp"
)

type AddCommandParser struct {
	Message string
}

// regexp to find all hashtags
const reHashTags = `(?m:#\w+)`

func (p AddCommandParser) GetHashtags() []string {
	re := regexp.MustCompile(reHashTags)
	return re.FindAllString(p.Message, -1)
}

func (p AddCommandParser) GetURL() string {
	return xurls.Strict().FindString(p.Message)
}

func (p AddCommandParser) GetComment() string {
	return p.Message
}
