package color

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestColor(t *testing.T) {
	Convey("test color init", t, func() {
		c := NewColor(100, 100, 100, 100)
		colorAsBinary := c.BinaryRGBA()
		So(len(colorAsBinary), ShouldEqual, 4)
		So(colorAsBinary[0], ShouldEqual, 100)
	})
}
