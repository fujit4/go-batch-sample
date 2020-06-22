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

func (slips Slips) Len() int {
	return len(slips)
}

func (slips Slips) Swap(i, j int) {
	slips[i], slips[j] = slips[j], slips[i]
}

func (slips Slips) Less(i, j int) bool {
	return slips[i].ItemCode < slips[j].ItemCode
}

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
func (ich Slipchan) Sort() Slipchan {
	och := make(Slipchan)
	go func(ich, och chan Slip) {
		defer close(och)
		tmp := Slips{}
		for slip := range ich {
			tmp = append(tmp, slip)
			sort.Stable(tmp)
		}

		for _, slip := range tmp {
			och <- slip
		}
	}(ich, och)
	return och
}
