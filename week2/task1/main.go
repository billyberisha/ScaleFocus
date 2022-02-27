package main

import "fmt"

func main() {
	daysInMonth(3, 2016)
}

func daysInMonth(month int, year int) (int, bool) {
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		fmt.Println("31")
		break
	case 4, 6, 9, 11:
		fmt.Println("30")
		break
	case 2:
		if (month%4 == 0) && (month%100 != 0 || month%400 == 0) {
			fmt.Println("29")
		} else {
			fmt.Println("28")
		}
		break

	}
	return 1, true
}
