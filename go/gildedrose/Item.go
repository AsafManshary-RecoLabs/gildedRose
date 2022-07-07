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
