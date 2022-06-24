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
