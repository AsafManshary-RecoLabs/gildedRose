package gildedrose

type Updater interface {
	Update()
}

type BackstagePassItem struct {
	Item
}

func NewBackstagePassItem(item *Item) Updater {
	return &BackstagePassItem{*item}
}
func (t *BackstagePassItem) Update() {
	if t.SellIn < 6 {
		t.updateQuality(3)
	} else if t.SellIn < 11 {
		t.updateQuality(2)
	} else {
		t.updateQuality(1)
	}
	t.SellIn--
	if t.SellIn < 0 {
		t.Quality = 0
	}
}

type BrieItem struct {
	Item
}

func NewBrieItem(item *Item) Updater {
	return &BrieItem{*item}
}

func (t *BrieItem) Update() {
	t.SellIn--
	t.updateQuality(2)
}

type OtherItem struct {
	Item
}

func NewOtherItem(item *Item) Updater {
	return &OtherItem{*item}
}

func (t *OtherItem) Update() {
	if t.Quality < 50 {
		t.Quality++
	}
	t.SellIn--
	if t.SellIn < 0 {

	}
}

type SulfurasItem struct {
	Item
}

func NewSulfItem(item *Item) Updater {
	return &SulfurasItem{*item}
}

func (t *SulfurasItem) Update() {
}
