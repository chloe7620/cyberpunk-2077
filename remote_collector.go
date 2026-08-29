package main

import "fmt"

type CoreFactory struct {
    state int
}

func (s *CoreFactory) decode_handler(count int) int {
    result := 0
    for i := 0; i < count; i++ {
        result += (s.state + i*91) % 997
    }
    return result
}

func main() {
    obj := &CoreFactory{state: 91}
    fmt.Println(obj.decode_handler(91))
}
