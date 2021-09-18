package main

import "fmt"

type SliceHeader struct {
	Data uintptr
	Len int 
	Cap int 
}



func main(){
	// create emplty slice.
	aSlice := [] float64 {}
	// printing len and capacity.
	fmt.Println(aSlice, len(aSlice), cap(aSlice))
	// add elements to slice.
	aSlice = append(aSlice, 345.78)
	aSlice = append(aSlice, 45.67)
	fmt.Println(aSlice, "current length is  --->", len(aSlice))

	// create a slice specifying the length.
	sliceWithLen := make([] int, 5)
	sliceWithLen[0] = 1
	sliceWithLen[1] = 2
	sliceWithLen[2] = 3
	sliceWithLen[3] = 4
	sliceWithLen[4] = 5

	// append element to slice.
	sliceWithLen = append(sliceWithLen,6)

	fmt.Println(sliceWithLen)

	fmt.Println("---- TwoD array")
	// 2D slice. 
	twoD := [][] int {{1,2,3},{4,5,6}}
	//loop thru the twoD slice.
	for _,v := range twoD{
		for _,k := range v {
			fmt.Println(k, "")

		}
		fmt.Println()
	}

	make2D := make([][]int, 2)
    fmt.Println(make2D)
    make2D[0] = []int{1, 2, 3, 4}
    make2D[1] = []int{-1, -2, -3, -4}
    fmt.Println(make2D)




}