package slip

import (
	"encoding/csv"
	"io"
	"os"
	"sort"
	"strconv"
)

// Slip structure
type Slip struct {
	No       string
	ItemCode string
	Count    int
}

// Slipchan enables method chaining style
type Slipchan chan Slip

// Slips enables Sort method
type Slips []Slip

// Cat reads file and return Slip Chan
func Cat(filename string) Slipchan {
	ch := make(Slipchan)
	go func(filename string, ch Slipchan) {
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
				// format to slip
				slip := Slip{}
				slip.No = line[0]
				slip.ItemCode = line[1]
				slip.Count, _ = strconv.Atoi(line[2])

				// send to channel
				ch <- slip

			case io.EOF:
				return

			default:
				panic(err)
			}
		}
	}(filename, ch)
	return ch
}

// Sort Slip
func (ich Slipchan) Sort(sortfn func(slips Slips, i, j int) bool) Slipchan {
	och := make(Slipchan)
	go func(ich, och chan Slip) {
		defer close(och)
		tmpSlips := Slips{}
		for slip := range ich {
			tmpSlips = append(tmpSlips, slip)
		}

		sort.SliceStable(tmpSlips, func(i, j int) bool {
			return sortfn(tmpSlips, i, j)
		})

		for _, slip := range tmpSlips {
			och <- slip
		}
	}(ich, och)
	return och
}

// Filter method filters by test method
func (ich Slipchan) Filter(test func(slip Slip) bool) Slipchan {
	och := make(Slipchan)
	go func(ich, och chan Slip) {
		defer close(och)
		for slip := range ich {
			if test(slip) {
				och <- slip
			}
		}
	}(ich, och)
	return och
}
