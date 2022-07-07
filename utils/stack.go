package utils

import "sync"

type FloatStack struct {
	stack []float64
	lock  sync.Mutex
	sum   float64
	limit int
}

//newFloatStack returns a new FloatStack with the given limit.
func NewFloatStack(limit int) *FloatStack {
	return &FloatStack{
		limit: limit,
	}
}

// Push adds a new element to the stack.
func (s *FloatStack) Push(x float64) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if len(s.stack) == s.limit {
		s.sum -= s.stack[0]
		s.stack = s.stack[1:]
	}
	s.stack = append(s.stack, x)
	s.sum += x
}

// Pop removes the top element from the stack and returns it.
func (s *FloatStack) Pop() float64 {
	s.lock.Lock()
	defer s.lock.Unlock()
	if len(s.stack) == 0 {
		return 0
	}
	x := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	s.sum -= x
	return x
}

// Len returns the length of the stack.
func (s *FloatStack) Len() int {
	s.lock.Lock()
	defer s.lock.Unlock()
	return len(s.stack)
}

// Cap returns the capacity of the stack.
func (s *FloatStack) Cap() int {
	s.lock.Lock()
	defer s.lock.Unlock()
	return cap(s.stack)
}

//sum returns the sum of the stack.
func (s *FloatStack) Sum() float64 {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.sum
}

//Avg returns the average of the stack.
func (s *FloatStack) Avg() float64 {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.sum / float64(len(s.stack))
}
