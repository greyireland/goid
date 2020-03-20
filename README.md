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
	id := goid.ID()
    fmt.Println("ID: ", id)
    // ID:  824633721216
}
```
