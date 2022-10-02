package container

type IndexIterator interface {
	Index() int
	Value() interface{}
	Next() bool
	Begin()
}

type IndexBiIterator interface {
	IndexIterator
	Prev() bool
	End()
}

type KeyIterator interface {
	Key() interface{}
	Value() interface{}
	Next() bool
	Begin()
}

type KeyBiIterator interface {
	KeyIterator
	Prev() bool
	End()
}
