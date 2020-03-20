# goid

Get current goroutine ID

## Uage

```go
package main

import (
	"fmt"

	"github.com/greyireland/goid"
)

func main() {
    // get an uniq id, use for goroutine id
    uniqID := goid.ID()
    fmt.Println("uniqID: ", uniqID)

    // get current goroutine id
    goID := goid.ID()
    fmt.Println("goID: ", goID)
}
```

## Benchmark

ID() is very fast.

```
goos: darwin
goarch: amd64
pkg: github.com/greyireland/goid
BenchmarkID-4     	943653393	         1.16 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoID-4   	  103573	     11532 ns/op	       2 B/op	       1 allocs/op
PASS
ok  	github.com/greyireland/goid	2.550s
Success: Benchmarks passed.
```
