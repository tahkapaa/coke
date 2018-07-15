package actions

import "github.com/gobuffalo/buffalo"

// TestTesthandler default implementation.
func TestTesthandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("test/testhandler.html"))
}
