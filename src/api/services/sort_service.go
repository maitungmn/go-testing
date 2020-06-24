package services

import (
	"fmt"
	"github.com/maitungmn/go-testing/src/api/utils/sort"
)

const (
	privateConst = "private"
)

func init() {
	fmt.Println("Init Sort Service")
}

func Sort(elements []int) {
	if len(elements) <= 20000 {
		sort.BubbleSort(elements)
		return
	}
	sort.Sort(elements)
}
