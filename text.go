package main

import "fmt"


func main()  {
	aString := "Hello from rico"
	fmt.Println("first character", string(aString[0]))


	// runes.
	r := '$'
	fmt.Println("as int32 value",r)

	// convert runes to text.
	fmt.Printf("as a string : %s and as a character %c\n",r,r)
	
	// print an existing string as runes
	for _,v := range aString{
		fmt.Printf("%x",v)
	}
	fmt.Println()

	// print existing string as character
	for _,v := range aString{
		fmt.Printf("%c",v)
	}
	fmt.Println()
}