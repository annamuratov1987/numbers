package main

import (
	customList "numbers/pkg/list"
	"fmt"
	"strconv"
	"strings"
)

func ReadNumbers(s string) (customList.LinkedList, customList.LinkedList) {
	str := strings.Split(s, "[")

	firstNum := customList.NewLinkedList()
	secondNum := customList.NewLinkedList()


	for i:=0; i<len(str[1]); i++ {
		if str[1][i] == ']' {
			break
		}
		if str[1][i] != ',' {
			num, err := strconv.Atoi((string(str[1][i])))
			if err!=nil {continue}
			firstNum.Add(num)
		}
	}

	for i:=0; i<len(str[2]); i++ {
		if str[2][i] == ']' {
			break
		}
		if str[1][i] != ',' {
			num, err := strconv.Atoi((string(str[2][i])))
			if err!=nil {continue}
			secondNum.Add(num)
		}
	}

	return firstNum, secondNum
}

func AddNumbers(firstNum customList.LinkedList, secondNum customList.LinkedList) customList.LinkedList {
	result := customList.LinkedList{}
	sum := 0
	div := 0
	var numFirst, numSecond  interface{} = nil, nil
	i := 0
	for {
		numFirst = firstNum.Get(i)
		numSecond = secondNum.Get(i)
		if (numFirst == nil)&&(numSecond == nil)&&(div==0) {
			break
		}
		if numFirst == nil {
			numFirst = 0
		}
		if numSecond == nil {
			numSecond = 0
		}
		sum = numFirst.(int) + numSecond.(int) + div
		div = sum / 10
		result.Add(sum % 10)
		i++
	}
	return result
}

func main()  {

	var input string
	fmt.Scanf("%s\n", &input)

	firsNum, secondNum := ReadNumbers(input)
	sumList := AddNumbers(firsNum, secondNum)
	fmt.Print(sumList)

}
