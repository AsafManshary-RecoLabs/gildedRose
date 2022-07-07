package gildedrose

import "testing"

func TestUpdateQuality(t *testing.T) {
	type args struct {
		items []*Item
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "handle empty items",
			args: args{items: []*Item{}},
		},
		{
			name: "handle Brie",
			args: args{
				items: []*Item{
					{Quality: 1, Name: "Brie", SellIn: 1}, {Quality: 2, Name: "Brie", SellIn: 0},
				},
			},
		},
		{
			name: "handle Tafkal",
			args: args{
				items: []*Item{
					{Quality: 1, Name: "Tafkal", SellIn: 1}, {Quality: 2, Name: "Tafkal", SellIn: 0},
				},
			},
		},
		{
			name: "handle Sulf",
			args: args{
				items: []*Item{
					{Quality: 1, Name: "Sulf", SellIn: 1}, {Quality: 2, Name: "Sulf", SellIn: 0},
				},
			},
		},
		{
			name: "handle Other",
			args: args{
				items: []*Item{
					{Quality: 1, Name: "Other", SellIn: 1}, {Quality: 2, Name: "Other", SellIn: 0},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateQuality(tt.args.items)
		})
	}
}

func TestUpdateQualityPerItem(t *testing.T) {
	type args struct {
		item *Item
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateQualityPerItem(tt.args.item)
		})
	}
}
