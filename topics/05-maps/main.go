package main

import (
	"fmt"
)

func main() {

	fmt.Println("Maps 1 -----------------------")

	days := make(map[int]string)
	days[0] = "Monday"
	days[1] = "Tuesday"
	days[2] = "Wednesday"
	days[3] = "Thursday"
	days[4] = "Friday"
	days[5] = "Saturday"
	days[6] = "SUNDAY"

	fmt.Println(days, len(days))

	days[6] = "Sunday"

	fmt.Println(days, len(days))

	delete(days, 6)

	fmt.Println(days, len(days))

	fmt.Println("Maps 2 -----------------------")

	students := make(map[string][]int)
	students["Kurt"] = []int{1, 3, 5, 7, 9}
	students["Miles"] = []int{2, 4, 6, 8}
	fmt.Println(students, len(students))
	fmt.Println(students["Miles"][3])

	fmt.Println("Maps 3 -----------------------")

	users := make(map[string]string)
	users["Kurt"] = "Grunge"
	users["Miles"] = "Jazz"

	user, exist := users["Kurt"]
	fmt.Println(user, exist)

	user, exist = users["NoExist"]
	fmt.Println(user, exist)
}
