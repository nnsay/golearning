package unittest

import (
	"testing"

	bdd "github.com/smartystreets/goconvey/convey"
)

func TestBDD(t *testing.T) {
	bdd.Convey("Give two number", t, func() {
		a := 2
		b := 4

		bdd.Convey("When add two number", func() {
			ret := a + b

			bdd.Convey("Then result is still even", func() {
				bdd.So(ret%2, bdd.ShouldEqual, 0)
			})

		})

	})

}
