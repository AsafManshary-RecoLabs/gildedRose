package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

func (t *Item) updateQuality(quality int) {
	if t.Quality+quality >= 50 {
		t.Quality = 50
	} else if t.Quality < 50 {
		t.Quality += quality
	}
}

type BackstagePassItem struct {
	Item
}

func NewBackstagePassItem(item *Item) Updater {
	return &BackstagePassItem{*item}
}

type BrieItem struct {
	Item
}

func NewBrieItem(item *Item) Updater {
	return &BrieItem{*item}
}

type OtherItem struct {
	Item
}

func NewOtherItem(item *Item) Updater {
	return &OtherItem{*item}
}

type SulfurasItem struct {
	Item
}

func NewSulfurasItem(item *Item) Updater {
	return &SulfurasItem{*item}
}
