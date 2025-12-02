package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var operation string
	var numbers string
	for {
		fmt.Print("Введите операцию (AVG-среднее, SUM-cумма, MED-медиана): ")
		fmt.Scan(&operation)
		if operation != "AVG" && operation != "SUM" && operation != "MED" {
			fmt.Println("Неправильно введена операция, попробйте снова")
			continue
		}
		fmt.Print("Введите целые числа через запятую: ")
		fmt.Scan(&numbers)
		slice := GoToSlice(numbers)
		switch operation {
		case "AVG":
			AVG(slice)
		case "SUM":
			SUM(slice)
		case "MED":
			MED(slice)
		}
		break
	}
}

func GoToSlice(numbers string) []int {
	slice_res := make([]int, 0)
	slice := strings.Split(numbers, ",")
	for i := range slice {
		res, _ := strconv.Atoi(slice[i])
		slice_res = append(slice_res, res)
	}
	return slice_res
}

func AVG(slice []int) {
	res := 0
	for _, v := range slice {
		res += v
	}
	result := float64(res / len(slice))
	fmt.Printf("AVG = %.1f", result)
}

func SUM(slice []int) {
	res := 0
	for _, v := range slice {
		res += v
	}
	fmt.Printf("SUM = %d", res)
}

func MED(slice []int) {
	sort.Ints(slice)
	if len(slice)%2 != 0 {
		res := len(slice)/2 + 1
		fmt.Println(slice[res])
	} else {
		res := len(slice) / 2
		result := float64((slice[res] + slice[res+1]) / 2)
		fmt.Printf("MED = %.1f", result)
	}
}
