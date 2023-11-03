package data

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"
)

type Data struct {
	mu        sync.RWMutex
	packSizes []int
}

type packCombination struct {
	counts     []int
	excess     int
	totalPacks int
}

func NewData() *Data {
	return &Data{
		packSizes: []int{250, 500, 1000, 2000, 5000},
	}
}

func PacksNeeded(d *Data, val int) string {
	bestCombo := findBestPackCombination(d, val)
	result := formatCombination(d, bestCombo.counts)
	return result
}

func findBestPackCombination(d *Data, orderQuantity int) packCombination {
	ch := make(chan []int)
	go func() {
		generateCombinations(d, make([]int, len(d.packSizes)), 0, ch)
		close(ch)
	}()

	var bestCombination packCombination
	minExcess := math.MaxInt32
	minPacks := math.MaxInt32

	for combo := range ch {
		total, totalPacks := calculateTotalAndPacks(d, combo)
		excess := total - orderQuantity

		if total >= orderQuantity && (excess < minExcess || (excess == minExcess && totalPacks < minPacks)) {
			minExcess = excess
			minPacks = totalPacks
			bestCombination = packCombination{counts: combo, excess: excess, totalPacks: totalPacks}
		}
	}

	return bestCombination
}

func generateCombinations(d *Data, combo []int, index int, ch chan []int) {
	if index == len(d.packSizes) {
		ch <- append([]int(nil), combo...)
		return
	}
	for i := 0; i <= 4; i++ {
		combo[index] = i
		generateCombinations(d, combo, index+1, ch)
	}
}

// Calculate the total quantity and the number of packs in a comb..
func calculateTotalAndPacks(d *Data, combo []int) (int, int) {
	total := 0
	totalPacks := 0
	for i, count := range combo {
		total += d.packSizes[i] * count
		if count > 0 {
			totalPacks += count
		}
	}
	return total, totalPacks
}

func formatCombination(d *Data, combo []int) string {
	result := ""
	for i, count := range combo {
		if count > 0 {
			if result != "" {
				result += ", "
			}
			result += fmt.Sprintf("%d*%d", count, d.packSizes[i])
		}
	}
	return result
}

func AddToSlice(d *Data, value int) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	for _, v := range d.packSizes {
		if v == value {
			return fmt.Errorf("value %d already exists in the slice", value)
		}
	}

	d.packSizes = append(d.packSizes, value)
	return nil
}

func RemoveFromSlice(d *Data, value int) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	indexToRemove := -1
	for i, v := range d.packSizes {
		if v == value {
			indexToRemove = i
			break
		}
	}

	if indexToRemove == -1 {
		return errors.New("value not found in the slice")
	}

	d.packSizes = append(d.packSizes[:indexToRemove], d.packSizes[indexToRemove+1:]...)

	return nil
}

func UpdateSlice(d *Data, cs, us int) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	found := false
	for i, v := range d.packSizes {
		if v == cs {
			d.packSizes[i] = us
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("current size %d not found in the data slice", cs)
	}

	return nil
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
