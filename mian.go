package main

import (
	"fmt"
)

func main() {
	input := 0
	inputs := make([]int, 0, 0)
	for {
		fmt.Scanln(&input)
		if input == -1 {
			break
		}
		inputs = append(inputs, input)
	}
	result := maopao(inputs)
	fmt.Println("maopao: ", result)
	result = xuanze(inputs)
	fmt.Println("xuanze: ", result)
	result = charu(inputs)
	fmt.Println("charu: ", result)
	result = xier(inputs)
	fmt.Println("xier: ", result)

	// result = guibing(inputs)
	// fmt.Println("xier: ", result)

	result = kuaipai(inputs)
	fmt.Println("kuaipai: ", result)
}

//冒泡排序
//遍历N轮
//每轮两两比较，把最大的沉底
func maopao(a []int) []int {
	b := make([]int, 0, 0)
	b = append(b, a...)
	lenb := len(b)
	fmt.Println("maopao first:", b)
	for i := 0; i < lenb; i++ { //表示循环多少次
		for j := 0; j < lenb-i-1; j++ { //表示遍历到第几个，已确定大小的在最后
			if b[j] > b[j+1] { //两两相邻比较
				b[j], b[j+1] = b[j+1], b[j]
			}
		}
		fmt.Println(b)
	}
	return b
}

//选择排序
//遍历一轮后找最小的防在已排序的末尾
func xuanze(a []int) []int {
	b := make([]int, 0, 0)
	b = append(b, a...)
	lenb := len(b)
	fmt.Println("xuanze first:", b)
	for i := 0; i < lenb-1; i++ { //表示找第几个数
		min := i
		for j := i + 1; j < lenb; j++ {
			if b[min] > b[j] { //记录最小的数的索引
				min = j
			}
		}
		b[min], b[i] = b[i], b[min] //交换
		fmt.Println(b)
	}
	return b
}

//插入排序
//每轮插入一个
//从已排好序的数组从前往后找，插在比其大的数的前面
func charu(a []int) []int {
	b := make([]int, 0, 0)
	b = append(b, a...)
	lenb := len(b)
	fmt.Println("charu first:", b)
	for i := 1; i < lenb; i++ { //从第二个数开始
		j := i - 1
		current := b[i]
		for j >= 0 && current < b[j] { //把比当前的数大的值后移，从后往前
			b[j+1] = b[j]
			j--
		}
		b[j+1] = current
		fmt.Println(b)
	}
	return b
}

func xier(arr []int) []int {
	b := make([]int, 0, 0)
	b = append(b, arr...)
	fmt.Println("xier first:", b)
	length := len(b)
	gap := 1
	for gap < length/3 {
		gap = gap*3 + 1
	}
	for gap > 0 {
		for i := gap; i < length; i++ {
			temp := b[i]
			j := i - gap
			for j >= 0 && b[j] > temp {
				b[j+gap] = b[j]
				j -= gap
			}
			b[j+gap] = temp
		}
		fmt.Println("xier second:", b)
		gap = gap / 3
	}
	return b
}

func guibing(arr []int) []int {
	len := len(arr)
	if len < 2 {
		return arr
	}
	middle := len / 2
	left := arr[0:middle]
	right := arr[middle:]
	return merge(guibing(left), guibing(right))
}

func merge(left []int, right []int) []int {
	var result []int
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	result = append(result, left...)
	result = append(result, right...)

	fmt.Println("guibing: ", result)
	return result
}

func kuaipai(arr []int) []int {
	b := make([]int, 0, 0)
	b = arr
	qsort(b)
	return b
}

func qsort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	pivot := parition(arr)
	qsort(arr[:pivot])
	qsort(arr[pivot+1:])
}

//这个是从两边找值再进行交换
func parition(arr []int) int {
	pivot := arr[0]
	left, right := 0, len(arr)-1
	for left < right {
		for left < right && pivot < arr[right] {
			right--
		}
		for left < right && pivot > arr[left] {
			left++
		}
		arr[left], arr[right] = arr[right], arr[left]
	}
	arr[left] = pivot
	return left
}

// func parition(arr []int) int {
// 	pivot := arr[0]
// 	left, right := 0, len(arr)-1
// 	for left < right {
// 		for left < right && pivot < arr[right] {
// 			right--
// 		}
// 		arr[left] = arr[right] //先把小的值房第一个，覆盖pivot 这时候right的位置空了出来
// 		for left < right && pivot > arr[left] {
// 			left++
// 		}
// 		arr[right] = arr[left]
// 	}
// 	arr[left] = pivot
// 	return left
// }

func Quick2Sort(values []int) {
	if len(values) <= 1 {
		return
	}
	mid, i := values[0], 1

	head, tail := 0, len(values)-1
	for head < tail {
		fmt.Println(values)
		if values[i] > mid {
			values[i], values[tail] = values[tail], values[i]
			tail--
		} else {
			values[i], values[head] = values[head], values[i]
			head++
			i++
		}
	}
	values[head] = mid
	Quick2Sort(values[:head])
	Quick2Sort(values[head+1:])
}
