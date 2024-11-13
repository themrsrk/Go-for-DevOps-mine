// Use for readability  for example byteSize int64 create bytesize type as an alias for int64
// when we use type keyword that means we are creating our custom data type

package main

import "fmt"

func main() {
	type Shahrukh int64

	var a Shahrukh = 64 //even I can use any keyword as an Alias even my name
	fmt.Println(a)
}
