package main

func removeDuplicates(s string) string {
	stack := make([]rune, 0)
	for _, r := range s {
		if len(stack) == 0 || stack[len(stack)-1] != r {
			stack = append(stack, r)
			continue
		}
		if stack[len(stack)-1] == r {
			stack = append(stack[:len(stack)-1])
		}
	}
	return string(stack)
}
