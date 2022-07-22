package item_database

import (
	"eviecoin/items"
	"github.com/fatih/color"
)

var rawItems = map[string]items.LocalItemData{}

func loadItems(items_ ...items.LocalItemData) {
	for _, item := range items_ {
		rawItems[item.Id] = item
		color.Green("[items] Loaded item: %s", item.Label)
	}
}

func NewManager() ItemDatabase {
	_init()
	return ItemDatabase{
		GetItem:    GetItem,
		GetRawItem: GetRawItem,
		GetItems:   GetItems,
	}
}

type ItemDatabase struct {
	GetItem    func(string) (items.Item, error)
	GetRawItem func(string) (items.RawItemData, error)
	GetItems   func() []items.Item
}
