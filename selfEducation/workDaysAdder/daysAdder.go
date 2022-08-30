package main

import "fmt"

type Days int

const (
	monday Days = iota + 1
	tuesday
	wednesday
	thursday
	friday
	saturday
	sunday
)

func main() {
	d := friday.addWorkDays(3)
	fmt.Println(d)
}

func (d Days) addWorkDays(n int) string {
	count := n
	workDays := []string{
		"monday",
		"tuesday",
		"wednesday",
		"thursday",
		"friday",
	}
	result := d
	for n > 0 {
		result++
		n--
		if result > 5 {
			result = 1
		}
	}

	resultStr := fmt.Sprintf("%s + %d workdays is %s", workDays[d-1], count, workDays[result-1])
	return resultStr
}
