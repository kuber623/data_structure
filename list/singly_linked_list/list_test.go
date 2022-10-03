package singly_linked_list

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/kuber623/data_structure/common"
)

func TestList(t *testing.T) {
	var list *List
	Convey("singly linked list", t, func() {
		Convey("new", func() {
			list = New(1, 2, 3, 4, 5)
			So(list.String(), ShouldResemble, `SinglyLinkedList
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
			So(list.String(), ShouldResemble, `SinglyLinkedList
1, 2, 3, 4, 5, 6, 7, 8, 9`)
		})
		Convey("add nil", func() {
			list.Add()
			So(list.String(), ShouldResemble, `SinglyLinkedList
1, 2, 3, 4, 5, 6, 7, 8, 9`)
		})
		Convey("remove", func() {
			list.Remove(7)
			So(list.String(), ShouldResemble, `SinglyLinkedList
1, 2, 3, 4, 5, 6, 7, 9`)
		})
		Convey("remove out of range", func() {
			list.Remove(100)
			So(list.String(), ShouldResemble, `SinglyLinkedList
1, 2, 3, 4, 5, 6, 7, 9`)

			list.Remove(-100)
			So(list.String(), ShouldResemble, `SinglyLinkedList
1, 2, 3, 4, 5, 6, 7, 9`)
		})
		Convey("insert", func() {
			list.Insert(5, 1, 2, 3, 4)
			So(list.String(), ShouldResemble, `SinglyLinkedList
1, 2, 3, 4, 5, 1, 2, 3, 4, 6, 7, 9`)
		})
		Convey("insert out of range", func() {
			list.Insert(100, 1, 2, 3, 4)
			So(list.String(), ShouldResemble, `SinglyLinkedList
1, 2, 3, 4, 5, 1, 2, 3, 4, 6, 7, 9`)
		})
		Convey("set", func() {
			list.Set(0, 2)
			So(list.String(), ShouldResemble, `SinglyLinkedList
2, 2, 3, 4, 5, 1, 2, 3, 4, 6, 7, 9`)
		})
		Convey("set out of range", func() {
			list.Set(100, 3)
			So(list.String(), ShouldResemble, `SinglyLinkedList
2, 2, 3, 4, 5, 1, 2, 3, 4, 6, 7, 9`)
		})
		Convey("sort", func() {
			list.Sort(common.NumberComparator)
			So(list.String(), ShouldResemble, `SinglyLinkedList
1, 2, 2, 2, 3, 3, 4, 4, 5, 6, 7, 9`)
		})
	})
}
