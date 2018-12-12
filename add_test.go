package add_test

import (
	"fmt"
	"testing"

	"github.com/Mazafard/add"
	. "github.com/smartystreets/goconvey/convey"
)

func TestConvert(t *testing.T) {
	Convert := add.Convert

	Convey("Should convert correctly", t, func() {
		Convey("Small numbers should convert correctly", func() {
			So(Convert(0), ShouldEqual, "صفر")
			So(Convert(1), ShouldEqual, "یک")
			So(Convert(5), ShouldEqual, "پنج")
			So(Convert(10), ShouldEqual, "ده")
			So(Convert(11), ShouldEqual, "یازده")
			So(Convert(12), ShouldEqual, "دوازده")
			So(Convert(17), ShouldEqual, "هفده")
		})
		Convey("Tens should convert correctly", func() {
			So(Convert(20), ShouldEqual, "بیست")
			So(Convert(30), ShouldEqual, "سی")
			So(Convert(40), ShouldEqual, "چهل")
			So(Convert(50), ShouldEqual, "پنجاه")
			So(Convert(60), ShouldEqual, "شصت")
			So(Convert(90), ShouldEqual, "نود")
		})
		Convey("Combined numbers should convert correctly", func() {
			So(Convert(21), ShouldEqual, "بیست و یک")
			So(Convert(34), ShouldEqual, "سی و چهار")
			So(Convert(49), ShouldEqual, "چهل و نه")
			So(Convert(53), ShouldEqual, "پنجاه و سه")
			So(Convert(68), ShouldEqual, "شصت و هشت")
			So(Convert(99), ShouldEqual, "نود و نه")
		})
		Convey("Big numbers should convert correctly", func() {
			So(Convert(100), ShouldEqual, "صد")
			So(Convert(200), ShouldEqual, "دویست")
			So(Convert(500), ShouldEqual, "پانصد")
			So(Convert(123), ShouldEqual, "صد و بیست و سه")
			So(Convert(666), ShouldEqual, "ششصد و شصت و شش")
			So(Convert(1024), ShouldEqual, "یک هزار و بیست و چهار")
		})
		Convey("Negative numbers should convert correctly", func() {
			So(Convert(-123), ShouldEqual, "منفی صد و بیست و سه")
		})

	})
}

func ExampleConvert() {
	fmt.Println(add.Convert(11))
	fmt.Println(add.Convert(123))
	fmt.Println(add.Convert(-99))
	// Output: یازده
	// صد و بیست و سه
	// منفی نود و نه
}
