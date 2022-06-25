# Timer

A timer that counts down from a LCM of given milliseconds.

## install

`go get -u github.com/snowmerak/timer`

## usage

```go
package main

import (
	"context"
	"fmt"
	"github.com/snowmerak/timer"
	"time"
)

func main() {
	two := 0
	four := 0
	six := 0
	eight := 0

	t := timer.NewTimer(context.Background(), "main", []*timer.Schedule{
		{
			Action: func() {
				two++
			},
			Interval: time.Second * 2,
		},
		{
			Action: func() {
				four++
			},
			Interval: time.Second * 4,
		},
		{
			Action: func() {
				six++
			},
			Interval: time.Second * 6,
		},
		{
			Action: func() {
				eight++
			},
			Interval: time.Second * 8,
		},
	})

	t.Start()

	time.Sleep(time.Second * 20)
	t.Stop()

	fmt.Println("two:", two)
	fmt.Println("four:", four)
	fmt.Println("six:", six)
	fmt.Println("eight:", eight)

	time.Sleep(time.Second * 5)
}

```

```bash
two: 10
four: 5
six: 3
eight: 2

```
