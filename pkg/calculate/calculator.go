package calculator

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type InvalidExpressionError struct {
    message string
}

func (e *InvalidExpressionError) Error() string {
    return e.message
}


func Calculate(expression string) (string, error) {
    expression = strings.ReplaceAll(expression, " ", "")
    if expression == "" {
        return "", &InvalidExpressionError{"Expression is empty"}
    }

	
    for _, char := range expression {
        if !((char >= '0' && char <= '9') || char == '+' || char == '-' || char == '*' || char == '/') {
            return "", &InvalidExpressionError{"Invalid Characters in Expression"}
        }
    }
    
    nums := strings.Split(expression, "+")
	sum := 0.0
    for _, numStr := range nums{
        num, err := strconv.ParseFloat(numStr, 64)
        if err != nil {
            return "", &InvalidExpressionError{"Invalid Number In Expression"}
        }
        sum += num
    }
    

	return fmt.Sprintf("%.2f",sum), nil // Return the result
}
