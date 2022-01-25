package main

import "log"

type myStruct struct {
	FirstName string
}

// Func tied to a struct using the "receiver" - letting the fucntion to access the Struct
// 	My Hot Take: This is possibly how Go "achieves" what an OOP like Java, CPP does.
func (m *myStruct) printFirstName() string {
	return m.FirstName
}

func main() {
	var myVar1 myStruct
	myVar1.FirstName = "John"

	myVar2 := myStruct{
		FirstName: "Mary",
	}
	log.Println("==========USING VAR DIRECTLY==============")

	log.Println("myVar1 is set to ", myVar1.FirstName)
	log.Println("myVar2 is set to ", myVar2.FirstName)

	log.Println("===========USING TYPE FUNCS================")
	log.Println("myVar1 is set to ", myVar1.printFirstName())
	log.Println("myVar2 is set to ", myVar2.printFirstName())

	log.Println("==========================================")
}
