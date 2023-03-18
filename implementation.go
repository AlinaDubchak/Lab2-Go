package lab2

import "fmt"

func SolvePostfixExpression(expression string) (string, error) {
	stack := make([]float32, 0)
	for i := 0; i < len(expression); i++ {
		switch string(expression[i]) {
		case "+", "-", "*", "/", "^":
			if len(stack) < 2 {
				return "", errors.New("invalid expression")
			}
		}
		switch string(expression[i]) {
		case " ":
			continue
		case "+":
			num2 := stack[len(stack)-1]
			num1 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			result := num1 + num2
			stack = append(stack, result)
		case "*":
			num2 := stack[len(stack)-1]
			num1 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			result := num1 * num2
			stack = append(stack, result)
		case "-":
			num2 := stack[len(stack)-1]
			num1 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			result := num1 - num2
			stack = append(stack, result)
		case "/":
			num2 := stack[len(stack)-1]
			num1 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			result := num1 / num2
			stack = append(stack, result)
		case "^":
			num2 := stack[len(stack)-1]
			num1 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			result := math.Pow(float64(num1), float64(num2))
			stack = append(stack, float32(result))
		default:
			num, err := strconv.ParseFloat(string(expression[i]), 32)
			if err != nil {
				return "", errors.New("invalid expression")
			}
			stack = append(stack, float32(num))
		}
	}
	if len(stack) != 1 {
		return "", errors.New("invalid expression")
	}
	strRes := fmt.Sprintf("%.2f", stack[0])
	return strRes, nil
}
