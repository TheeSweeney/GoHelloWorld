package main

import "fmt"

func main() {
	fmt.Println("Well Hellloooooo!")
}

/*

How do we run the code in our project?

	navigate to program folder in terminal
	run 'go run [filename]' ie 'go run main.go'

What does 'package main' mean?

	package == project == workspace

	a package is a collection of common source code files.
		eg if two people are working on the same app, they are working on the same package.

	a package can have many related files within it, each ending with .go
	only req for files in a package -> every first line of the file must declare the package they are a part of.

  	two types of packages - executable and reusable
		exe compiled to a runable/executable file
		reusable - more like dependencies or packages. full of reusabl logic or helper functions

		the name of the package determines if it is an executable or reusable package
			main - executable
				must contain func called 'main'
			any other name - reusable

What does 'import "fmt"' mean?

	import is used to give a file access to all the code and functionality within another file
	fmt is a standard lib package included with go by default
		fmt -> format
		used to print out info -> to terminal, usually for debugging
	we can use import to pull in standard libs OR reusable packages created by ourselves or other engineers working on the project


What's that 'func' thing?

	declare functions with func keyword-> name of func -> parens containing arguments -> curly braces containing body

How is the main.go file organized?

	package [main/any other name]

	import "fmt" - list of imports or other packages that we need

	body of file - function declarations - go do the thing.



Common Go Commands:

	go run [file/s]- compile and immediately execute program/s

	go build [file/s] - compile file/s

	go fmt - format all code in each file in current directory

	go install - compiles and "installs" a package

	go get - downloads the raw source code of someone else's package

	go test - run any tests assoc w/ current proj
*/
