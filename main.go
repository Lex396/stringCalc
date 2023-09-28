package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Input")
	expression, _ := reader.ReadString('\n')

	expressionCheck := check(expression)
	result := calculate(expressionCheck)
	fmt.Print(result)
}

func calculate(expr []string) string {

	var res string

	var oper1 string = rmQout(expr[0])
	var oper2 string = rmQout(expr[2])
	var operator string = expr[1]

	var oper2Int, _ = strconv.Atoi(expr[2])

	switch operator {
	case "+":
		res = oper1 + oper2
	case "-":
		res = strings.ReplaceAll(oper1, oper2, "")
	case "*":
		res = strings.Repeat(oper1, oper2Int)
	case "/":
		for i, r := range oper1 {
			res = res + string(r)
			if i > 0 && (i+1)%oper2Int == 0 {
				break
			}
		}
	default:
		fmt.Println("незвестный оператор")
		os.Exit(1)

	}
	res = truncateText(res, 40)
	res = addQout(res)

	return res
}

func check(expression string) []string {
	expr := strings.TrimSpace(expression)
	// var regx string = "(\".*\")\\t\\n\\f\\r*([-+*\\/])\\t\\n\\f\\r*((\".*\")|[0-9])"
	// (".*")\s([-+*//])\s(".*"|\d)
	// "[\w\s]+"|([-+*//])|\d|("[\w\s]+")
	// "[^"]*"|[-+*/]|[0-9]
	// "[a-zA-z0-9_]+[^"]*"|[-+*/]|[0-9]
	var regx string = "\"[a-zA-z0-9_]+[^\"]*\"|[-+*/]|[0-9]|([a-zA-z]+)$"

	//parts := strings.SplitAfter(expr, "")
	parts := regexp.MustCompilePOSIX(regx).FindAllString(expr, -1)

	if len(parts) > 3 {
		fmt.Print("неверное выражение, много аргументов")
		os.Exit(1)
	} else if len(parts) < 2 {
		fmt.Println("неверное выражение, мало аргументов")
		os.Exit(1)
	}

	if parts[0][0] != '"' || parts[0][len(parts[0])-1] != '"' {
		fmt.Println("первым аргументом должна быть строка заключенная в символ \" \" ")
		os.Exit(1)
	}

	if len(parts[0]) > 10 {
		fmt.Println("первым аргументом должна быть строка не более 10 символов")
		os.Exit(1)
	}

	part2, _ := strconv.Atoi(parts[2])

	if part2 != 0 {
		if part2 < 0 || part2 > 10 {
			fmt.Print("неверный второй операнд, число должно быть от 1 до 10")
			os.Exit(1)
		} else {
			if parts[1] == "+" || parts[1] == "-" {
				if string(parts[2]) != "string" {
					fmt.Println("с данным оператором второй аргумент должен быть строкой и заключен в \" \" ")
					os.Exit(1)
				}
			}
		}
	} else {

		if len(parts[2]) > 10 {
			fmt.Println("вторым аргументом должна быть строка не более 10 символов")
			os.Exit(1)
		}

		if parts[2][0] != '"' || parts[2][len(parts[2])-1] != '"' {
			fmt.Println("второй аргумент должен быть заключенн в  \" \" ")
			os.Exit(1)
		}

		if parts[1] == "/" {
			fmt.Println("неподдерживаемая операция, нельзя делить строку на строку ")
			os.Exit(1)
		}

		if parts[1] == "*" {
			fmt.Println("неподдерживаемая операция, нельзя умножать строку на строку ")
			os.Exit(1)
		}
	}

	return parts
}

func rmQout(s string) string {
	return strings.ReplaceAll(s, "\"", "")
}

func addQout(s string) string {
	return "\"" + s + "\""
}

func truncateText(s string, max int) string {
	if len(s) > max {
		ii := 0
		for i := range s {
			ii++
			if ii > max {
				return s[:i] + "..."
			}
		}
	}
	return s
}
