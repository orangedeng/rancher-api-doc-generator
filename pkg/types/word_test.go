package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWord(t *testing.T) {
	type testcase struct {
		name     string
		source   Words
		expected string
	}

	for _, c := range []testcase{
		{
			name: "title link",
			source: Link{
				Anchor: "test",
				Title:  "title",
				Link:   "/abcd/dddd",
			},
			expected: `[test](/abcd/dddd "title")`,
		},
		{
			name: "raw link",
			source: Link{
				Anchor: "test",
				Title:  "",
				Link:   "/abcd/dddd",
			},
			expected: "[test](/abcd/dddd)",
		},
		{
			name: "title image",
			source: Image{
				Anchor: "test",
				Title:  "title",
				Link:   "/abcd/dddd",
			},
			expected: `![test](/abcd/dddd "title")`,
		},
		{
			name: "raw image",
			source: Image{
				Anchor: "test",
				Title:  "",
				Link:   "/abcd/dddd",
			},
			expected: "![test](/abcd/dddd)",
		},
		{
			name:     "emphasis blod",
			source:   Blod("test", "emphasis", "blod"),
			expected: "**test emphasis blod**",
		},
		{
			name:     "emphasis italic",
			source:   Italic("test", "emphasis", "italic"),
			expected: "*test emphasis italic*",
		},
		{
			name:     "code block",
			source:   Code("test my code"),
			expected: "`test my code`",
		},
	} {
		assert.Equalf(t, c.expected, c.source.String(), "test case %s failed", c.name)
	}
}
