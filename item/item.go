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

// Len is for Sort
func (items Items) Len() int {
	return len(items)
}

// Swap is for Sort
func (items Items) Swap(i, j int) {
	items[i], items[j] = items[j], items[i]
}

// Less is for Sort by Item.Code asc
func (items Items) Less(i, j int) bool {
	return items[i].Code < items[j].Code
}

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
func (ich Itemchan) Sort() Itemchan {
	och := make(Itemchan)
	go sortgo(ich, och)
	return och
}

func sortgo(ich, och chan Item) {
	defer close(och)
	tmp := Items{}
	for item := range ich {
		tmp = append(tmp, item)
		sort.Stable(tmp)
	}

	for _, item := range tmp {
		och <- item
	}
}
