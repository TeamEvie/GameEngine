package items

type Rarity int

const (
	Common Rarity = iota
	Uncommon
	Rare
	Epic
	Legendary
)

type RawItemData struct {
	Id string `json:"id"`
}

type LocalItemData struct {
	Id          string
	Label       string
	Description string
	Price       int
	Rarity      Rarity
	Emoji       string
}

type Item struct {
	Raw         LocalItemData
	Name        string
	Description string
	Price       int
	Rarity      string
	Emoji       string
}
