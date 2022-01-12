package types

type MarkdownContent interface {
	Append(...interface{}) MarkdownContent
	String() string
	Write(path string) error
}

type Line interface {
	String() string
	GetLine() string
}

type MultipleLine interface {
	String() string
	GetLines() []string
}

type Words interface {
	String() string
}
