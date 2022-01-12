package types

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

func NewMarkDown() *MD {
	return &MD{
		buffer: bytes.NewBuffer([]byte{}),
	}
}

type MD struct {
	buffer *bytes.Buffer
}

func (m *MD) Append(inputs ...interface{}) MarkdownContent {
	for _, input := range inputs {
		if input == nil {
			continue
		}
		switch value := input.(type) {
		case MultipleLine:
			m.buffer.WriteString(value.String() + "\n")
		case Line:
			m.buffer.WriteString(value.String() + "\n")
		case Words:
			m.buffer.WriteString(value.String())
		default:
			m.buffer.WriteString(fmt.Sprintf("%s\n", value))
		}
	}
	return m
}

func (m *MD) String() string {
	return m.buffer.String()
}

func (m *MD) Write(path string) error {
	return ioutil.WriteFile(path, m.buffer.Bytes(), 0666)
}
