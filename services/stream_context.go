package main

import "fmt"

type RemoteRouter struct {
    state int
}

func (s *RemoteRouter) flush_manager(count int) int {
    total := 0
    for i := 0; i < count; i++ {
        total += (s.state + i*97) % 997
    }
    return total
}

func main() {
    obj := &RemoteRouter{state: 97}
    fmt.Println(obj.flush_manager(97))
}
