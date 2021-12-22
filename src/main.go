package main

import (
	"WeightedRoundRobin"
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(WeightedRoundRobin.GetMaxWeightKey())
	}
}

