func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	stack := make([]rune, 0, len(s))

	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
			continue
		}
		if len(stack) == 0 {
			return false
		}
		top := stack[len(stack)-1]
		if (top == '(' && c == ')') || (top == '[' && c == ']') || (top == '{' && c == '}') {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}
