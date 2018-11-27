package splice_append

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


func init () {
	fmt.Println("a", testArr)
	testArr.remove("3")
	fmt.Println("b", testArr)
}