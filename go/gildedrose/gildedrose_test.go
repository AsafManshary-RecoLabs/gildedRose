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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateQuality(tt.args.items)
		})
	}
}
