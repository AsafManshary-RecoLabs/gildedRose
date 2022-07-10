package gildedrose

type Updater interface {
	// Update - updates the item according to its new state
	Update()
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

func (t *BrieItem) Update() {
	t.SellIn--
	t.updateQuality(2)
}

func (t *OtherItem) Update() {
	if t.Quality < 50 {
		t.Quality++
	}
	t.SellIn--
}

func (t *SulfurasItem) Update() {
}
