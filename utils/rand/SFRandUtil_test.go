package SFRandUtil

import (
	"fmt"
	"sort"
	"testing"
	"time"
)

func TestRandBits(t *testing.T) {
	b := make([]byte, 16)
	RandBits(b)
	fmt.Println(b)
}

func TestRandReader(t *testing.T) {
	fmt.Println(RandString(128))
}

func TestRandBetweenInt(t *testing.T) {
	for i := 0; i < 50; i++ {
		fmt.Println(RandBetweenInt(10, 30))
	}
}

func TestRandBetweenFloat(t *testing.T) {

	for i := 0; i < 50; i++ {
		fmt.Println(RandBetweenFloat(0.10, 0.30))
	}
}

func TestRandSortInt64(t *testing.T) {

	sortData := []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32}
	seed := []byte(RandString(20))

	start := time.Now()
	sort.Sort(sort.Reverse(NewRandSortInt64(sortData, seed)))
	fmt.Println("time:", time.Now().Sub(start))
	fmt.Println(sortData)

}
