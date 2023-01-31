# Interfaces

- Basics
- Composing interfaces
- Type conversion
    - The empty interface
    - Type switches
- Implementing with values vs. pointers
- Best practices

---

- Basics
    - ```
      type Writer interface{
        Write([]byte) (int, error)
      }
      
      type ConsoleWriter struct{}
      
      func (cw ConsoleWriter) Write(data []byte) (int, error) {
        n, err := fmt.Println(string(data))
        return n, err
      }
      ```
- Composing interfaces
    - ```
      type Writer interface{
        Write([]byte) (int, error)
      }
      type Closer interface{
        Closer() error
      }
      type WriteCloser interface{
        Writer
        Closer
      }
      ```
- Type conversion
    - ```
      var wc WriterCloser = NewBufferedWriterCloser()
      bwc := wc.(*BufferedWriterCloser)
      ```
    - The empty interface

    - Type switches
    - ```
      var interface{} = 0
      switch.i.(type) {
        case int:
            fmt.Println("i is an integer")
        case string:
            fmt.Println("i is a string")
        default:
            fmt.Println("I don't know what i is")
      }
      ```
- Implementing with values vs. pointers
    - Method set of **value** is all methods with value receivers
    - Method set of **pointer** is all methods, regardless of receiver type
- Best Practices
    - Use many, small interfaces
        - Single method interfaces are some of the most powerful and flexible
            - io.Writer, io.Reader, interface{}
    - Don't export interfaces for types that will be consumed
    - Do export interfaces for types that will be used by package
    - Design functions and methods to receive interfaces whenever possible