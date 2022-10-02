package list

import (
	"github.com/kuber623/data_structure/common"
	"github.com/kuber623/data_structure/container"
)

type List interface {
	Get(index int) (interface{}, bool)
	Add(values ...interface{})
	Remove(index int)
	Insert(index int, values ...interface{})
	Set(index int, value interface{})
	Sort(comparator common.Comparator)

	container.Container
	// Empty() bool
	// Size() int
	// Clear()
	// Values() []interface{}
	// String() string
}
