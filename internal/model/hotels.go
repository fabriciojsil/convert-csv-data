package model

import (
	"reflect"
	"sort"
)

type Hotels struct {
	Hotels []Hotel `xml:"hotel" json:"hotels"`
}

//InsertHotel Just insert a hotel in the slice
func (h *Hotels) InsertHotel(hotel Hotel) {
	h.Hotels = append(h.Hotels, hotel)
}

//InsertSortedBy insert a hotel ASC/DESC ordered by field
func (h *Hotels) InsertSortedBy(hotel Hotel, field, order string) {
	if field == "Stars" {
		h.insertIntSortedBy(hotel, field, order)
	} else {
		h.insertStringSortedBy(hotel, field, order)
	}
}

func (h *Hotels) insertIntSortedBy(hotel Hotel, field, order string) {
	index := sort.Search(len(h.Hotels), h.sortByInt(hotel, field, order))
	h.appendToHotels(hotel, index)
}

func (h *Hotels) insertStringSortedBy(hotel Hotel, field, order string) {
	index := sort.Search(len(h.Hotels), h.sortByString(hotel, field, order))
	h.appendToHotels(hotel, index)
}

func (h *Hotels) sortByString(hotel Hotel, field, order string) func(i int) bool {
	return func(i int) bool {
		a := getStringValueProperty(h.Hotels[i], field)
		b := getStringValueProperty(hotel, field)
		if order == "desc" {
			return a < b
		}
		return a > b
	}
}

func (h *Hotels) sortByInt(hotel Hotel, field, order string) func(i int) bool {
	return func(i int) bool {
		a := getIntValueProperty(h.Hotels[i], field)
		b := getIntValueProperty(hotel, field)
		if order == "desc" {
			return a < b
		}
		return a > b
	}
}

func (h *Hotels) appendToHotels(hotel Hotel, index int) {
	h.Hotels = append(h.Hotels, Hotel{})
	copy(h.Hotels[index+1:], h.Hotels[index:])
	h.Hotels[index] = hotel
}

func getStringValueProperty(hotel Hotel, field string) string {
	value := getReflectValue(hotel, field)
	return value.String()
}

func getIntValueProperty(hotel Hotel, field string) int64 {
	value := getReflectValue(hotel, field)
	return value.Int()
}

func getReflectValue(hotel Hotel, field string) reflect.Value {
	r := reflect.ValueOf(hotel)
	return reflect.Indirect(r).FieldByName(field)
}
