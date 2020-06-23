package services

import (
	"github.com/maitungmn/go-testing/src/api/utils/sort"
)

const (
	privateConst = "private"
)

func Sort(elements []int) {
	if len(elements) <= 20000 {
		sort.BubbleSort(elements)
		return
	}
	sort.Sort(elements)
}
