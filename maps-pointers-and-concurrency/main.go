package main

import (
	"fmt"
	"time"
)

type Item struct {
	name   string
	things map[string]string
}

func NewItem() *Item {
	return &Item{
		name:   "",
		things: make(map[string]string),
	}
}

func NewItemNoPointer() Item {
	return Item{
		name:   "",
		things: make(map[string]string),
	}
}

func main() {
	{
		fmt.Println("Pointer")
		item := NewItem()

		var channel chan *Item = make(chan *Item)

		done := make(chan struct{})

		go change(channel, done)

		fmt.Println("before:", item)
		channel <- item
		time.Sleep(1 * time.Second)
		fmt.Println("after:", item)

		done <- struct{}{}
	}

	fmt.Println("")
	fmt.Println("----------------------------------------")
	fmt.Println("")

	{
		fmt.Println("No Pointer: ")

		item := NewItemNoPointer()

		done := make(chan struct{})

		var channel chan Item = make(chan Item)

		go changeNoPointer(channel, done)

		fmt.Println("before:", item)
		channel <- item

		// for {
		// 	item.things["name"] = "Not Christian"
		// 	channel <- item
		// }

		fmt.Println("after:", item)

		done <- struct{}{}
	}

}

func change(in <-chan *Item, done <-chan struct{}) {
	for {
		select {
		case item := <-in:
			item.things["name"] = "Christian"
			item.name = "Christian"
		case <-done:
			return
		}
	}
}

func changeNoPointer(in <-chan Item, done <-chan struct{}) {
	for {
		select {
		case item := <-in:
			item.name = "Christian"
			item.things["name"] = "Christian"
		case <-done:
			return
		}
	}
}
