package parser

import (
	"os"
	"reflect"
	"testing"

	"github.com/fabriciojsil/convert-csv-data/internal/model"
)

func TestCSV(t *testing.T) {
	headers, hotel, hotel2 := createExpectation()
	t.Run("Testing parseLine ", func(t *testing.T) {
		expected := hotel
		csv := `The Gibson,"63847 Lowe Knoll, East Maxine, WA 97030-4876",5,Dr. Sinda Wyman,1-270-665-9933x1626, http://www.paucek.com/search.htm`
		reader := CSVReader{}
		reader.headers = headers
		result, _ := reader.parseLineToHotel(csv)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("The Struct is diferent, expected %v Actual %v", expected, result)
		}
	})

	t.Run("Testing GetHeaders", func(t *testing.T) {
		expected := headers

		csv := `name,address,stars,contact,phone,uri`
		reader := CSVReader{}
		reader.getHeaders(csv)

		if !reflect.DeepEqual(reader.headers, expected) {
			t.Errorf("The Struct is diferent, expected %v  Actual %v", expected, reader.headers)
		}
	})

	t.Run("Testing Corvert", func(t *testing.T) {

		expected := []model.Hotel{hotel, hotel2}

		file, _ := os.Open("../../data/data.csv")

		reader := CSVReader{}
		hotels := reader.Convert(file)

		if !reflect.DeepEqual(hotels, expected) {
			t.Errorf("The Struct is diferent, expected %v  Actual %v", expected, hotels)
		}
	})
}

func createExpectation() (map[string]int, model.Hotel, model.Hotel) {
	headers := map[string]int{
		"name":    0,
		"address": 1,
		"stars":   2,
		"contact": 3,
		"phone":   4,
		"uri":     5,
	}
	hotel := model.Hotel{
		Name:    "The Gibson",
		Stars:   5,
		Address: "63847 Lowe Knoll, East Maxine, WA 97030-4876",
		Contact: "Dr. Sinda Wyman",
		Phone:   "1-270-665-9933x1626",
		URL:     "http://www.paucek.com/search.htm",
	}
	hotel2 := model.Hotel{
		Name:    "Martini Cattaneo",
		Stars:   0,
		Address: "Stretto Bernardi 004, Quarto Mietta nell'emilia, 07958 Torino (OG)",
		Contact: "Rosalino Marchetti",
		Phone:   "+39 627 68225719",
		URL:     "http://www.farina.org/blog/categories/tags/about.html",
	}
	return headers, hotel, hotel2

}
