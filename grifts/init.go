package grifts

import (
	"github.com/SerhiiCho/reciper_go/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
