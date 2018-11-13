package stacker

type stack struct {
	size uint
	arr  []string
}

func (s *stack) Push(op string) {
	s.arr = append(s.arr, op)
	s.size++
}

func (s *stack) Pop() string {
	topOp := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	s.size--
	return topOp
}

func (s *stack) Peek() string {
	return s.arr[len(s.arr)-1]
}

type floatStack struct {
	size uint
	arr  []float64
}

func (s *floatStack) Push(op float64) {
	s.arr = append(s.arr, op)
	s.size++
}

func (s *floatStack) Pop() float64 {
	topOp := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	s.size--
	return topOp
}

func (s *floatStack) Peek() float64 {
	return s.arr[len(s.arr)-1]
}
