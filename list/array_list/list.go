package array_list

import (
	"fmt"
	"strings"

	"github.com/kuber623/data_structure/common"
)

const (
	growthFactor = float32(2.0)  // growth by 100%
	shrinkFactor = float32(0.25) // shrink when size is 25% of capacity
)

type List struct {
	elements []interface{}
	size     int
}

func New(values ...interface{}) *List {
	return &List{
		elements: values,
		size:     len(values),
	}
}

func (list *List) Get(index int) (interface{}, bool) {
	if !list.withinRange(index) {
		return nil, false
	}
	return list.elements[index], true
}

func (list *List) Add(values ...interface{}) {
	list.growBy(len(values))
	for _, value := range values {
		list.elements[list.size] = value
		list.size++
	}
}

func (list *List) Remove(index int) {
	if !list.withinRange(index) {
		return
	}
	list.elements[index] = nil
	copy(list.elements[index:], list.elements[index+1:list.size])
	list.size--
	list.shrink()
}

func (list *List) Insert(index int, values ...interface{}) {
	if !list.withinRange(index) {
		if index == list.size {
			list.Add(values)
		}
		return
	}

	l := len(values)
	list.growBy(l)
	copy(list.elements[index+l:list.size+l], list.elements[index:list.size]) // left shift element
	copy(list.elements[index:], values)
	list.size += l
}

func (list *List) Set(index int, value interface{}) {
	if !list.withinRange(index) {
		return
	}
	list.elements[index] = value
}

func (list *List) Sort(comparator common.Comparator) {
	common.Sort(list.elements[:list.size], comparator)
}

func (list *List) Empty() bool {
	return list.size == 0
}

func (list *List) Size() int {
	return list.size
}

func (list *List) Clear() {
	list.size = 0
}

func (list *List) Values() []interface{} {
	return list.elements[:list.size]
}

func (list *List) String() string {
	str := "ArrayList\n"
	values := make([]string, 0)
	for _, value := range list.elements[:list.size] {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str += strings.Join(values, ", ")
	return str
}

func (list *List) withinRange(index int) bool {
	return index >= 0 && index < list.size
}

func (list *List) resize(cap int) {
	newElems := make([]interface{}, cap, cap)
	copy(newElems, list.elements)
	list.elements = newElems
}

func (list *List) growBy(n int) {
	capacity := len(list.elements)
	if capacity < list.size+n {
		list.resize(int(float32(capacity+n) * growthFactor))
	}
}

func (list *List) shrink() {
	if shrinkFactor == 0.0 {
		return
	}
	if float32(len(list.elements))*shrinkFactor > float32(list.size) {
		return
	}
	list.resize(list.size)
}
