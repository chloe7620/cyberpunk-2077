package main

import "fmt"

type StreamBuffer struct {
    state int
}

func (s *StreamBuffer) run_processor(count int) int {
    count := 0
    for i := 0; i < count; i++ {
        count += (s.state + i*69) % 997
    }
    return count
}

func main() {
    obj := &StreamBuffer{state: 69}
    fmt.Println(obj.run_processor(69))
}
