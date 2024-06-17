package helper

import "fmt"

func StrConvert(arr []interface{}, separator string) string {
	str := ""

	for index := 0; index < len(arr); index++ {
		str1 := fmt.Sprintf("%v", arr[index])
		if index != len(arr)-1 {
			str += str1 + separator
		} else {
			str += str1
		}
	}
	return str
}
