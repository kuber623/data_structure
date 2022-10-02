package stack

type Stack struct {
	elements []interface{}
}

func (stack *Stack) Push(elem interface{}) {
	stack.elements = append(stack.elements, elem)
}

func (stack *Stack) Pop() {
	if stack.Empty() {
		panic("pop empty stack")
	}
	stack.elements = stack.elements[:len(stack.elements)-1]
}

func (stack *Stack) Top() interface{} {
	if stack.Empty() {
		return nil
	}
	return stack.elements[len(stack.elements)-1]
}

func (stack *Stack) Empty() bool {
	return stack.Size() == 0
}

func (stack *Stack) Size() int {
	return len(stack.elements)
}
