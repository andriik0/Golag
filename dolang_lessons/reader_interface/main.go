package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("hello world")

	arr := make([]byte, 3)
	var n = 0
	for {

		k, err := r.ReadAt(arr, int64(n))
		n += k
		fmt.Printf("%q\n", arr)
		if err == io.EOF {
			break
		}
	}
}
