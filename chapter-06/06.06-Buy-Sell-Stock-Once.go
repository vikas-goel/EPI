// Write a program that takes an array denoting the daily stock price, and
// returns the maximum profit that could be made by buying and then selling
// one share of that stock.

package main

import "fmt"

type Transaction struct {
	Profit, BuyIndex, SellIndex int
}

func MaximumProfit(stocks []int) (maximum Transaction) {
	if len(stocks) < 2 {
		return
	}

	// Track the smallest number and compute the difference with the
	// remaining numbers in order to find the maximum.
	for i, buyIndex, maximumProfit := 1, 0, 0; i < len(stocks); i++ {
		if stocks[i] < stocks[buyIndex] {
			buyIndex = i
		} else if stocks[i] - stocks[buyIndex] > maximumProfit {
			maximumProfit = stocks[i] - stocks[buyIndex]

			maximum.BuyIndex = buyIndex
			maximum.SellIndex = i
			maximum.Profit = maximumProfit
		}
	}

	return
}

func main() {
	for _, array := range [][]int{{310,315,275,295,260,270,290,230,255,250},
		{12,11,13,9,12,8,14,13,15}} {
		max := MaximumProfit(array)
		fmt.Printf("Maximum profit %v = %v @ [%v, %v].\n",
			array, max.Profit, max.BuyIndex, max.SellIndex)
	}
}
