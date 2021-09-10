package equation

import "sync"

type stack struct {
	list  []uint
	mutex *sync.Mutex
}

func newStack() stack {
	return stack{
		list:  make([]uint, 0),
		mutex: &sync.Mutex{},
	}
}

func (s *stack) push(item uint) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.list = append(s.list, item)
}

func (s *stack) pop() (uint, bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if len(s.list) == 0 {
		return 0, false
	}

	value := s.list[len(s.list)-1]
	s.list = s.list[:len(s.list)-1]

	return value, true
}

func (s *stack) peek() (uint, bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if len(s.list) < 1 {
		return 0, false
	}

	return s.list[len(s.list)-1], true
}

func (s *stack) size() int {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return len(s.list)
}
