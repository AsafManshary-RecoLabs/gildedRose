package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

func UpdateQuality(items []*Item) (err error) {
	for i := 0; i < len(items); i++ {
		err = UpdateQualityPerItem(items[i])
		if err == nil {
			return err
		}
	}
	return nil
}
func UpdateOther(item *Item) (err error) {
	if item.Quality < 50 {
		item.Quality = item.Quality + 1
	}
	item.SellIn = item.SellIn - 1
	if item.Quality < 50 {
		item.Quality = item.Quality + 1
	}
	if item.SellIn < 0 {

	}
	return err
}
func UpdateSulf(item *Item) (err error) {
	return err
}

func UpdateBrie(item *Item) (err error) {
	if item.Quality < 50 {
		item.Quality = item.Quality + 1
	}
	item.SellIn = item.SellIn - 1
	if item.Quality < 50 {
		item.Quality = item.Quality + 1
	}
	if item.SellIn < 0 {

	}
	return err
}
func UpdateTafkal(item *Item) (err error) {
	if item.Quality < 50 {
		item.Quality = item.Quality + 1
		if item.SellIn < 11 {
			if item.Quality < 50 {
				item.Quality = item.Quality + 1
			}
		}
		if item.SellIn < 6 {
			if item.Quality < 50 {
				item.Quality = item.Quality + 1
			}
		}
	}
	item.SellIn = item.SellIn - 1
	if item.Quality < 50 {
		item.Quality = item.Quality + 1
	}
	if item.SellIn < 0 {
		item.Quality = item.Quality - item.Quality
	}
	return err
}
func UpdateQualityPerItem(item *Item) (err error) {
	const Tafkal = "Backstage passes to a TAFKAL80ETC concert"
	const Sulfuras = "Sulfuras, Hand of Ragnaros"
	const Brie = "Aged Brie"
	if item.Name == Tafkal {
		return UpdateTafkal(item)
	}
	if item.Name == Brie {
		return UpdateBrie(item)
	}
	if item.Name == Sulfuras {
		return UpdateSulf(item)
	}
	return UpdateOther(item)
}
