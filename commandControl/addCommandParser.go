package commandControl

import (
	"mvdan.cc/xurls"
	"regexp"
	"strings"
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
	// remove command
	re := regexp.MustCompile("/music\\b")
	m := re.ReplaceAllString(p.Message, "")

	return m
}

func (p AddCommandParser) getTitle() string {
	// remove command
	re := regexp.MustCompile("/music\\b")
	m := re.ReplaceAllString(p.Message, "")

	lines := strings.Split(m, "\n")
	return lines[0]
}
