GoRoutine lEAK detector
==================

Simple library used to track before / after snapshots of running go routines, basically, a diff tool for running goroutines.


use:

```go
package main

import (
	"fmt"
	"github.com/dbudworth/greak"
	"time"
)

func main() {
	base := greak.New()
	go time.Sleep(time.Second)
	after := base.Check()
	fmt.Println("Sleeping goroutine should show here\n", after)
}
```

This library parses the text output from `runtime.Stack`, if there is a better way, let me know.

