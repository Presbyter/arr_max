package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	arr := []float64{1.23, 2.32, 4.53, 7.23, 1.22, 1.11}
	a, b, err := fooD(arr)
	if err != nil {
		fmt.Printf("get an error: %s", err)
		return
	}
	fmt.Println(arr)
	fmt.Printf("MaxA: %v, MaxB: %v", a, b)
}

func fooA(arr []float64) (float64, float64, error) {
	if len(arr) < 2 {
		return 0, 0, fmt.Errorf("arr length must more then two")
	}
	sort.SliceStable(arr, func(i int, j int) bool { return arr[i] >= arr[j] })
	return arr[0], arr[1], nil
}

func fooB(arr []float64) (float64, float64, error) {
	if len(arr) < 2 {
		return 0, 0, fmt.Errorf("arr length must more then two")
	}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] > arr[i] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr[0], arr[1], nil
}

func fooC(arr []float64) (float64, float64, error) {
	if len(arr) < 2 {
		return 0, 0, fmt.Errorf("arr length must more then two")
	}

	sortFunc(arr, 0, len(arr)-1)

	return arr[0], arr[1], nil
}

func sortFunc(arr []float64, left, right int) {
	if left > right {
		return
	}
	l, r := left, right
	pivot := arr[l]
	for l < r {
		for l < r && pivot >= arr[r] {
			r--
		}
		arr[l] = arr[r]
		for l < r && pivot <= arr[l] {
			l++
		}
		arr[r] = arr[l]
	}
	arr[l] = pivot
	sortFunc(arr, left, l-1)
	sortFunc(arr, l+1, right)
}

type ArrEx []float64

func (a ArrEx) Len() int {
	return len(a)
}

func (a ArrEx) Less(i, j int) bool {
	return a[i] > a[j]
}

func (a ArrEx) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a *ArrEx) Push(x interface{}) {
	*a = append(*a, x.(float64))
}

func (a *ArrEx) Pop() interface{} {
	n := len(*a)
	res := (*a)[n-1]
	*a = (*a)[0 : n-1]
	return res
}

func fooD(arr []float64) (float64, float64, error) {
	if len(arr) < 2 {
		return 0, 0, fmt.Errorf("arr length must more then two")
	}
	arrEx := &ArrEx{}

	for _, v := range arr {
		heap.Push(arrEx, v)
	}

	return heap.Pop(arrEx).(float64), heap.Pop(arrEx).(float64), nil
}
