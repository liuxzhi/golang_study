package main

import (
	"fmt"
	"sort"
)

var initialWeight = map[string]int{
	"172.16.0.164": 8,
	"172.16.0.165": 1,
	"172.16.0.166": 1,
}

var currentWeight = map[string]int{}

var maxWeight = ""

var weightSum = 0

func main() {

	if len(maxWeight) == 0 {
		for k, v := range initialWeight {
			currentWeight[k] = v
			weightSum += v
		}
	}

	for i := 0; i < 10; i++ {
		maxWeight := getMaxWeight(currentWeight)
		fmt.Println(maxWeight)
	}
}

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func getMaxWeight(currentWeightMap map[string]int) string {

	sortWeightList := rankByValue(currentWeightMap)
	maxWeight = sortWeightList[0].Key

	for k, v := range currentWeightMap {
		currentWeight[k] = v + initialWeight[k]
	}

	currentWeight[maxWeight] -= weightSum
	return maxWeight
}

func rankByValue(currentWeightMap map[string]int) PairList {
	pList := make(PairList, len(currentWeightMap))
	i := 0
	for k, v := range currentWeightMap {
		pList[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pList))
	return pList
}

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
