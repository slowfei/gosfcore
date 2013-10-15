package SFRandUtil

import (
	"fmt"
	"testing"
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
