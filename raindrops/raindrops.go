package raindrops

import "fmt"

func Convert(number int) string {
	var isPling bool = number%3 == 0
	var isPlang bool = number%5 == 0
	var isPlong bool = number%7 == 0
	var isNone bool = !isPling && !isPlang && !isPlong

	var result = ""
	if isPling {
		result = result + "Pling"
	}
	if isPlang {
		result = result + "Plang"
	}
	if isPlong {
		result = result + "Plong"
	}
	if isNone {
		result = fmt.Sprint(number)
	}
	return result
}
