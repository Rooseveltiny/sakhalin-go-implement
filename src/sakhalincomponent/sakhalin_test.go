package sakhalincomponent

import (
	"context"
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSakhalinInit(t *testing.T) {
	Convey("init test sakhalin", t, func(c C) {
		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		newSakhalin := NewSakhalinConnection("/ws", ctx)
		fmt.Println(newSakhalin)
		cancel()
	})
}
