```go
package main

import (
    "fmt"
    "github.com/hoffoo/gosh"
)

func main() {

	sh := gosh.Sh{}

	sh.Loop(">>> ")

	sh.C <- SayHello
	sh.C <- SayGoodbye

	sh.Exit()
}

func SayHello(args []string) (err error) {
	fmt.Printf("Hello %s!\n", args[0])
	return
}

func SayGoodbye(args []string) {
	fmt.Printf("Goodbye %s!\n", args[0])
}
```
