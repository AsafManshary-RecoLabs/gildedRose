package gildedrose

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {
		UpdateQualityPerItem(items[i])
	}
}
func UpdateQualityPerItem(item *Item) {
	getUpdater(item).Update()
}
