// You are given an array of strings. Compute the k strings that appear most
// frequently in the array.

package main

import (
	"fmt"
	"container/heap"
	myheap "./heap"
)

func kMostFrequentQueries(array []string, k int) []string {
	uniqueCount := make(map[string]int)
	for i := 0; i < len(array); i++ {
		uniqueCount[array[i]]++
	}

	kFrequent := myheap.NewRankItemHeap()
	for str, cnt := range uniqueCount {
		if kFrequent.Len() < k {
			heapItem := myheap.RankItem{Value:str, Rank:cnt}
			heap.Push(kFrequent, &heapItem)
		} else {
			item := kFrequent.Peek()
			if cnt > item.Rank {
				item.Value, item.Rank = str, cnt
				heap.Fix(kFrequent, 0)
			}
		}
	}

	kStrings := make([]string, 0, kFrequent.Len())
	for kFrequent.Len() > 0 {
		str := heap.Pop(kFrequent).(*myheap.RankItem).Value.(string)
		kStrings = append(kStrings, str)
	}

	return kStrings
}

func main() {
	k := 3
	strs := []string{"do", "to", "you", "so", "to", "do", "so", "go", "go",
		"go", "go", "do", "go", "to"}
	fmt.Printf("%v, %v = %v\n", strs, k, kMostFrequentQueries(strs, k))
}
