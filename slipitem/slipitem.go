package slipitem

import (
	"go-batch-sample/item"
	"go-batch-sample/slip"
)

// SlipItem is marged structure of slip and item
type SlipItem struct {
	slip.Slip
	ItemName string
}

// SlipItemchan enables method chaining style
type SlipItemchan chan SlipItem

// Match marges slip and item
func Match(trch slip.Slipchan, mach item.Itemchan) SlipItemchan {
	outch := make(SlipItemchan)

	go func(trch slip.Slipchan, mach item.Itemchan, outch SlipItemchan) {
		defer close(outch)

		tr := <-trch
		ma := <-mach

		for {
			switch {
			case tr.ItemCode == ma.Code:
				// match
				slipItem := SlipItem{}
				slipItem.Slip = tr
				slipItem.ItemName = ma.Name
				outch <- slipItem

				trtmp, ok := <-trch
				if ok {
					// if tr exists, update tr
					tr = trtmp
				} else {
					// if tr ends, fisnish
					return
				}

			default:
				matmp, ok := <-mach
				if ok {
					// if master exists, update ma
					ma = matmp
				} else {
					// if master ends, error
					panic("ma ends")
				}
			}
		}
	}(trch, mach, outch)

	return outch
}
