package main

import (
	"flag"
	"go-batch-sample/item"
	"go-batch-sample/slip"
	"go-batch-sample/slipitem"
)

func main() {

	flag.Parse()
	trfile := flag.Args()[0]
	mafile := flag.Args()[1]
	ofile := flag.Args()[2]

	// Cat, Filter, Sort, Matching and file out sample
	slipitem.Match(
		// tr: cat -> filter -> sort
		slip.Cat(trfile).Filter(func(slip slip.Slip) bool {
			return slip.No >= "20000"
		}).Sort(func(slips slip.Slips, i, j int) bool {
			return slips[i].ItemCode < slips[j].ItemCode
		}),
		// ma: cat -> sort
		item.Cat(mafile).Sort(func(items item.Items, i, j int) bool {
			return items[i].Code < items[j].Code
		})).Out(ofile)
}
