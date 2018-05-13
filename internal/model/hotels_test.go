package model

import (
	"reflect"
	"testing"
)

func TestHotels(t *testing.T) {
	t.Run("Adding ASC ordered By Name", func(t *testing.T) {
		h0, h1, h2 := CreateHotel()
		hotels := Hotels{}
		hotels.InsertSortedBy(h2, "Name", "asc")
		hotels.InsertSortedBy(h0, "Name", "asc")
		hotels.InsertSortedBy(h1, "Name", "asc")

		if !reflect.DeepEqual(hotels.Hotels[0], h0) {
			t.Errorf("Expected %v Actual %v", h0, hotels.Hotels[0])
		}
		if !reflect.DeepEqual(hotels.Hotels[1], h1) {
			t.Errorf("Expected %v Actual %v", h1, hotels.Hotels[1])
		}
		if !reflect.DeepEqual(hotels.Hotels[2], h2) {
			t.Errorf("Expected %v Actual %v", h2, hotels.Hotels[2])
		}
	})

	t.Run("Adding DESC ordered By Name", func(t *testing.T) {
		h0, h1, h2 := CreateHotel()
		hotels := Hotels{}
		hotels.InsertSortedBy(h2, "Name", "desc")
		hotels.InsertSortedBy(h0, "Name", "desc")
		hotels.InsertSortedBy(h1, "Name", "desc")

		if !reflect.DeepEqual(hotels.Hotels[0], h2) {
			t.Errorf("Expected %v Actual %v", h0, hotels.Hotels[2])
		}
		if !reflect.DeepEqual(hotels.Hotels[1], h1) {
			t.Errorf("Expected %v Actual %v", h1, hotels.Hotels[1])
		}
		if !reflect.DeepEqual(hotels.Hotels[2], h0) {
			t.Errorf("Expected %v Actual %v", h2, hotels.Hotels[0])
		}
	})

	t.Run("Adding DESC ordered Address", func(t *testing.T) {
		h0, h1, h2 := CreateHotel()
		hotels := Hotels{}
		hotels.InsertSortedBy(h2, "Address", "desc")
		hotels.InsertSortedBy(h0, "Address", "desc")
		hotels.InsertSortedBy(h1, "Address", "desc")

		if !reflect.DeepEqual(hotels.Hotels[0], h2) {
			t.Errorf("Expected %v Actual %v", h0, hotels.Hotels[2])
		}
		if !reflect.DeepEqual(hotels.Hotels[1], h1) {
			t.Errorf("Expected %v Actual %v", h1, hotels.Hotels[1])
		}
		if !reflect.DeepEqual(hotels.Hotels[2], h0) {
			t.Errorf("Expected %v Actual %v", h2, hotels.Hotels[0])
		}
	})

	t.Run("Adding DESC ordered Stars ASC", func(t *testing.T) {
		h0, h1, h2 := CreateHotel()
		hotels := Hotels{}
		hotels.InsertSortedBy(h2, "Stars", "asc")
		hotels.InsertSortedBy(h0, "Stars", "asc")
		hotels.InsertSortedBy(h1, "Stars", "asc")

		if !reflect.DeepEqual(hotels.Hotels[0], h0) {
			t.Errorf("Expected %v Actual %v", h0, hotels.Hotels[0])
		}
		if !reflect.DeepEqual(hotels.Hotels[1], h1) {
			t.Errorf("Expected %v Actual %v", h1, hotels.Hotels[1])
		}
		if !reflect.DeepEqual(hotels.Hotels[2], h2) {
			t.Errorf("Expected %v Actual %v", h2, hotels.Hotels[2])
		}
	})

	t.Run("Adding DESC ordered Stars DESC", func(t *testing.T) {
		h0, h1, h2 := CreateHotel()
		hotels := Hotels{}
		hotels.InsertSortedBy(h2, "Stars", "desc")
		hotels.InsertSortedBy(h0, "Stars", "desc")
		hotels.InsertSortedBy(h1, "Stars", "desc")

		if !reflect.DeepEqual(hotels.Hotels[0], h2) {
			t.Errorf("Expected %v Actual %v", h2, hotels.Hotels[2])
		}
		if !reflect.DeepEqual(hotels.Hotels[1], h1) {
			t.Errorf("Expected %v Actual %v", h1, hotels.Hotels[1])
		}
		if !reflect.DeepEqual(hotels.Hotels[2], h0) {
			t.Errorf("Expected %v Actual %v", h0, hotels.Hotels[0])
		}
	})

	t.Run("Adding Hotel", func(t *testing.T) {
		h0, h1, h2 := CreateHotel()
		hotels := Hotels{}
		hotels.InsertHotel(h2)
		hotels.InsertHotel(h0)
		hotels.InsertHotel(h1)

		if !reflect.DeepEqual(hotels.Hotels[0], h2) {
			t.Errorf("Expected %v Actual %v", h2, hotels.Hotels[2])
		}
		if !reflect.DeepEqual(hotels.Hotels[1], h0) {
			t.Errorf("Expected %v Actual %v", h1, hotels.Hotels[1])
		}
		if !reflect.DeepEqual(hotels.Hotels[2], h1) {
			t.Errorf("Expected %v Actual %v", h0, hotels.Hotels[0])
		}
	})

}

func CreateHotel() (Hotel, Hotel, Hotel) {
	hotelPosition1 := Hotel{
		Name:    "a",
		Stars:   1,
		Address: "d",
		Contact: "Dr. Sinda Wyman",
		Phone:   "1-270-665-9933x1626",
		URL:     "http://www.paucek.com",
	}
	hotelPosition2 := Hotel{
		Name:    "b",
		Stars:   2,
		Address: "e",
		Contact: "",
		Phone:   "1-270-665-9933x1626",
		URL:     "http://www.paucek.com",
	}
	hotelPosition3 := Hotel{
		Name:    "c",
		Stars:   3,
		Address: "f",
		Contact: "",
		Phone:   "1-270-665-9933x1626",
		URL:     "http://www.paucek.com",
	}
	return hotelPosition1, hotelPosition2, hotelPosition3
}
