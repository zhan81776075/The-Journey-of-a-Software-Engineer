package main

import "fmt"

func sum(a []int) int {
	s := 0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return s
}

func main() {
    //例子1: slice的基础用法
	var arr1 [6]int
	var slice1 []int = arr1[2:5] // item at index 5 not included!

	// load the array with integers: 0,1,2,3,4,5
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}

	// print the slice
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}

	fmt.Printf("The length of arr1 is %d\n", len(arr1))
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	// grow the slice
	slice1 = slice1[0:4]
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	ret := sum(slice1)
	fmt.Printf("The sum of slice is %d\n", ret)
	for i, value := range slice1 {
		fmt.Printf("The num %d of slice1 is %d\n", i, value)
	}

    //例子2:通过make初始化slice
	var slice2 []int = make([]int, 10)
	// load the array/slice:
	for i := 0; i < len(slice2); i++ {
		slice2[i] = 5 * i
	}

	// print the slice:
	for i := 0; i < len(slice2); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice2[i])
	}
	fmt.Printf("The length of slice1 is %d\n", len(slice2))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice2))

    //例子3:range循环
    items := [...]int{10, 20, 30, 40, 50}
    var sumItem int = 0
    for _, item := range items {
        sumItem += item * 2
    }
	fmt.Printf("The sum of items is %d\n", sumItem)
	// grow the slice beyond capacity
	//slice1 = slice1[0:7 ] // panic: runtime error: slice bound out of range

    //例子4:切片的复制与追加
    sl_from := []int{1, 2, 3}
    sl_to := make([]int, 10)

    n := copy(sl_to, sl_from)
    fmt.Println(sl_to)
    fmt.Printf("Copied %d elements\n", n) // n == 3

    sl3 := []int{1, 2, 3}
    sl3 = append(sl3, 4, 5, 6)
    fmt.Println(sl3)
	fmt.Printf("The length of sl3 is %d\n", len(sl3))
	fmt.Printf("The capacity of sl3 is %d\n", cap(sl3))

    //例子5:修改字符串
    s := "hello"
    c := []byte(s)
    c[0] = 'c'
    s2 := string(c) // s2 == "cello"
	fmt.Printf("s2 is %s\n", s2)
}
