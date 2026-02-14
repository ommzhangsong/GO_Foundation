package main

import "fmt"

func slice_append[T any](slice []T, index int, value T) []T {
	return append(slice[:index], append([]T{value}, slice[index:]...)...)
}
func slice_insert_usecopy[T any](slice []T, index int, value T) []T {
	var zero T
	slice = append(slice, zero)
	copy(slice[index+1:], slice[index:])
	slice[index] = value
	return slice
}
func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	brr := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(slice_append(arr, 3, 1))
	fmt.Println(slice_insert_usecopy(brr, 3, 1))

}
