package container

import "github.com/kuber623/data_structure/common"

// Container is base interface that all data structures implement
type Container interface {
	Empty() bool
	Size() int
	Clear()
	Values() []interface{}
}

// GetSortedValues returns sorted container's elements with respect to the passed comparator
// Does not affect the ordering of elements within the container
func GetSortedValues(container Container, comparator common.Comparator) []interface{} {
	values := container.Values()
	if len(values) < 2 {
		return values
	}
	common.Sort(values, comparator)
	return values
}
