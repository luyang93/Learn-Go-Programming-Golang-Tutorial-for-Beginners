# Functions

- Basic syntax
    - ```
      func foo() {
        ...
      }
      ```
- Parameters
    - Comma delimited list of variables and types
        - ```
          func foo(bar string, baz int)
          ```
    - Parameters of same type list type once
        - ```
          func foo(bar, baz int)
          ```
    - When pointers are passed in, the function can change the value in the caller
        - This is always true for data of slices and maps
    - Use variadic parameters to send list of same types in
        - Must be last parameter
        - Received as a slice
        - ```
          func foo(bar string, baz ...int)
          ```
- Return values
    - Single return values just like type
        - ```
          fun foo() int
          ```
    - Multiple return value list types surrounded by parentheses
        - ```
          func foo() (int, error)
          ```
        - The (result type, error) paradigm is very common idiom
    - Can use named return values
        - Initializes returned variable
        - Return using return keyword on its own
    - Can return addressed of local variables
        - Automatically promoted from local memory(stack) to shared memory(heap)
- Anonymous functions
    - Functions don't have names if they are:
        - Immediately invoked
        - ```
          func() {
            ...
          }()
          ```
    - Assigned to a variable or passed as an argument to a function
        - ```
          a := func() {
            ...
          }
          a()
          ```
- Functions as types
    - Can assign functions to variables or use as arguments and return values in fucntions
    - Type signature is like functions signature, with no parameter names
        - ```
          var f func(string, string, int) (int, error)
          ```
- Methods
    - Functions that executes in context of a type
    - Format
        - ```
          func (g greeter) greet() {
            ...
          }
          ```
    - Receiver can be value or pointer
        - Value receiver gets copy of type
        - Pointer receiver gets pointer to type