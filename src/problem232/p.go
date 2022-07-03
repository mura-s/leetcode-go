package problem232

type Stack []int

func (s *Stack) Push(x int) {
	*s = append(*s, x)
}
func (s *Stack) Pop() int {
	if s.Len() == 0 {
		return -1
	}
	x := (*s)[s.Len()-1]
	*s = (*s)[:s.Len()-1]
	return x
}

func (s *Stack) Peek() int {
	if s.Len() == 0 {
		return -1
	}
	return (*s)[s.Len()-1]
}

func (s *Stack) Len() int {
	return len(*s)
}

func (s *Stack) IsEmpty() bool {
	return s.Len() == 0
}

type MyQueue struct {
	FirstStack  *Stack
	SecondStack *Stack
}

func Constructor() MyQueue {
	return MyQueue{
		FirstStack:  &Stack{},
		SecondStack: &Stack{},
	}
}

func (this *MyQueue) Push(x int) {
	this.FirstStack.Push(x)
}

func (this *MyQueue) Pop() int {
	if !this.SecondStack.IsEmpty() {
		return this.SecondStack.Pop()
	}
	this.move()
	return this.SecondStack.Pop()
}

func (this *MyQueue) Peek() int {
	if !this.SecondStack.IsEmpty() {
		return this.SecondStack.Peek()
	}
	this.move()
	return this.SecondStack.Peek()
}

func (this *MyQueue) move() {
	for !this.FirstStack.IsEmpty() {
		x := this.FirstStack.Pop()
		this.SecondStack.Push(x)
	}
}

func (this *MyQueue) Empty() bool {
	return this.FirstStack.IsEmpty() && this.SecondStack.IsEmpty()
}
