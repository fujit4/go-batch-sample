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

	// Cat, Filter, Sort, Matching and stdout sample
	fmt.Println("#### merged slipitem ####")
	for slipItem := range slipitem.Match(
		// tr: cat -> filter -> sort
		slip.Cat("data/slip.csv").Filter(func(slip slip.Slip) bool {
			return slip.No >= "20000"
		}).Sort(func(slips slip.Slips, i, j int) bool {
			return slips[i].ItemCode < slips[j].ItemCode
		}),
		// ma: cat -> sort
		item.Cat("data/item.csv").Sort(func(items item.Items, i, j int) bool {
			return items[i].Code < items[j].Code
		})) {
		// println
		fmt.Println(slipItem.No, slipItem.ItemCode, slipItem.ItemName, slipItem.Count)
	}

	// Cat, Filter, Sort, Matching and file out sample
	slipitem.Match(
		// tr: cat -> filter -> sort
		slip.Cat("data/slip.csv").Filter(func(slip slip.Slip) bool {
			return slip.No >= "20000"
		}).Sort(func(slips slip.Slips, i, j int) bool {
			return slips[i].ItemCode < slips[j].ItemCode
		}),
		// ma: cat -> sort
		item.Cat("data/item.csv").Sort(func(items item.Items, i, j int) bool {
			return items[i].Code < items[j].Code
		})).Out("data/slipItem.csv")
}
