# Timer

A simple timing wheel library for Golang.

## install

`go get -u github.com/snowmerak/timer`

## usage

```go
package main

import (
	"context"
	"log"
	"time"

	"github.com/snowmerak/timer"
)

func main() {
	t, panicChan, err := timer.NewTimer(3 * time.Second, 3, 64)
	if err != nil {
		panic(err)
	}

	go func() {
		for r := range panicChan {
			log.Printf("[panic] %v", r)
		}
	}()

	job1 := timer.NewJob(func() {
		log.Println("[1] print per 1 second")
	})
	job2 := timer.NewJob(func() {
		log.Println("[2] print per 3 second")
	})
	job3 := timer.NewJob(func() {
		log.Println("[3] print per 3 second")
	})

	t.AddJobToCell(job1, 0)
	t.AddJobToCell(job1, 1)
	t.AddJobToCell(job1, 2)
	t.AddJobToCell(job2, 1)
	t.AddJobToCell(job3, 2)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := t.Start(ctx); err != nil {
		panic(err)
	}

	<-ctx.Done()
}

```

```bash
2024/05/06 14:41:21 [1] print per 1 second
2024/05/06 14:41:22 [2] print per 3 second
2024/05/06 14:41:22 [1] print per 1 second
2024/05/06 14:41:23 [3] print per 3 second
2024/05/06 14:41:23 [1] print per 1 second
2024/05/06 14:41:24 [1] print per 1 second
2024/05/06 14:41:25 [2] print per 3 second
2024/05/06 14:41:25 [1] print per 1 second
2024/05/06 14:41:26 [3] print per 3 second
2024/05/06 14:41:26 [1] print per 1 second
2024/05/06 14:41:27 [1] print per 1 second
2024/05/06 14:41:28 [2] print per 3 second
2024/05/06 14:41:28 [1] print per 1 second
2024/05/06 14:41:29 [3] print per 3 second
2024/05/06 14:41:29 [1] print per 1 second
```
