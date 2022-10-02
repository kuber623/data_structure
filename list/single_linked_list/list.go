package single_linked_list

import (
	"fmt"
	"strings"

	"github.com/kuber623/data_structure/common"
)

type List struct {
	head *element
	tail *element
	size int
}

type element struct {
	value interface{}
	next  *element
}

func New(values ...interface{}) *List {
	list := &List{}
	list.Add(values)
	return list
}

func (list *List) Get(index int) (interface{}, bool) {
	if !list.withinRange(index) {
		return nil, false
	}
	element := list.head
	for i := 0; i != index; i, element = i+1, element.next {
	}
	return element.value, true
}

func (list *List) Add(values ...interface{}) {
	for _, value := range values {
		element := &element{value: value}
		if list.size == 0 {
			list.head = element
			list.tail = element
		} else {
			list.tail.next = element
			list.tail = element
		}
		list.size++
	}
}

func (list *List) Append(values ...interface{}) {
	list.Add(values)
}

func (list *List) Remove(index int) {
	if !list.withinRange(index) {
		return
	}
	if list.size == 1 {
		list.Clear()
		return
	}

	before, element := new(element), list.head
	for i := 0; i != index; i, element = i+1, element.next {
		before = element
	}

	if element == list.head {
		list.head = element.next
	}
	if element == list.tail {
		list.tail = before
	}
	if before != nil {
		before.next = element.next
	}
	element = nil
	list.size--
}

func (list *List) Insert(index int, values ...interface{}) {
	if !list.withinRange(index) || len(values) == 0 {
		if index == list.size {
			list.Add(values)
		}
		return
	}

	before, element := new(element), list.head
	for i := 0; i != index; i, element = i+1, element.next {
		before = element
	}

	sublist := New(values)
	if element == list.head {
		sublist.tail.next = list.head
		list.head = sublist.head
	} else {
		before.next = sublist.head
		sublist.tail.next = element
	}
	list.size += sublist.size
	sublist.Clear()
}

func (list *List) Set(index int, value interface{}) {
	if !list.withinRange(index) {
		return
	}

	element := list.head
	for i := 0; i != index; i, element = i+1, element.next {
	}
	element.value = value
}

func (list *List) Sort(comparator common.Comparator) {
	values := list.Values()
	common.Sort(values, comparator)
	list.Clear()
	list.Add(values...)
}

func (list *List) Empty() bool {
	return list.size == 0
}

func (list *List) Size() int {
	return list.size
}

func (list *List) Clear() {
	list.head = nil
	list.tail = nil
	list.size = 0
}

func (list *List) Values() []interface{} {
	values := make([]interface{}, list.size, list.size)
	for index, element := 0, list.head; element != nil; index, element = index+1, element.next {
		values[index] = element.value
	}
	return values
}

func (list *List) String() string {
	str := "SinglyLinkedList\n"
	values := make([]string, 0, list.size)
	for element := list.head; element != nil; element = element.next {
		values = append(values, fmt.Sprintf("%v", element.value))
	}
	str += strings.Join(values, ", ")
	return str
}

func (list *List) withinRange(index int) bool {
	return index >= 0 && index < list.size
}
