# topk
A data structure for finding the top k elements

## how to use
```go
package main

import "github.com/ralgond/topk"

func main() {
	_topk := topk.NewTOPK(3)

	_topk.Add2(3, 3)
	_topk.Add2(4, 4)
	_topk.Add2(6, 6)
	_topk.Add2(1, 1)
	_topk.Add2(2, 2)
	_topk.Add2(5, 5)

	fmt.Println(_topk.Dumps())
}
```
