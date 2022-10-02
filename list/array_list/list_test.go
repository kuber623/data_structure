package array_list

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/kuber623/data_structure/common"
)

func TestList(t *testing.T) {
	var list *List
	Convey("array list", t, func() {
		Convey("new", func() {
			list = New(1, 2, 3, 4, 5)
			So(list.String(), ShouldResemble, `ArrayList
1, 2, 3, 4, 5`)
		})
		Convey("get", func() {
			value, ok := list.Get(1)
			So(value, ShouldResemble, 2)
			So(ok, ShouldBeTrue)
		})
		Convey("get out of range", func() {
			value, ok := list.Get(100)
			So(value, ShouldBeNil)
			So(ok, ShouldBeFalse)

			value, ok = list.Get(-100)
			So(value, ShouldBeNil)
			So(ok, ShouldBeFalse)
		})
		Convey("add", func() {
			list.Add(6, 7, 8, 9)
			So(list.String(), ShouldResemble, `ArrayList
1, 2, 3, 4, 5, 6, 7, 8, 9`)
		})
		Convey("add nil", func() {
			list.Add()
			So(list.String(), ShouldResemble, `ArrayList
1, 2, 3, 4, 5, 6, 7, 8, 9`)
		})
		Convey("remove", func() {
			ok := list.Remove(7)
			So(list.String(), ShouldResemble, `ArrayList
1, 2, 3, 4, 5, 6, 7, 9`)
			So(ok, ShouldBeTrue)
		})
		Convey("remove out of range", func() {
			ok := list.Remove(100)
			So(list.String(), ShouldResemble, `ArrayList
1, 2, 3, 4, 5, 6, 7, 9`)
			So(ok, ShouldBeFalse)

			ok = list.Remove(-100)
			So(list.String(), ShouldResemble, `ArrayList
1, 2, 3, 4, 5, 6, 7, 9`)
			So(ok, ShouldBeFalse)
		})
		Convey("insert", func() {
			ok := list.Insert(5, 1, 2, 3, 4)
			So(list.String(), ShouldResemble, `ArrayList
1, 2, 3, 4, 5, 1, 2, 3, 4, 6, 7, 9`)
			So(ok, ShouldBeTrue)
		})
		Convey("insert out of range", func() {
			ok := list.Insert(100, 1, 2, 3, 4)
			So(list.String(), ShouldResemble, `ArrayList
1, 2, 3, 4, 5, 1, 2, 3, 4, 6, 7, 9`)
			So(ok, ShouldBeFalse)
		})
		Convey("set", func() {
			ok := list.Set(0, 2)
			So(list.String(), ShouldResemble, `ArrayList
2, 2, 3, 4, 5, 1, 2, 3, 4, 6, 7, 9`)
			So(ok, ShouldBeTrue)
		})
		Convey("set out of range", func() {
			ok := list.Set(100, 3)
			So(list.String(), ShouldResemble, `ArrayList
2, 2, 3, 4, 5, 1, 2, 3, 4, 6, 7, 9`)
			So(ok, ShouldBeFalse)
		})
		Convey("sort", func() {
			list.Sort(common.NumberComparator)
			So(list.String(), ShouldResemble, `ArrayList
1, 2, 2, 2, 3, 3, 4, 4, 5, 6, 7, 9`)
		})
	})
}
