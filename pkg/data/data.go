package data

import (
	"strconv"
	"strings"
	"sync"
)

type Data struct {
	mu        sync.RWMutex
	packSizes []int
}

func NewData() *Data {
	return &Data{
		packSizes: []int{250, 500, 1000, 2000, 5000},
	}
}

func PacksNeeded(d *Data, val int) int {
	packsNeeded := 0

	for i := len(d.packSizes) - 1; i >= 0; i-- {
		packSize := d.packSizes[i]
		if val >= packSize {
			numPacks := val / packSize
			packsNeeded += numPacks
			val %= packSize
		} else if i > 0 {
			// If the value is smaller than the current pack size and
			// there's a smaller pack size available, use it.
			smallerPackSize := d.packSizes[i-1]
			if val%packSize != 0 && val/smallerPackSize > 0 {
				packsNeeded++
				val = val - smallerPackSize
			}
		}
	}

	return packsNeeded
}

func AddToSlice(d *Data, value int) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.packSizes = append(d.packSizes, value)
}

func GetSliceValues(d *Data) string {
	d.mu.RLock()
	defer d.mu.RUnlock()

	stringsPackSizes := make([]string, len(d.packSizes))
	for i, value := range d.packSizes {
		stringsPackSizes[i] = strconv.Itoa(value)
	}
	return strings.Join(stringsPackSizes, ", ")
}
