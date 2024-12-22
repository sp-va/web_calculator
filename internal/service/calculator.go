package service

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func evalSimpleExpression(a, b float64, action string) (float64, error) {
	var result float64
	switch action {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			return 0, errors.New("попытка деления на ноль")
		}
		result = a / b
	case "+-":
		result = a - b
	case "--":
		result = a + b
	case "*-":
		result = -(a * b)
	case "/-":
		if b == 0 {
			return 0, errors.New("попытка деления на ноль")
		}
		result = -(a / b)
	default:
		return 0, errors.New("некорректный символ")
	}
	return result, nil
}

func lookForHighpriority(actions []string) bool {
	highPriority := []string{"*", "/", "*-", "/-"}

	for _, b := range actions {
		for _, a := range highPriority {
			if b == a {
				return true
			}
		}
	}
	return false
}

func evalInsideParethensis(expr string) (float64, error) {
	numbers := []float64{}
	actions := []string{}

	var addNumber string
	numberIsNegative := false
	for i := 0; i < len(expr); i++ {
		v := string(expr[i])
		if v == "+" || v == "-" || v == "*" || v == "/" {
			if len(actions) > 0 && (actions[len(actions)-1] == v) {
				return 0, errors.New("два оператора подряд")
			}
			if len(addNumber) != 0 {
				result, err := strconv.ParseFloat(addNumber, 32)
				if err != nil {
					return 0, errors.New("не получилосьт перевести строку в float")
				}
				if numberIsNegative {
					numbers = append(numbers, -result)
					numberIsNegative = false
				} else {
					numbers = append(numbers, result)
				}
				addNumber = ""
			} else {
				numberIsNegative = true
			}
			if !numberIsNegative {
				if i+1 < len(expr) && string(expr[i+1]) == "-" {
					actions = append(actions, v+"-")
					i++
				} else {
					actions = append(actions, v)
				}
			}
		} else {
			addNumber += v
		}
	}
	result, err := strconv.ParseFloat(addNumber, 32)
	if err != nil {
		return 0, errors.New("не получилосьт перевести строку в float")
	}
	numbers = append(numbers, result)
	if len(numbers) != (len(actions) + 1) {
		return 0, errors.New("неверное количество операторов")
	}

	for {
		if lookForHighpriority(actions) {
			for i := 0; i < len(actions); i++ {
				if actions[i] == "*" || actions[i] == "/" || actions[i] == "*-" || actions[i] == "/-" {
					replaceValue, err := evalSimpleExpression(numbers[i], numbers[i+1], actions[i])
					if err != nil {
						return 0, err
					}
					numbers = append(numbers[:i], numbers[i+1:]...)
					numbers[i] = replaceValue
					actions = append(actions[:i], actions[i+1:]...)
					i--
				}
			}
		} else {
			break
		}
	}
	for i := 0; i < len(actions); i++ {
		replaceValue, err := evalSimpleExpression(numbers[i], numbers[i+1], actions[i])
		if err != nil {
			return 0, err
		}
		numbers = append(numbers[:i], numbers[i+1:]...)
		numbers[i] = replaceValue
		actions = append(actions[:i], actions[i+1:]...)
		i--
	}

	return numbers[0], nil
}

func Calc(expression string) (float64, error) {
	removedSpaces := strings.ReplaceAll(expression, " ", "")
	countOpenParethensis := strings.Count(removedSpaces, "(")
	countCloseParethensis := strings.Count(removedSpaces, ")")
	if countOpenParethensis != countCloseParethensis {
		return 0, errors.New("где-то не закрыта скобка")
	}

	for {
		if strings.ContainsAny(removedSpaces, "(") {
			var openPosition int
			var closePosition int
			for i, v := range removedSpaces {
				if v == '(' {
					openPosition = i
				}
				if v == ')' {
					closePosition = i
					break
				}
			}
			subStr := removedSpaces[openPosition+1 : closePosition]
			subStrResult, err := evalInsideParethensis(subStr)
			if err != nil {
				return 0, err
			}
			removedSpaces = removedSpaces[:openPosition] + fmt.Sprintf("%f", subStrResult) + removedSpaces[closePosition+1:]
			continue
		}
		finalResult, err := evalInsideParethensis(removedSpaces)
		if err != nil {
			return 0, err
		}
		return finalResult, nil
	}
}
