func trap(height []int) int {
	var ans int
	
	stack := NewStack()
	for i, num := range height {
		for !stack.IsEmpty() && num > height[stack.Peek()] {
			current := stack.Pop()
			if stack.IsEmpty() {
				break
			}
			// current 左边的值肯定比 current 大
			// 而 i 对应的值是右边的值，肯定比 current 大
			left, right := stack.Peek(), i
			height := min(height[right], height[left]) - height[current]
			ans += (right - left - 1) * height
		}
		stack.Push(i)
	}
	
	return ans
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

type Stack struct {
	values []int
}

func NewStack() *Stack {
	return &Stack{values: make([]int, 0)}
}

func (s *Stack) Push(v int) {
	s.values = append(s.values, v)
}

func (s *Stack) Pop() int {
	pop, temp := s.values[len(s.values) - 1], s.values[:len(s.values) - 1]
	s.values = temp
	return pop
}

func (s *Stack) Peek() int {
	return s.values[len(s.values) - 1]   
}

func (s *Stack) IsEmpty() bool {
	return len(s.values) == 0
}