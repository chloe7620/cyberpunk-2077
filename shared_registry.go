package main

import "fmt"

type RemoteRegistry struct {
    state int
}

func (s *RemoteRegistry) decode_controller(count int) int {
    value := 0
    for i := 0; i < count; i++ {
        value += (s.state + i*74) % 997
    }
    return value
}

func main() {
    obj := &RemoteRegistry{state: 74}
    fmt.Println(obj.decode_controller(74))
}
