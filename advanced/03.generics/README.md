# Generics

Golang supports generic types and parameters

## Generic Types

```go
type ErrorResponse struct {
	Error string
}
type SuccessResponse struct {
	Message string
}
type Response[T any] struct {
	Code int
	Data T
}
```

In the above example, we created `Response` which is a generic type. Response accepts a generic variable `T` which can be anything so we give it a type as `any`. `any` is short form for `interface{}`.

Generics are basically types which we define at the time of initializing the value.

```go
func main() {
	success := &Response[*SuccessResponse]{ //T is replaced by *SuccessResponse
		Code: 200,
		Data: &SuccessResponse{
			Message: "Success",
		},
	}
	err := &Response[*ErrorResponse]{ //T is replaced by *ErrorResponse
		Code: 200,
		Data: &ErrorResponse{
			Error: "Error",
		},
	}
	fmt.Println(success.Data.Message)
	fmt.Println(err.Data.Error)
}

```

In the above example, we used the response type to create a success response and an error response. Both have a `data` key which can hold different types based on the type of generic type provided at the time of `initialization`.

## Type Parameters

```go
func Find[T comparable](data []T, find T) int {
	for index, value := range data {
		if value == find {
			return index
		}
	}
	return -1
}
```

In the above function, we have created a function which expects generic arguments. We have defined the generic in the square brackets `[]` to the left of the actual arguments.

With a generic variable, we can pass a constraint. A constraint defines what things can be done with the generic. In our case we have passed `comparable` which tells that two values of `T` can be compared with `==` or `!=` operators. You can also use `any` as a constraint.

When we use it, we will use it as follows

```go
func main(){
  fmt.Println(Find([]int{1, 2, 3, 4, 5}, 4))
	fmt.Println(Find([]string{"Abc", "Def", "Ghi", "Jkl"}, "ADG"))
}
```