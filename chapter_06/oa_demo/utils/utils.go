package utils

import "fmt"

func Cal(n1 float64, n2 float64, operator byte) float64 {
	// 一个计算的函数
	var res float64
	switch operator {
		case '+':
			res = n1 + n2
		case '-':
			res = n1 - n2
		case '*':
			res = n1 * n2
		case '/':
			res = n1 / n2
		default:
			fmt.Println("操作符号错误...")
		
	}
	return res
}
