package main

import "fmt"

type ErrorResponse struct {
	Error string
}
type SuccessResponse struct {
	Message string
}
type Response[T interface{}] struct {
	Code int
	Data T
}

func Find[T comparable](data []T, find T) int {
	for index, value := range data {
		if value == find {
			return index
		}
	}
	return -1
}
func main() {
	success := &Response[*SuccessResponse]{
		Code: 200,
		Data: &SuccessResponse{
			Message: "Success",
		},
	}
	err := &Response[*ErrorResponse]{
		Code: 200,
		Data: &ErrorResponse{
			Error: "Error",
		},
	}
	fmt.Println(success.Data.Message)
	fmt.Println(err.Data.Error)

	fmt.Println(Find([]int{1, 2, 3, 4, 5}, 4))
	fmt.Println(Find([]string{"Abc", "Def", "Ghi", "Jkl"}, "ADG"))
}
