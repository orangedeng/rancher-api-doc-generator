package types

import (
	"bytes"
	"fmt"
	"strings"
)

var (
	_ Words = WordWrapper{}
	_ Words = Link{}
	_ Words = Image{}
)

type WordWrapper struct {
	wrapper string
	content string
}

func (w WordWrapper) String() string {
	return w.wrapper + w.content + w.wrapper
}

func NewWords(inputs ...interface{}) Words {
	holder := ""
	for i := 0; i < len(inputs); i++ {
		holder += "%s"
	}
	return SimpleWord(fmt.Sprintf(holder, inputs...))
}

type (
	SimpleWord string
)

func (s SimpleWord) String() string {
	return string(s)
}

type Link struct {
	Anchor string
	Title  string
	Link   string
}

func (l Link) String() string {
	buffer := bytes.NewBuffer([]byte{})
	buffer.WriteString("[" + l.Anchor + "](" + l.Link)
	if l.Title != "" {
		buffer.WriteString(` "` + l.Title + `"`)
	}
	buffer.WriteString(")")
	return buffer.String()
}

type Image struct {
	Anchor string
	Title  string
	Link   string
}

func (l Image) String() string {
	buffer := bytes.NewBuffer([]byte{})
	buffer.WriteString("![" + l.Anchor + "](" + l.Link)
	if l.Title != "" {
		buffer.WriteString(` "` + l.Title + `"`)
	}
	buffer.WriteString(")")
	return buffer.String()
}

var (
	EmphasisBlodWrapper   = "**"
	EmphasisItalicWrapper = "*"
	CodeBlockWrapper      = "`"
)

func Blod(input ...string) Words {
	return WordWrapper{
		wrapper: EmphasisBlodWrapper,
		content: strings.Join(input, " "),
	}
}

func Italic(input ...string) Words {
	return WordWrapper{
		wrapper: EmphasisItalicWrapper,
		content: strings.Join(input, " "),
	}
}

func Code(input ...string) Words {
	return WordWrapper{
		wrapper: CodeBlockWrapper,
		content: strings.Join(input, " "),
	}
}

func NewLine() Words {
	return SimpleWord("  \n")
}
