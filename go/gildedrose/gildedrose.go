package gildedrose

// UpdateQuality - updates list of items
func UpdateQuality(items []*Item) {
	for _, item := range items {
		updateQualityPerItem(item)
	}
}
func updateQualityPerItem(item *Item) {
	getUpdater(item).Update()
}
