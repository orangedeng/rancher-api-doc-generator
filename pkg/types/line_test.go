package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLine(t *testing.T) {
	type testcase struct {
		name   string
		source Line
		expect string
	}
	for _, c := range []testcase{
		{
			name:   "h1",
			source: H1Heading("test heading 1"),
			expect: "# test heading 1\n",
		},
		{
			name:   "h2",
			source: H2Heading("test heading 2"),
			expect: "## test heading 2\n",
		},
		{
			name:   "h3",
			source: H3Heading("test heading 3"),
			expect: "### test heading 3\n",
		},
		{
			name:   "h4",
			source: H4Heading("test heading 4"),
			expect: "#### test heading 4\n",
		},
		{
			name:   "h5",
			source: H5Heading("test heading 5"),
			expect: "##### test heading 5\n",
		},
		{
			name:   "h6",
			source: H6Heading("test heading 6"),
			expect: "###### test heading 6\n",
		},
	} {
		assert.Equalf(t, c.expect, c.source.String(), "test case %s failed", c.name)
	}
}
