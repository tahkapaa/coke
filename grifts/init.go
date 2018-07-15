package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/tahkapaa/test_stuff/buffalo/coke/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
