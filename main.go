package main;

import (
	"fmt"
)

type TestArr [] string

var testArr TestArr = [] string {"1", "2", "3"}

func (test TestArr)remove(str string) *TestArr{
	for i, item := range testArr {
		if item == str {
			fmt.Println("i:", i)
			testArr = append(testArr[:i], testArr[i+1:]...)
		}
	}
	return &testArr
}


func main () {
	fmt.Println("a", testArr)
	testArr.remove("2")
	fmt.Println("b", testArr)
}