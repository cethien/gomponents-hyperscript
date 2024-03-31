// Package hyperscript adds hyperscript components to gomponents.
package hyperscript

import (
	g "github.com/maragudk/gomponents"
	//lint:ignore ST1001 this is fine
	. "github.com/maragudk/gomponents/html"
)

func Hyperscript(script string) g.Node {
	return g.Attr("_", script)
}

func HyperscriptBlock(script string) g.Node {
	return Script(
		g.Attr("type", "text/hyperscript"),
		g.Text(script),
	)
}
