package gildedrose

type Updater interface {
	Update()
}

type TafkalItem struct {
	Item
}

func NewTakfkalItem(item *Item) Updater {
	return &TafkalItem{*item}
}
func (t *TafkalItem) Update() {
	if t.Quality < 50 {
		t.Quality = t.Quality + 1
		if t.SellIn < 11 {
			if t.Quality < 50 {
				t.Quality = t.Quality + 1
			}
		}
	}
	if t.SellIn < 6 {
		if t.Quality < 50 {
			t.Quality = t.Quality + 1
		}
	}
	t.SellIn = t.SellIn - 1
	if t.Quality < 50 {
		t.Quality = t.Quality + 1
	}
	if t.SellIn < 0 {
		t.Quality = t.Quality - t.Quality
	}
}

type BrieItem struct {
	Item
}

func NewBrieItem(item *Item) Updater {
	return &BrieItem{*item}
}

func (t *BrieItem) Update() {
	if t.Quality < 50 {
		t.Quality = t.Quality + 1
	}
	t.SellIn = t.SellIn - 1
	if t.Quality < 50 {
		t.Quality = t.Quality + 1
	}
	if t.SellIn < 0 {

	}
}

type OtherItem struct {
	Item
}

func NewOtherItem(item *Item) Updater {
	return &OtherItem{*item}
}

func (t *OtherItem) Update() {
	if t.Quality < 50 {
		t.Quality = t.Quality + 1
	}
	t.SellIn = t.SellIn - 1
	if t.Quality < 50 {
		t.Quality = t.Quality + 1
	}
	if t.SellIn < 0 {

	}
}

type SulfItem struct {
	Item
}

func NewSulfItem(item *Item) Updater {
	return &SulfItem{*item}
}

func (t *SulfItem) Update() {
}
