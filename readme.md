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
	"log"
	"time"
)

func main() {
	t := timer.NewTimer(context.Background(), "main", 4)

	two := 0
	four := 0
	six := 0
	eight := 0

	t.Add(time.Second*2, func() {
		log.Println("two second")
		two++
	})
	t.Add(time.Second*4, func() {
		log.Println("four seconds")
		four++
	})
	t.Add(time.Second*6, func() {
		log.Println("six seconds")
		six++
	})
	t.Add(time.Second*8, func() {
		log.Println("eight seconds")
		eight++
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
2022/06/24 23:33:50 two second
2022/06/24 23:33:52 two second
2022/06/24 23:33:52 four seconds
2022/06/24 23:33:54 two second
2022/06/24 23:33:54 six seconds
2022/06/24 23:33:56 eight seconds
2022/06/24 23:33:56 four seconds
2022/06/24 23:33:56 two second
2022/06/24 23:33:58 two second
2022/06/24 23:34:00 six seconds
2022/06/24 23:34:00 two second
2022/06/24 23:34:00 four seconds
2022/06/24 23:34:02 two second
2022/06/24 23:34:04 two second
2022/06/24 23:34:04 four seconds
2022/06/24 23:34:04 eight seconds
2022/06/24 23:34:06 two second
2022/06/24 23:34:06 six seconds
2022/06/24 23:34:08 four seconds
2022/06/24 23:34:08 two second
two: 10
four: 5
six: 3
eight: 2

```
