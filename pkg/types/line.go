package types

import (
	"bytes"
	"strings"
)

type HeaderLevel int
type SimpleLine string

func (s SimpleLine) String() string {
	return string(s)
}

func (s SimpleLine) GetLine() string {
	return s.String()
}

const (
	h1 HeaderLevel = iota + 1
	h2
	h3
	h4
	h5
	h6
)

var (
	_ Line = Heading{}
	_ Line = HorizontalBreaker()
)

type Heading struct {
	level   HeaderLevel
	content string
}

func (h Heading) String() string {
	buffer := bytes.NewBuffer([]byte{})
	for i := 0; i < int(h.level); i++ {
		buffer.WriteRune('#')
	}
	buffer.WriteString(" " + h.content)
	return buffer.String() + "\n"
}

func (h Heading) GetLine() string {
	return h.String()
}

func H1Heading(input ...string) Line {
	return Heading{
		level:   h1,
		content: strings.Join(input, " "),
	}
}

func H2Heading(input ...string) Line {
	return Heading{
		level:   h2,
		content: strings.Join(input, " "),
	}
}

func H3Heading(input ...string) Line {
	return Heading{
		level:   h3,
		content: strings.Join(input, " "),
	}
}

func H4Heading(input ...string) Line {
	return Heading{
		level:   h4,
		content: strings.Join(input, " "),
	}
}

func H5Heading(input ...string) Line {
	return Heading{
		level:   h5,
		content: strings.Join(input, " "),
	}
}

func H6Heading(input ...string) Line {
	return Heading{
		level:   h6,
		content: strings.Join(input, " "),
	}
}

func HorizontalBreaker() SimpleLine {
	return SimpleLine("---")
}
