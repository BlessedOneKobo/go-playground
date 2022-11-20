package main

import "fmt"

func transform_array[T, K any](arr []T, fn func (T, int, []T) K) []K {
	t_arr := make([]K, len(arr))
	for i := 0; i < len(arr); i++ {
		t_arr[i] = fn(arr[i], i, arr)
	}
	return t_arr
}

func filter_array[T any](arr []T, fn func (T, int, []T) bool) []T {
	f_arr := make([]T, len(arr))
	j := 0
	for i := 0; i < len(arr); i++ {
		if cond := fn(arr[i], i, arr); cond {
			f_arr[j] = arr[i]
			j++
		}
	}
	return f_arr[:j]
}

func main() {
	arr := []int{3, 4, 9, -1, 3, 9}
	multiplied_by_two := transform_array(arr, func (val, index int, arr []int) int {
		return val * 2
	})
	only_negative := filter_array(multiplied_by_two, func (val, index int, arr []int) bool {
		return val < 0
	})
	fmt.Println("arr:")
	fmt.Println(arr)
	fmt.Println("multiplied_by_two:")
	fmt.Println(multiplied_by_two)
	fmt.Println("only_negative:")
	fmt.Println(only_negative)
}