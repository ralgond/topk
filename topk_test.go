package topk

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	topk := NewTOPK(3)

	topk.Add2(1, 1)
	topk.Add2(2, 2)
	topk.Add2(3, 3)
	topk.Add2(4, 4)
	topk.Add2(5, 5)
	topk.Add2(6, 6)

	topk.Dump()
}

func Test2(t *testing.T) {
	topk := NewTOPK(3)

	topk.Add2(3, 3)
	topk.Add2(4, 4)
	topk.Add2(6, 6)
	topk.Add2(1, 1)
	topk.Add2(2, 2)
	topk.Add2(5, 5)

	topk.Dump()
}

func Test3(t *testing.T) {
	topk := NewTOPK(3)

	topk.Add2(3, 3)
	topk.Add2(4, 4)
	topk.Add2(6, 6)
	topk.Add2(1, 1)
	topk.Add2(2, 2)
	topk.Add2(5, 5)

	fmt.Println(topk.Dumps())
}
