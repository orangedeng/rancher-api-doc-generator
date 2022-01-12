package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiline(t *testing.T) {
	type testcase struct {
		name   string
		source MultipleLine
		expect string
	}

	for _, c := range []testcase{
		{
			name:   "paragraph normal",
			source: Paragraph("test"),
			expect: "test\n",
		},
		{
			name:   "paragraph with words",
			source: Paragraph("prefix ", Code("code"), " suffix"),
			expect: "prefix `code` suffix\n",
		},
		{
			name:   "paragraph with new line and word",
			source: Paragraph("prefix ", Code("code"), " suffix.", NewLine(), "second line ", Blod("blod")),
			expect: "prefix `code` suffix.  \nsecond line **blod**\n",
		},
		{
			name:   "code paragraph with kind",
			source: CodeParagraph("json", `{"tmp":"abcd"}`),
			expect: "```json\n{\"tmp\":\"abcd\"}\n```\n",
		},
		{
			name:   "code paragraph without kind",
			source: CodeParagraph("", `{"tmp":"abcd"}`),
			expect: "```\n{\"tmp\":\"abcd\"}\n```\n",
		},
		{
			name:   "blockquote with paragraph",
			source: Blockquotes(Paragraph("prefix ", Code("code"), " suffix.", NewLine(), "second line ", Blod("blod"))),
			expect: "> prefix `code` suffix.  \n> second line **blod**\n",
		},
		{
			name: "blockquote with paragraph and blockquote",
			source: Blockquotes(
				Paragraph(
					"prefix ", Code("code"), " suffix.", NewLine(),
					"second line ", Blod("blod"), NewLine(),
					Blockquotes(Paragraph("inner")),
				),
			),
			expect: "> prefix `code` suffix.  \n> second line **blod**  \n>> inner\n",
		},
		{
			name: "test raw list",
			source: NewList(true).
				Append(NewWords("abcd"), nil).
				Append(Italic("code"), nil).
				Append(NewWords("it is ", Link{Anchor: "line", Link: "/testlink"}), nil),
			expect: `1. abcd
1. *code*
1. it is [line](/testlink)
`,
		},
		{
			name: "test sub paragraph in list item",
			source: NewList(false).
				Append(NewWords("first"), nil).
				Append(NewWords("second"), nil).
				Append(NewWords("third"), Paragraph("test paragraph ", Italic("italic"), NewLine(), "nextline")).
				Append(NewWords("fourth"), nil),
			expect: `- first
- second
- third

   test paragraph *italic*  
   nextline

- fourth
`,
		},
		{
			name: "test simple table",
			source: NewTable(TablesHeader{"header1", TableAlignNormal}, TablesHeader{"header2", TableAlignLeft}, TablesHeader{"header3", TableAlignCenter}, TablesHeader{"header4", TableAlignRight}).
				NewRow().Append(SimpleWord("colume 1")).Append(SimpleWord("colume 2")).Append(SimpleWord("colume 3")).Append(SimpleWord("colume 4")).
				NewRow().Append(Italic("italic")).Append(Blod("blod")).
				Table(),
			expect: `| header1 | header2 | header3 | header4 |
| --- | :-- | :-: | --: |
| colume 1 | colume 2 | colume 3 | colume 4 |
| *italic* | **blod** |||
`,
		},
	} {
		assert.Equalf(t, c.expect, c.source.String(), "test case %s failed", c.name)
	}
}

func TestTableGetSetColumn(t *testing.T) {
	testTable := NewTable(
		TablesHeader{Title: "column1"},
		TablesHeader{Title: "column2"},
	)
	testTable.NewRow().Append(SimpleWord("r1c1"), SimpleWord("r1c2"))
	testTable.NewRow().Append(SimpleWord("r2c1"), SimpleWord("r2c2"))
	if !assert.Equal(t, `| column1 | column2 |
| --- | --- |
| r1c1 | r1c2 |
| r2c1 | r2c2 |
`, testTable.String(), "failed to build a simple table") {
		t.Fatal()
	}
	testTable.GetRow(0).SetColumn(1, SimpleWord("modified"))
	if !assert.Equal(t, `| column1 | column2 |
| --- | --- |
| r1c1 | modified |
| r2c1 | r2c2 |
`, testTable.String(), "failed to modify row column") {
		t.Fatal()
	}
	newTable := NewTable(
		TablesHeader{Title: "column1-1"},
		TablesHeader{Title: "column2-1"},
	)
	for _, row := range testTable.CloneRows() {
		newTable.Append(row)
	}
	newTable.GetRow(1).SetColumn(1, SimpleWord("modified"))
	if !assert.Equal(t, `| column1-1 | column2-1 |
| --- | --- |
| r1c1 | modified |
| r2c1 | modified |
`, newTable.String(), "failed to modify row column") {
		t.Fatal()
	}
}
