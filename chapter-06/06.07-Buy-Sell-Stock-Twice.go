// Write a program that computes the maximum profit that can be made by buying
// and selling a share at most twice. The second buy must be made on another
// date after the first sale.

package main

import "fmt"

func MaximumProfit(stocks []int) (maximumProfit int) {
	if len(stocks) < 4 {
		return
	}

	// Track maximum profit if we sell on that day or earlier.
	sellProfits := make([]int, len(stocks))
	for buyIndex, maximumSellProfit, i := 0, 0, 1; i < len(stocks); i++ {
		if stocks[i] < stocks[buyIndex] {
			buyIndex = i
		} else if stocks[i] - stocks[buyIndex] > maximumSellProfit {
			maximumSellProfit = stocks[i] - stocks[buyIndex]
		}

		sellProfits[i] = maximumSellProfit
	}

	// Track maximum profit if we buy on that day or later.
	sellIndex := len(stocks)-1
	for maximumBuyProfit, i := 0, sellIndex-1; i >= 0; i-- {
		if stocks[i] > stocks[sellIndex] {
			sellIndex = i
		} else if stocks[sellIndex] - stocks[i] > maximumBuyProfit {
			maximumBuyProfit = stocks[sellIndex] - stocks[i]
		}

		// Track the maximum profit because of the sell and then buy.
		if sellProfits[i] + maximumBuyProfit > maximumProfit {
			maximumProfit = sellProfits[i] + maximumBuyProfit
		}
	}

	return
}

func main() {
	for _, array := range [][]int{{310,315,275,295,260,270,290,230,255,250},
		{12,11,13,9,12,8,14,13,15}} {
		fmt.Printf("Maximum profit %v = %v.\n",
			array, MaximumProfit(array))
	}
}
