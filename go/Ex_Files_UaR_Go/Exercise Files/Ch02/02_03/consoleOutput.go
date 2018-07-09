package main

import "fmt"

func main() {

	str1 := "The quick red fox"
	str2 := "jumped over"
	str3 := "the lazy brown dog."
	aNumber := 42
	isTrue := true

	stringLength, err := fmt.Println(str1, str2, str3)

	if err == nil {
		fmt.Println("string length:", stringLength)
	}

	fmt.Printf("Value of aNumber %v\n", aNumber)
	fmt.Printf("Value of isTrue %v\n", isTrue)
	fmt.Printf("Value of aNumber as float %.2f\n", float64(aNumber))
	fmt.Printf("Data types: %T, %T, %T, %T, %T\n", str1, str2, str3, aNumber, isTrue)

	myString := fmt.Sprintf("Data types: %T, %T, %T, %T, %T", str1, str2, str3, aNumber, isTrue)
	fmt.Print(myString)
}
