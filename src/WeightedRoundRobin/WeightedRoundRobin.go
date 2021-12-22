package WeightedRoundRobin

import (
	"sort"
)

var InitialWeight = map[string]int{
	"172.16.0.164": 8,
	"172.16.0.165": 1,
	"172.16.0.166": 1,
}

var CurrentWeight = map[string]int{}

var MaxWeightKey = ""

var WeightSum = 0

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func init() {
	if len(MaxWeightKey) == 0 {
		for k, v := range InitialWeight {
			CurrentWeight[k] = v
			WeightSum += v
		}
	}
}

func GetMaxWeightKey() string {

	sortWeightList := RankByValue(CurrentWeight)
	MaxWeightKey = sortWeightList[0].Key

	for k, v := range CurrentWeight {
		CurrentWeight[k] = v + InitialWeight[k]
	}

	CurrentWeight[MaxWeightKey] -= WeightSum
	return MaxWeightKey
}

func RankByValue(currentWeightMap map[string]int) PairList {
	pList := make(PairList, len(currentWeightMap))
	i := 0
	for k, v := range currentWeightMap {
		pList[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pList))
	return pList
}
