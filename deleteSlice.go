package main

import (
	"fmt"
	"os"
	"strconv"
)


func main(){
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("NEED AN INTEGER VALUE COME ON")
		return 
	}

	index := arguments[1]

	i , err := strconv.Atoi(index)

	if err != nil{
		fmt.Println(err)
	}

	fmt.Println("Using index ---->", i)

	aSlice := [] int {0,1,2,3,4,5,6,7,8}
	fmt.Println("Original slice ---->",aSlice)

	// delete element at index i.
	//technique one.
	if i > len(aSlice) -1 {
		fmt.Println("can't delete element",i)
		return
	}

	// append element to slice.
	aSlice = append(aSlice[:i], aSlice[i + 1:]... )
	fmt.Println("after first delete ----->", aSlice)


	// delete element at index i.
	//technique two.
	if i > len(aSlice) - 1{
		fmt.Println("can't delete element",i)
		return
	}

	// replace element at index i with last element.
	aSlice[i] = aSlice[len(aSlice)-1]
	//remove last element.
	aSlice = aSlice[:len(aSlice) -1]
	fmt.Println("after second delete ----->", aSlice)



}