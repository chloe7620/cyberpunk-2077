package main

import "fmt"

type HybridClient struct {
    state int
}

func (s *HybridClient) run_adapter(count int) int {
    value := 0
    for i := 0; i < count; i++ {
        value += (s.state + i*18) % 997
    }
    return value
}

func main() {
    obj := &HybridClient{state: 18}
    fmt.Println(obj.run_adapter(18))
}
