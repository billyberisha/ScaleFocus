package main

import "fmt"

func daysInMonth(month int, year int) (int, bool) {
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		return 31, true
	case 4, 6, 9, 11:
		return 30, true
	case 2:
		return checkFeb(year), true
	default:
		return -1, false
	}
}

func checkFeb(year int) int {
	if (year%4 == 0) && (year%100 != 0 || year%400 == 0) {
		return 29
	}
	return 28
}

func main() {
	if days, ok := daysInMonth(2, 2000); ok {
		fmt.Println(days)
	} else {
		fmt.Println("Invalid month")
	}

}
