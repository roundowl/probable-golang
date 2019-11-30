package main

import "fmt"

type myInt int

// IntPlusTwo Increase int by 2
func (i *myInt) IntPlusTwo() {
	*i += myInt(2)
}

// ZeroDivision Dividing by zero here
func ZeroDivision() {
	var int1 uint = 1
	var int2 uint = 1
	var int3 int

	int2 = 0
	int3 = int(int1 / int2)

	fmt.Println(int3)

}
