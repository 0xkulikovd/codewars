// https://www.codewars.com/kata/5277c8a221e209d3f6000b56

// i'm bad at programming, but...

package kata

func ValidBraces(str string) bool {
	braces := map[string]string{"(": ")", "{": "}", "[": "]"}
	stack := []string{} // while parsing "good" string - only opening braces should go in stack
	for _, val := range str {
		for opening, closing := range braces {
			if len(stack) != 0 && stack[len(stack)-1] == closing { // closing brace in stack => "bad" string
				return false
			}
			if len(stack) != 0 && string(val) == closing && stack[len(stack)-1] == opening { // cheking pairs of braces - opening in stack & closing in string
				stack = stack[:len(stack)-1]
				break
			} else if string(val) == opening || string(val) == closing { // appending if check failed
				stack = append(stack, string(val))
				break
			}
		}
	}
	return len(stack) == 0 // if string consist only of opening braces - leftover in stack should signal about "bad" string
}
