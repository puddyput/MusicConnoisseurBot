package commandControl

import (
	"regexp"
	"strings"
)

type AddCommandParser struct {
	Message string
}

// regexp to find all hashtags
const reHashTags = `(?m:#\w+)`

// regexp to find all URLs
const reURL = `(?m:^(http:\/\/www\.|https:\/\/www\.|http:\/\/|https:\/\/)?` +
	`[a-z0-9]+([\-\.]{1}[a-z0-9]+)*\.[a-z]{2,5}(:[0-9]{1,5})?(\/.*)?$)`

func (p AddCommandParser) GetHashtags() []string {
	re := regexp.MustCompile(reHashTags)
	return re.FindAllString(p.Message, -1)
}

func (p AddCommandParser) GetURL() string {
	// taken from https://www.regextester.com/93652
	re := regexp.MustCompile(reURL)
	result := re.FindString(p.Message)
	return strings.TrimRight(result, " 	")
}

func (p AddCommandParser) GetComment() string {
	return p.Message
}
