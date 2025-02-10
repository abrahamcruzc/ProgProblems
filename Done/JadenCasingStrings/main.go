package main

import "strings"

func ToJadenCase(str string) string{
	strArr := strings.Fields(str)

	for index, value := range strArr {
		strArr[index] = strings.ToUpper(value[:1]) + strings.ToLower(value[1:])
	}
	
	return strings.Join(strArr, " ")
}

