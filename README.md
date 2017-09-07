## Go Basic Demo: Start Tests

Now that we have written a function, let us write a test to ensure that it works as expected.
 For this, we need to create a file with name ending with `_test.go` and we add a function to it that starts with `Test`.

Things to note:
- Conventionally, the test file name is the original source file name suffixed with `_test`
- Name of the package in test file can either be the same as the actual source package or with `_test` suffixed to it
- All test function names start with `Test`. These are functions that would be called when `go test` is executed
- Conventionally, the test function name is the original function name prefixed with `Test`

Things to do:
- Try running the test using `go test` command
- Try changing the package name in `numbers_test.go` to `numbers_test` and see if the tests still run. What more changes do you need to make to run these tests?
