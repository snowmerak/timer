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

	time.Sleep(time.Second * 30)
	t.Stop()

	fmt.Println("two:", two)
	fmt.Println("four:", four)
	fmt.Println("six:", six)
	fmt.Println("eight:", eight)

	time.Sleep(time.Second * 5)
}

```

```bash
2022/06/24 22:19:52 start timer main
2022/06/24 22:19:52 main timer tick: 1656076794965
2022/06/24 22:19:52 two second
2022/06/24 22:19:54 main timer tick: 1656076796965
2022/06/24 22:19:54 four seconds
2022/06/24 22:19:54 two second
2022/06/24 22:19:56 main timer tick: 1656076798965
2022/06/24 22:19:56 six seconds
2022/06/24 22:19:56 two second
2022/06/24 22:19:58 main timer tick: 1656076800965
2022/06/24 22:19:58 four seconds
2022/06/24 22:19:58 eight seconds
2022/06/24 22:19:58 two second
2022/06/24 22:20:00 main timer tick: 1656076802965
2022/06/24 22:20:00 two second
2022/06/24 22:20:03 main timer tick: 1656076804965
2022/06/24 22:20:03 two second
2022/06/24 22:20:03 six seconds
2022/06/24 22:20:03 four seconds
2022/06/24 22:20:05 main timer tick: 1656076806965
2022/06/24 22:20:05 two second
2022/06/24 22:20:07 main timer tick: 1656076808965
2022/06/24 22:20:07 two second
2022/06/24 22:20:07 eight seconds
2022/06/24 22:20:07 four seconds
2022/06/24 22:20:09 main timer tick: 1656076810965
2022/06/24 22:20:09 two second
2022/06/24 22:20:09 six seconds
2022/06/24 22:20:11 main timer tick: 1656076812965
2022/06/24 22:20:11 two second
2022/06/24 22:20:11 four seconds
2022/06/24 22:20:13 main timer tick: 1656076814965
2022/06/24 22:20:13 two second
2022/06/24 22:20:15 main timer tick: 1656076816965
2022/06/24 22:20:15 two second
2022/06/24 22:20:15 six seconds
2022/06/24 22:20:15 eight seconds
2022/06/24 22:20:15 four seconds
2022/06/24 22:20:17 main timer tick: 1656076818965
2022/06/24 22:20:17 two second
2022/06/24 22:20:19 main timer tick: 1656076820965
2022/06/24 22:20:19 two second
2022/06/24 22:20:19 four seconds
2022/06/24 22:20:21 main timer tick: 1656076822965
2022/06/24 22:20:21 two second
2022/06/24 22:20:21 six seconds
2022/06/24 22:20:22 stop timer main
two: 15
four: 7
six: 5
eight: 3

```
