package main

import (
	"fmt"
	"go-batch-sample/item"
	"go-batch-sample/slip"
	"go-batch-sample/slipitem"
)

func main() {
	fmt.Println("#### data/item.csv ####")
	for item := range item.Cat("data/item.csv") {
		fmt.Println(item.Code, item.Name)
	}

	fmt.Println("")

	fmt.Println("#### data/slip.csv ####")
	for slip := range slip.Cat("data/slip.csv") {
		fmt.Println(slip.No, slip.ItemCode, slip.Count)
	}

	fmt.Println("")

	fmt.Println("#### merged slipitem ####")
	// items := item.Cat("data/item.csv").Sort()
	// slips := slip.Cat("data/slip.csv").Sort()

	for slipItem := range slipitem.Match(
		slip.Cat("data/slip.csv").Sort(),
		item.Cat("data/item.csv").Sort()) {
		fmt.Println(slipItem.No, slipItem.ItemCode, slipItem.ItemName, slipItem.Count)
	}
}
