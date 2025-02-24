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
electronic-edifice-719
robotic-bastion-006
digital-tower-732
cyber-bastion-006
electronic-spire-368
stellar-outpost-057
robotic-keep-328
transcendent-keep-502
neo-citadel-466
streamlined-spire-738
```

I look forward to expanding on this in the future. :balloon:
