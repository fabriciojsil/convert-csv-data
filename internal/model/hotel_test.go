package model

import (
	"reflect"
	"testing"
)

func TestHotelModel(t *testing.T) {
	hotel := Hotel{
		Name:    "The Gibson",
		Stars:   5,
		Address: "63847 Lowe Knoll, East Maxine, WA 97030-4876",
		Contact: "Dr. Sinda Wyman",
		Phone:   "1-270-665-9933x1626",
		URL:     "http://www.paucek.com",
	}

	t.Run("creating a new Hotel with valid URL", func(t *testing.T) {
		result, _ := HotelFactory(
			"The Gibson",
			"63847 Lowe Knoll, East Maxine, WA 97030-4876",
			"Dr. Sinda Wyman",
			"1-270-665-9933x1626",
			"http://www.paucek.com",
			5)
		if !reflect.DeepEqual(result, hotel) {
			t.Errorf("The Struct is diferent, expected %v Actual %v", hotel, result)
		}
	})

	t.Run("creating a new Hotel with invalid URL without Scheme", func(t *testing.T) {

		_, err := HotelFactory(
			"The Gibson",
			"63847 Lowe Knoll, East Maxine, WA 97030-4876",
			"Dr. Sinda Wyman",
			"1-270-665-9933x1626",
			"www.paucek.com",
			5)

		if err == nil {
			t.Errorf("it Should return an error %v", err)
		}
	})

	t.Run("creating a new Hotel with invalid URL without Scheme", func(t *testing.T) {

		_, err := HotelFactory(
			"The Gibson",
			"63847 Lowe Knoll, East Maxine, WA 97030-4876",
			"Dr. Sinda Wyman",
			"1-270-665-9933x1626",
			"paucek",
			5)

		if err == nil {
			t.Errorf("it Should return an error %v", err)
		}
	})

	t.Run("creating a new Hotel with valid star", func(t *testing.T) {
		result, _ := HotelFactory(
			"The Gibson",
			"63847 Lowe Knoll, East Maxine, WA 97030-4876",
			"Dr. Sinda Wyman",
			"1-270-665-9933x1626",
			"http://www.paucek.com",
			5)
		if !reflect.DeepEqual(result, hotel) {
			t.Errorf("The Struct is diferent, expected %v Actual %v", hotel, result)
		}
	})

	t.Run("creating a new Hotel with negative star", func(t *testing.T) {
		hotel.Stars = 0
		result, _ := HotelFactory(
			"The Gibson",
			"63847 Lowe Knoll, East Maxine, WA 97030-4876",
			"Dr. Sinda Wyman",
			"1-270-665-9933x1626",
			"http://www.paucek.com",
			-5)

		if !reflect.DeepEqual(result, hotel) {
			t.Errorf("The Struct is diferent, expected %v Actual %v", hotel, result)
		}
	})

	t.Run("creating a new Hotel with big then 5 star", func(t *testing.T) {
		hotel.Stars = 5
		result, _ := HotelFactory(
			"The Gibson",
			"63847 Lowe Knoll, East Maxine, WA 97030-4876",
			"Dr. Sinda Wyman",
			"1-270-665-9933x1626",
			"http://www.paucek.com",
			50)
		if !reflect.DeepEqual(result, hotel) {
			t.Errorf("The Struct is diferent, expected %v Actual %v", hotel, result)
		}
	})

}
