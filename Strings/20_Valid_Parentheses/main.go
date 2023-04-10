package ValidParentheses

func isValid(s string) bool {
	var stack []rune

	for _, char := range s {
		switch char {
		case '(':
			stack = append(stack, ')')
		case '{':
			stack = append(stack, '}')
		case '[':
			stack = append(stack, ']')
		default:
			if len(stack) == 0 || char != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
