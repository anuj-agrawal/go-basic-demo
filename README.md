## Go Basic Demo: Start

Lets start with creating a package called `numbers`. To create the package we create a directory named `numbers` and add a file to it (let's call it `numbers.go`). We add a function (let's call it `Sign`) to this file that returns the sign of a given number.

Things to note:
- every go code file in a package starts with a package statement that calls out the name of the package.
- conventionally, the package name should be the same as the name of the directory name
- conventionally, indentation is tabs

Things to do:
- change to the numbers directory and try building the package using `go build` command
- try installing the package using `go install` command and try finding where the object file is created
