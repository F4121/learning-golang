package main

import "fmt"

func main() {
	var nilai32 int32 = 32769
	var nilai64 int64 = int64(nilai32)

	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)

	// got value -32768 because cannot convert down from int32 to int16
	// int16 min -32768 to 32767
	// value is overflow from int32 set to int16
	// if you set value on nilai32 = 32769 the result on nilai16 should be -32767 or +1
	fmt.Println(nilai16)

	// Another convertion data from byte to string
	var name = "Jhon"
	var h = name[1]
	var eString = string(h)
	fmt.Println(eString)
}
