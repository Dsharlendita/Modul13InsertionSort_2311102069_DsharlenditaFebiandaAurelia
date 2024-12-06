package main

import "fmt"

func insertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func main() {
	var data []int
	for {
		var x int
		fmt.Scan(&x)
		if x < 0 {
			break
		}
		data = append(data, x)
	}

	insertionSort(data)
	fmt.Println("Data terurut:", data)

	if len(data) > 1 {
		diff := data[1] - data[0]
		berjarakTetap := true
		for i := 1; i < len(data)-1; i++ {
			if data[i+1]-data[i] != diff {
				berjarakTetap = false
				break
			}
		}

		if berjarakTetap {
			fmt.Printf("Data berjarak %d\n", diff)
		} else {
			fmt.Println("Data berjarak tidak tetap")
		}
	}
}
