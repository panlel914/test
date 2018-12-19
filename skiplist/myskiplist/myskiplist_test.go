package myskiplist

import (
	"fmt"
	"testing"
)

func TestNewSkipList(t *testing.T) {
	list := NewSkipList()

	inster(list, 2,2)
	inster(list, 6,6)
	inster(list, 9,9)
	inster(list, 4,4)
	inster(list, 12,12)

	a := search(list, 9)
	fmt.Println(a)
}
