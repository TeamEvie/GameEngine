package item_database

import "eviecoin/items"

func GetItem(s string) (items.Item, error) {
	item, ok := rawItems[s]

	if !ok {
		return items.Item{}, nil
	}

	return parseLocalItemToApiItem(item), nil
}

func GetItems() []items.Item {
	var items_ []items.Item

	for _, i := range rawItems {
		items_ = append(items_, parseLocalItemToApiItem(i))
	}

	return items_
}

func GetRawItem(s string) (items.RawItemData, error) {
	_, err := GetItem(s)

	if err != nil {
		return items.RawItemData{}, err
	}

	return items.RawItemData{
		Id: s,
	}, nil
}

func parseLocalItemToApiItem(i items.LocalItemData) items.Item {
	return items.Item{
		Raw:         i,
		Name:        i.Label,
		Description: i.Description,
		Price:       i.Price,
		Rarity:      parseRarity(i.Rarity),
		Emoji:       i.Emoji,
	}
}

func parseRarity(r items.Rarity) string {
	switch r {
	case items.Common:
		return "Common"
	case items.Uncommon:
		return "Uncommon"
	case items.Rare:
		return "Rare"
	case items.Epic:
		return "Epic"
	case items.Legendary:
		return "Legendary"
	}

	return "Unknown"
}
