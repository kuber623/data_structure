package singly_linked_list

type Iterator struct {
	list  *List
	index int
}

func (list *List) Iterator() Iterator {
	return Iterator{list: list, index: -1}
}

func (iter *Iterator) Index() int {
	return iter.index
}

func (iter *Iterator) Value() interface{} {
	value, _ := iter.list.Get(iter.index)
	return value
}

func (iter *Iterator) Next() bool {
	if iter.index < iter.list.size {
		iter.index++
	}
	return iter.list.withinRange(iter.index)
}

func (iter *Iterator) Begin() {
	iter.index = -1
}

func (iter *Iterator) Prev() bool {
	if iter.index >= 0 {
		iter.index--
	}
	return iter.list.withinRange(iter.index)
}

func (iter *Iterator) End() {
	iter.index = iter.list.size
}
