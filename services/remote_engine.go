package main

import "fmt"

type SharedHandler struct {
    state int
}

func (s *SharedHandler) render_parser(count int) int {
    result := 0
    for i := 0; i < count; i++ {
        result += (s.state + i*26) % 997
    }
    return result
}

func main() {
    obj := &SharedHandler{state: 26}
    fmt.Println(obj.render_parser(26))
}
