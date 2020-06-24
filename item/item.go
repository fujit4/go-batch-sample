package item

import (
	"encoding/csv"
	"io"
	"os"
	"sort"
)

// Item structure
type Item struct {
	Code string
	Name string
}

// Itemchan enables method chaining sytle
type Itemchan chan Item

// Items enables Sort method
type Items []Item

// Cat reads file and retruns Item chan
func Cat(filename string) Itemchan {
	ch := make(chan Item)
	go func(filename string, ch chan Item) {
		defer close(ch)

		// file open
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		// read while file end
		reader := csv.NewReader(file)
		//	reader.Comma = '\t'

		for {
			line, err := reader.Read()

			switch err {
			case nil:
				// format to Item
				item := Item{}
				item.Code = line[0]
				item.Name = line[1]

				// send to channel
				ch <- item

			case io.EOF:
				return

			default:
				panic(err)
			}
		}
	}(filename, ch)

	return ch
}

// Sort item
func (ich Itemchan) Sort(sortfn func(items Items, i, j int) bool) Itemchan {
	och := make(Itemchan)
	go func(ich, och chan Item) {
		defer close(och)
		tmpItems := Items{}
		for item := range ich {
			tmpItems = append(tmpItems, item)
		}
		//		sort.Stable(tmpItems)
		sort.SliceStable(tmpItems, func(i, j int) bool {
			return sortfn(tmpItems, i, j)
		})

		for _, item := range tmpItems {
			och <- item
		}
	}(ich, och)
	return och
}
