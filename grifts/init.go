package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/tahkapaa/coke/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
