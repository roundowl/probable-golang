package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	playWithTypes()
	var myAwesomeInt myInt = myInt(1)
	myAwesomeInt.IntPlusTwo()
	println(myAwesomeInt)
}

func playWithTypes() {
	var int1 uint8
	int1 = 1
	var int2 uint8 = 1
	int2++
	int3 := 4
	int3 = 3
	var int4 = 4
	var int5, int6 uint8 = 5, 6

	// I know it's overkill
	var float1 float64 = 1.1

	var bool1 = false
	var bool2 bool = true

	string1 := "weeee"

	arr1 := [...]int{1, 2, 3}
	var arr2 [3]int
	arr2[2] = 6

	slice1 := []int8{1, 2, 3}
	slice2 := make([]int8, len(slice1))
	copy(slice2, slice1)
	slice2[1] = 5

	var mapSquares map[int]int = map[int]int{
		1: 1,
		2: 4,
		3: 9,
	}

	fmt.Println(int1, int2, int3, int4, int5, int6)
	fmt.Println(float1, bool1, bool2)
	fmt.Println(utf8.RuneCountInString(string1))
	fmt.Println(string(string1[0]))
	fmt.Println(arr1, arr2)
	fmt.Println(slice1, slice2)
	fmt.Println(mapSquares)

}
