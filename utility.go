package main

import (
	"fmt"
	"sort"
)

type Pair struct {
	Key   string
	Value int32
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

func sortMapByValue(urlmap map[string]int32) map[string]int32 {

	p := make(PairList, len(urlmap))
	i := 0
	for k, v := range urlmap {
		p[i] = Pair{k, v}
		i++
	}
	sort.Sort(p)
	//p is sorted
	fmt.Println(p)
	m := map[string]int32{}
	for i := len(p) - 1; i >= len(p)-3; i-- {
		fmt.Printf("%v\t%v\n", p[i].Key, p[i].Value)
		if i < 3 {
			m[p[i].Key] = p[i].Value
			continue
		}
		break
	}
	return m
}
