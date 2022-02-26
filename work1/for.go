package main

import "fmt"

func main() {
	arr1 := [...]string{"i", "am", "stupid", "and", "weak"}

	for i := range arr1 {
		if i == 2 {
			arr1[i] = "smart"
		} else if i == 4 {
			arr1[i] = "strong"
		}
	}
	fmt.Println(arr1)

	for i := range arr1 {
		switch i {
		case 2:
			arr1[i] = "smart"
		case 4:
			arr1[i] = "strong"
		}
	}
	fmt.Println(arr1)
}
