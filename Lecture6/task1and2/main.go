package main

import "fmt"

type MagicList struct {
	LastItem *Item
}
type Item struct {
	Value    int
	PrevItem *Item
}

func add(list *MagicList, value int) {
	i := Item{
		Value: value,
	}

	if list.LastItem == nil {
		list.LastItem = &i
	} else {
		i.PrevItem = list.LastItem
		list.LastItem = &i
	}

}

func toSlice(list *MagicList) []int {
	item := list.LastItem
	var itemSlice []int
	for true {
		if item != nil {
			itemSlice = append(itemSlice, item.Value)
			item = item.PrevItem
		} else {
			break
		}
	}
	return itemSlice
}

func main() {
	magicList := &MagicList{}

	add(magicList, 5)
	add(magicList, 6)
	add(magicList, 7)
	add(magicList, 8)
	add(magicList, 9)
	fmt.Println(toSlice(magicList))

}
