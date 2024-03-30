package hyperscript_test

import (
	"strings"
	"testing"

	hs "github.com/cethien/gomponents-hyperscript"
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
	"github.com/stretchr/testify/assert"
)

func TestHyperscript(t *testing.T) {
	label := "HyperscriptBlock"

	input := hs.Hyperscript("on click alert('Hello world')")
	expected := g.Attr("_", "on click alert('Hello world')")

	t.Run(label, func(t *testing.T) {
		assert.Equal(t, expected, input)
	})
}

func TestHyperscriptBlock(t *testing.T) {
	label := "HyperscriptBlock"

	var inputSb strings.Builder
	hs.HyperscriptBlock(
		`def waitAndReturn()
wait 2s
return "I waited..."
end`,
	).Render(&inputSb)
	input := inputSb.String()

	var expectedSb strings.Builder
	Script(
		g.Attr("type", "text/hyperscript"),
		g.Text(
			`def waitAndReturn()
wait 2s
return "I waited..."
end`,
		),
	).Render(&expectedSb)
	expected := expectedSb.String()

	t.Run(label, func(t *testing.T) {
		assert.Equal(t, expected, input)
	})
}
