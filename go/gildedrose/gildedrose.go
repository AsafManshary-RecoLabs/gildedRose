package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {
		UpdateQualityPerItem(items[i])
	}
}
func UpdateQualityPerItem(item *Item) {
	getUpdater(item).Update()
}
