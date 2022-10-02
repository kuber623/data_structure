package array_list

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIterator(t *testing.T) {
	list := New(3, 5, 21, 23, 43, 222)
	iter := list.Iterator()
	Convey("array list iterator", t, func() {
		Convey("traversal", func() {
			values := make([]interface{}, 0, list.Size())
			for iter.Begin(); iter.Next(); {
				values = append(values, iter.Value())
			}
			So(values, ShouldResemble, list.Values())
		})
		Convey("reverse traversal", func() {
			values := make([]interface{}, 0, list.Size())
			for iter.End(); iter.Prev(); {
				values = append(values, iter.Value())
			}
			So(values, ShouldResemble, []interface{}{222, 43, 23, 21, 5, 3})
		})
	})
}
