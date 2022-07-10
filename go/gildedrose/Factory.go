package gildedrose

const (
	brie          = "Aged Brie"
	sulfuras      = "Sulfuras, Hand of Ragnaros"
	backstagePass = "Backstage passes to a TAFKAL80ETC concert"
)

// getUpdater - find a suitable updater for an item
func getUpdater(item *Item) (updater Updater) {
	switch updater = NewOtherItem(item); item.Name {
	case backstagePass:
		NewBackstagePassItem(item)
	case brie:
		NewBrieItem(item)
	case sulfuras:
		NewSulfurasItem(item)
	default:
		NewOtherItem(item)
	}
	return updater
}
