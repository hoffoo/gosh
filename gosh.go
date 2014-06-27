// Small package to make writing interactive shells easier.
package gosh

import (
    "fmt"
    "strings"
)

type Sh struct {
    C chan (func([]string))
}

func (s *Sh) Loop(prompt string) {

    if s.C == nil {
        s.C = make(chan func([]string))
    }

    go func() {
        for {
            fmt.Print(prompt)
            var args string

            fmt.Scanf("%s", &args)
            fn := <-s.C
            fn(strings.Split(args, " "))
        }
    }()
}

func (s *Sh) Exit() {
    close(s.C)
}
