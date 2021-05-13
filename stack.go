package sogstruct

type stack struct {
	values []string
}

// Construct a new stack
func newStack() *stack {
	return &stack{}
}

// Inserts an object at the top of the stack
func (s *stack) Push(value string) {
	s.values = append(s.values, value)
}

// Removes and returns the object at the top of the stack
func (s *stack) Pop() string {
	if len(s.values) == 0 {
		return ""
	}

	left := len(s.values) - 1
	result := s.values[left]

	if left == 0 {
		s.Clear()
	} else {
		s.values = s.values[:left]
	}

	return result
}

// Removes all objects from the stack
func (s *stack) Clear() {
	s.values = make([]string, 0)
}
