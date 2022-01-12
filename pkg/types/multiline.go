package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

var (
	_ MultipleLine = SimpleMultiline("")
	_ MultipleLine = Paragraph()
	_ MultipleLine = Blockquotes(SimpleMultiline(""))
	_ MultipleLine = CodeParagraph("test", "")
	_ MultipleLine = listItem{}
	_ MultipleLine = List{}
	_ MultipleLine = &Table{}
)

type TableAlign string

const (
	TableAlignNormal TableAlign = "---"
	TableAlignLeft   TableAlign = ":--"
	TableAlignCenter TableAlign = ":-:"
	TableAlignRight  TableAlign = "--:"
)

const (
	subContextIndent = "  "
)

type SimpleMultiline string

func (s SimpleMultiline) String() string {
	return string(s)
}
func (s SimpleMultiline) GetLines() []string {
	rtn := strings.Split(string(s), "\n")
	if rtn[len(rtn)-1] == "" {
		rtn = rtn[:len(rtn)-1]
	}
	return rtn
}

type LinePrefix struct {
	prefix string
	merge  bool
	lines  []string // assuming no \n in the end of each line
}

func (p LinePrefix) GetLines() []string {
	var lines []string
	for _, v := range p.lines {
		currentPrefix := p.prefix
		if p.merge && strings.HasPrefix(v, currentPrefix) {
			v = currentPrefix + v
		} else {
			v = currentPrefix + " " + v
		}
		lines = append(lines, v)
	}
	return lines
}

func (p LinePrefix) String() string {
	return strings.Join(p.GetLines(), "\n") + "\n"
}

func Paragraph(inputs ...interface{}) MultipleLine {
	holder := ""
	for i := 0; i < len(inputs); i++ {
		holder += "%s"
	}
	rtn := fmt.Sprintf(holder, inputs...)
	if !strings.HasSuffix(rtn, "\n") {
		rtn += "\n"
	}
	return SimpleMultiline(rtn)
}

func Blockquotes(input MultipleLine) MultipleLine {
	// assuming the last line is empty
	lines := strings.Split(input.String(), "\n")
	return LinePrefix{
		prefix: ">",
		merge:  true,
		lines:  lines[:len(lines)-1],
	}
}

func CodeParagraph(kind, content string) MultipleLine {
	if !strings.HasSuffix(content, "\n") {
		content += "\n"
	}
	return Paragraph("```", kind, "\n", content, "```")
}

func NewList(isOrdered bool) List {
	rtn := List{}
	if !isOrdered {
		rtn.prefix = "-"
	} else {
		rtn.prefix = "1."
	}
	return rtn
}

type List struct{ LinePrefix }

func (p List) Append(title Words, contents MultipleLine) List {
	p.lines = append(p.lines, listItem{
		title:    title,
		contents: contents,
	}.String())
	return p
}

type listItem struct {
	title    Words
	contents MultipleLine
}

func (p listItem) String() string {
	return strings.Join(p.GetLines(), "\n")
}

func (p listItem) GetLines() []string {
	var rtn []string
	rtn = append(rtn, p.title.String())
	if p.contents != nil {
		// empty line before sub content
		rtn = append(rtn, "")
		subContent := LinePrefix{
			prefix: subContextIndent,
			merge:  true,
			lines:  p.contents.GetLines(),
		}
		rtn = append(rtn, subContent.GetLines()...)
		// empty line after sub content
		rtn = append(rtn, "")
	}
	return rtn
}

type TablesHeader struct {
	Title string
	Align TableAlign
}

type Table struct {
	headers []TablesHeader
	rows    []Row
}

type Row interface {
	Append(...Words) Row
	NewRow() Row
	String() string
	Table() *Table
	SetColumn(int, Words)
	GetColumn(int) Words
	GetColumns() []Words
}

type row struct {
	table   *Table
	columns []Words
	index   int
}

func (r *row) Append(inputs ...Words) Row {
	count := len(r.columns)
	max := len(r.table.headers)
	if count+len(inputs) > max {
		logrus.Panic("add too much columns in table")
	}
	r.columns = append(r.columns, inputs...)
	return r
}

func (r *row) GetColumns() []Words {
	return r.columns
}

func (r *row) NewRow() Row {
	return r.table.NewRow()
}

func (r *row) String() string {
	buffer := bytes.NewBuffer([]byte{})
	buffer.WriteString("|")
	for i := 0; i < len(r.table.headers); i++ {
		if i < len(r.columns) && r.columns[i].String() != "" {
			buffer.WriteString(" " + r.columns[i].String() + " ")
		}
		buffer.WriteRune('|')
	}
	return buffer.String()
}

func (r *row) Table() *Table {
	return r.table
}

func (r *row) SetColumn(index int, input Words) {
	if index >= len(r.table.headers) {
		logrus.Panicf("set columen %d out of range %d", index, len(r.table.headers))
	}
	if index < len(r.columns) {
		r.columns[index] = input
	} else {
		for i := 0; i+len(r.columns) < index; i++ {
			r.columns = append(r.columns, SimpleWord(""))
		}
		r.columns = append(r.columns, input)
	}
}

func (r *row) GetColumn(index int) Words {
	return r.columns[index]
}

func NewTable(headers ...TablesHeader) *Table {
	for i, header := range headers {
		if string(header.Align) == "" {
			headers[i].Align = TableAlignNormal
		}
	}
	return &Table{
		headers: headers,
		rows:    []Row{},
	}
}

func (p *Table) Append(row Row) {
	newRow := p.NewRow()
	newRow.Append(row.GetColumns()...)
}

func (p *Table) NewRow() Row {
	r := &row{
		table:   p,
		index:   len(p.rows),
		columns: []Words{},
	}
	p.rows = append(p.rows, r)
	return r
}

func (p *Table) GetRow(index int) Row {
	return p.rows[index]
}

func (p *Table) CloneRows() []Row {
	rtn := make([]Row, len(p.rows))
	for i, v := range p.rows {
		rtn[i] = v
	}
	return rtn
}

func (p *Table) String() string {
	return strings.Join(p.GetLines(), "\n") + "\n"
}

func (p Table) GetLines() []string {
	var lines []string
	headerWriter := bytes.NewBuffer([]byte{})
	headerWriter.WriteString("|")
	seperator := bytes.NewBuffer([]byte{})
	seperator.WriteString("|")
	for _, header := range p.headers {
		headerWriter.WriteString(" " + header.Title + " |")
		seperator.WriteString(" " + string(header.Align) + " |")
	}

	lines = append(lines, headerWriter.String())
	lines = append(lines, seperator.String())
	for _, row := range p.rows {
		lines = append(lines, row.String())
	}
	return lines
}
