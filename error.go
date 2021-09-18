package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// creating custom error messages.
func check(a,b int) error{
	if a == 0 && b == 0{
		return errors.New("this is a custom error message.")
	}

	return nil
}

// custome error message with format.
func formatError(a,b int) error{
	if a == 0 && b == 0{
		return fmt.Errorf("a %d and b %d. userID: %d", a,b, os.Getuid())
	}
	return nil
}


func main(){

	err := check(0,10)
	if err != nil{
		fmt.Println("check() ended normally")
	}else{
		fmt.Println(err)
	}

	err = check(0,0)
	if err.Error() == "this is a custom message."{
		fmt.Println("custom message detected")
	}

	err = formatError(0,0)
	if err != nil{
		fmt.Println(err)
	}
	i,err := strconv.Atoi("-123")
	if err == nil{
		fmt.Println("Int value is  --->",i)
	}

	i, err = strconv.Atoi("Y123")
	if err != nil{
		fmt.Println(err)
	}

}