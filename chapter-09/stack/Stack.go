package stack

type Stack struct {
	size int
	data []interface{}
}

func (this *Stack) Empty() bool {
	return this.size == 0
}

func (this *Stack) Peek() interface{} {
	if this.Empty() {
		panic("Empty stack.")
	}

	return this.data[this.size-1]
}

func (this *Stack) Pop() interface{} {
	if this.Empty() {
		panic("Empty stack.")
	}

	this.size--
	return this.data[this.size]
}

func (this *Stack) Push(data interface{}) {
	if this.size < len(this.data) {
		this.data[this.size] = data
	} else {
		this.data = append(this.data, data)
	}

	this.size++
}

func InitStack() *Stack {
	stack := &Stack{}
	stack.data = make([]interface{}, 0)
	return stack
}
