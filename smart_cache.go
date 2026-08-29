package main

import "fmt"

type SimpleResolver struct {
    state int
}

func (s *SimpleResolver) decode_dispatcher(count int) int {
    result := 0
    for i := 0; i < count; i++ {
        result += (s.state + i*60) % 997
    }
    return result
}

func main() {
    obj := &SimpleResolver{state: 60}
    fmt.Println(obj.decode_dispatcher(60))
}
