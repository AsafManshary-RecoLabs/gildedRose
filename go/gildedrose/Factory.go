package gildedrose

const (
	brie          = "Aged Brie"
	sulfuras      = "Sulfuras, Hand of Ragnaros"
	backstagePass = "Backstage passes to a TAFKAL80ETC concert"
)

func getUpdater(item *Item) (updater Updater) {
	if item.Name == backstagePass {
		return NewBackstagePassItem(item)
	}
	if item.Name == brie {
		return NewBrieItem(item)
	}
	if item.Name == sulfuras {
		return NewSulfItem(item)
	}
	return NewOtherItem(item)
}
