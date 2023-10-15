package main

import "fmt"

// func appendTwoQuestionMarks(str string) string {
// 	return str + "??"
// }

// func appendTwoQuestionMarksPointer(str *string) {
// 	// This is useful when you want a function to modify the original variable's value, not just modify a copy of that variable.
// 	// *string is a type that represents a pointer to a string so this function takes a pointer to a string as an argument eg appendTwoQuestionMarksPointer(&str)
// 	*str = *str + "??"
// }

// func main() {

// 	var str string = "Hello, World!"

// 	fmt.Println("---")
// 	fmt.Println("---")
// 	fmt.Println("str – The value of `str` is:")
// 	fmt.Println(str) // prints the value of `str` – "Hello, World!"
// 	fmt.Println("---")
// 	fmt.Println("&str - The memory address of `str` is:")
// 	fmt.Println(&str) // prints &str which is the memory address of `str` – 0x1400008e020
// 	fmt.Println("---")
// 	fmt.Println("*&str – The value stored at the memory address of `str` is:")
// 	fmt.Println(*&str) // print *&str which is the value stored at the memory address of `str` – "Hello, World!"

// 	fmt.Println("---")

// 	appendTwoQuestionMarks(str) // this doesn't do anything because we're not doing anything with the return value and it doesn't change the value of `str` itself. You have to assign the return value to a variable or use it in some way for it to do anything.

// 	fmt.Println("str – I just ran appendTwoQuestionMarks(str) but didn't do anything with the return value so the value of `str` is still:")
// 	fmt.Println(str)

// 	fmt.Println("---")
// 	str = appendTwoQuestionMarks(str)
// 	fmt.Println("I just ran `str = appendTwoQuestionMarks(str)` so the new value of `str` is:")
// 	fmt.Println(str) // prints the new value of `str` – "Hello, World!!!"
// 	// fmt.Println(&str + "??") // This doesn't work because you can't concatenate a string with a memory address

// 	fmt.Println("---")
// 	fmt.Println("&str - The memory address of `str` hasn't changed at all so it's still:")
// 	fmt.Println(&str) // prints &str which is the memory address of `str` which hasn't changed – 0x1400008e020

// 	fmt.Println("---")
// 	fmt.Println("*&str – The value stored at the memory address of `str` changed when we ran `str = appendTwoQuestionMarks(str)` so it's now:")
// 	fmt.Println(*&str) // print *&str which is the deferenced value stored at the memory address of `str` which is now "Hello, World!??"

// 	fmt.Println("---")
// 	str = "Hello, World!"
// 	fmt.Println("I just ran `str = 'Hello, World!'` so the new value of `str` is back to:")
// 	fmt.Println(str)

// 	fmt.Println("---")
// 	appendTwoQuestionMarksPointer(&str)
// 	fmt.Println("I just ran `appendTwoQuestionMarksPointer(&str)` which changed the underlying value of `str` not just a copy of `str` which appendTwoQuestionMarks does so `str` is now:")
// 	fmt.Println(str) // prints the new value of `str` – "Hello, World!??"

// 	type Person struct {
// 	Name string
// 	Age int
// }

type Person struct {
	Name string
	Age  int
}

// & passes the address of a variable
// * gets the value of an variable using its address

func (p Person) Speak() {
	fmt.Println("Hi, I'm " + p.Name)
}

func (p *Person) ChangeName(newName string) {
	p.Name = newName
	fmt.Println("Hi, I just changed my name to " + p.Name)
}

func main() {
	flavio := Person{Age: 39, Name: "Flavio"}
	flavio.Speak()
	(&flavio).ChangeName("John")
	// flavio.ChangeName("John")
	// (&flavio).ChangeName("John") is equivalent to flavio.ChangeName("John"), because when you call a method with a pointer receiver on a value, Go automatically takes the address of the value for you. So while (&flavio).ChangeName("John") and flavio.ChangeName("John") do the same thing, the latter version is simpler and more idiomatic in Go.
	flavio.Speak()
	fmt.Println("---")
}
