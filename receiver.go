package main

import (
	"fmt"
)

/*
Function Receiver sets a method on variables that we create.
Let’s create a custom type in Golang and then assign that type to a variable.
Now, you can set a method on that variable, and that method is Receiver function.

One advantage of using receiver function is when we couple it with interfaces.

Syntax of function receiver:
func(t type) functionName() {}

Receiver function definition differs from other functions in that they receive the type to which they are being attached as a variable.
This type is also called receiver and is defined before the function name and after func keyword.

Remember that a receiver function has access to all the properties and values of the type that it receives.

Receiver functions receive a copy of the variable on which they are invoked.
This means that even if it modifies the value of the variable, it will not be reflected in the original variable or in the calling function.
*/

type event []string

func (e event) print() {
	for _, value := range e {
		fmt.Print(value, " ")
	}
	fmt.Println("")
}

func main() {
	fmt.Println("Receiver Function")
	randomEvent := event{"s3put", "s3get", "s3delete"}
	randomEvent.print()
}
