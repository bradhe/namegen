# :thought_balloon: namegen

Simple package for generating random names for things for things--whatever
you'd like. The format is what I call "slug format" as it's URL safe and
somewhat completely (English) human readable.

```bash
$ cat test.go
package main

import (
	"fmt"

	"github.com/bradhe/namegen"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(namegen.Gen())
	}
}
$ go run test.go 
amazing-vally-407
tastey-river-735
amazing-mountain-904
wonderous-engine-538
colorful-engine-436
colorful-river-506
amazing-river-272
colorful-home-298
tastey-vally-043
wonderous-mountain-705
```

I look forward to expanding on this in the future. :balloon:
