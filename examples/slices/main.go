package main

import "fmt"

func main() {
	//Slice and Append
	arr := []int{1, 2, 3, 4}
	fmt.Println(arr, len(arr), cap(arr))
	arr = append(arr, 5)
	fmt.Println(arr, len(arr), cap(arr))
	arr2 := []int{1, 2, 3}
	fmt.Println(arr2, len(arr2), cap(arr2))
	arr2 = append(arr2, 4) // Capacity changes to 6 (2x)
	// arr2 = append(arr2, 4, 5)             // Capacity changes to 6 (2x)
	// arr2 = append(arr2, 4, 5, 6)          // Capacity changes to 6 (2x)
	// arr2 = append(arr2, 4, 5, 6, 7)       // Capacity changes to 8 (2x + 2)
	// arr2 = append(arr2, 4, 5, 6, 7, 8, 9) // Capacity changes to 8 (2x + 2 + 2)
	fmt.Println(arr2, len(arr2), cap(arr2))
	fmt.Println(arr2, len(arr2), cap(arr2))

	//Appending Slice to Slice
	arr2 = append(arr2, arr...)
	fmt.Println(arr2, len(arr2), cap(arr2))

	//Slicing a Slice
	arr3 := arr2[5:7]
	fmt.Println(arr, arr2, arr3)
	arr3[1] = 25 //Updates original array (arr2) as well but (arr) won't change
	fmt.Println(arr, arr2, arr3, len(arr3), cap(arr3))
	arr3 = append(arr3, 82, 83, 84)
	fmt.Println(arr, arr2, arr3, len(arr3), cap(arr3))

	//copy
	myarr1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	myarr2 := myarr1[5:7]
	var myarr2Copy = make([]int, len(myarr2))
	copy(myarr2Copy, myarr2)
	fmt.Println(myarr2Copy, myarr2, myarr1)
	myarr2Copy = append(myarr2Copy, 9, 10, 11)
	fmt.Println(myarr2Copy, myarr2, myarr1)
}
