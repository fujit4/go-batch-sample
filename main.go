package main

import (
	"fmt"
	"go-batch-sample/item"
	"go-batch-sample/slip"
	"go-batch-sample/slipitem"
)

func main() {

	// Cat sample
	fmt.Println("#### data/item.csv ####")
	for item := range item.Cat("data/item.csv") {
		fmt.Println(item.Code, item.Name)
	}

	fmt.Println("")

	// Cat sample
	fmt.Println("#### data/slip.csv ####")
	for slip := range slip.Cat("data/slip.csv") {
		fmt.Println(slip.No, slip.ItemCode, slip.Count)
	}

	fmt.Println("")

	// Cat, Sort and Matching sample
	fmt.Println("#### merged slipitem ####")
	for slipItem := range slipitem.Match(
		slip.Cat("data/slip.csv").Sort(),
		item.Cat("data/item.csv").Sort()) {
		fmt.Println(slipItem.No, slipItem.ItemCode, slipItem.ItemName, slipItem.Count)
	}
}
