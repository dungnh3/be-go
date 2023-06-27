package main

import (
	"fmt"
	"time"
)

// Section 0: This is comment in golang
/*
This is block comment in golang
*/

func main() {
	// Section 1: Variable
	var str string
	val := "hello"
	fmt.Println(str)
	fmt.Println(val)

	// Section 2: Constant
	const Pi = 3.14
	fmt.Println(Pi)

	// Section 3: Data type
	var a bool = true    // Boolean
	var b int = 5        // Integer
	var c float32 = 3.14 // Floating point number
	var d string = "Hi!" // String
	fmt.Println("Boolean: ", a)
	fmt.Println("Integer: ", b)
	fmt.Println("Float:   ", c)
	fmt.Println("String:  ", d)

	// Section 4: Array
	var arr1 = [3]int{1, 2, 3}
	arr2 := [5]int{4, 5, 6, 7, 8}
	fmt.Println(arr1)
	fmt.Println(arr2)

	var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}
	fmt.Print(cars)

	prices := [3]int{10, 20, 30}
	fmt.Println(prices[0])
	fmt.Println(prices[2])

	// Section 5: condition if/else

	// Section 6: switch case
	day := 4
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	}

	// Section 7: Loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// Section 8: Struct
	type Person struct {
		name   string
		age    int
		job    string
		salary int
	}

	// Section 9: Map
	var m1 = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	m2 := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}
	fmt.Println(m1)
	fmt.Println(m2)

	// Section 10: Go routine
	for i := 0; i < 5; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("done go routine")

	// Section 11: Channel
	ch := make(chan int, 5)
	for i := 0; i < 5; i++ {
		go func(i int) {
			ch <- i
		}(i)
	}

	go func() {
		for i := range ch {
			fmt.Println(i)
		}
	}()
	time.Sleep(1 * time.Second)
	close(ch)

	// Section 12: defer
	demoDefer()
	time.Sleep(1 * time.Second)
}

func demoDefer() {
	defer func() {
		fmt.Println("defer function")
	}()
	defer func() {
		fmt.Println("pre defer function")
	}()
	fmt.Println("before defer function")
}
