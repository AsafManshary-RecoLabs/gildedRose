package gildedrose

const (
	brie     = "Aged Brie"
	sulfuras = "Sulfuras, Hand of Ragnaros"
	tafkal   = "Backstage passes to a TAFKAL80ETC concert"
)

func getUpdater(item *Item) (updater Updater) {
	if item.Name == tafkal {
		return NewTakfkalItem(item)
	}
	if item.Name == brie {
		return NewBrieItem(item)
	}
	if item.Name == sulfuras {
		return NewSulfItem(item)
	}
	return NewOtherItem(item)
}
