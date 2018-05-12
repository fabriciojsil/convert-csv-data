package presenter

import (
	"os"
	"testing"

	"github.com/fabriciojsil/convert-csv-data/internal/model"
)

func TestJSONPresenter(t *testing.T) {
	t.Run("Parse", func(t *testing.T) {
		t.Run("Parse Hotels to JSON", func(t *testing.T) {
			expected := `{"hotels":[{"name":"The Gibson","address":"63847 Lowe Knoll, East Maxine, WA 97030-4876","stars":5,"contact":"Dr. Sinda Wyman","phone":"1-270-665-9933x1626","url":"http://www.paucek.com/search.htm"},{"name":"Martini Cattaneo","address":"Stretto Bernardi 004, Quarto Mietta nell'emilia, 07958 Torino (OG)","stars":0,"contact":"Rosalino Marchetti","phone":"+39 627 68225719","url":"http://www.farina.org/blog/categories/tags/about.html"}]}`
			hotel, hotel2 := createHotels()
			jsonPresenter := &JSON{
				Hotels: model.Hotels{Hotels: []model.Hotel{hotel, hotel2}},
			}
			jsonPresenter.Parser()
			if expected != string(jsonPresenter.data) {
				t.Errorf("The Struct is diferent, expected %v Actual %v", expected, string(jsonPresenter.data))
			}
		})

		t.Run("Parse empty Hotels to JSON", func(t *testing.T) {
			expected := `{"hotels":[]}`
			hotels := []model.Hotel{}
			jsonPresenter := &JSON{
				Hotels: model.Hotels{Hotels: hotels},
			}
			jsonPresenter.Parser()
			if expected != string(jsonPresenter.data) {
				t.Errorf("The Struct is diferent, expected %v Actual %v", expected, string(jsonPresenter.data))
			}
		})

		t.Run("Parse nil Hotels to JSON", func(t *testing.T) {
			expected := `{"hotels":null}`
			jsonPresenter := &JSON{}
			jsonPresenter.Parser()
			if expected != string(jsonPresenter.data) {
				t.Errorf("The Struct is diferent, expected %v Actual %v", expected, string(jsonPresenter.data))
			}
		})
	})

	t.Run("Present", func(t *testing.T) {
		tearDown := setupTest(t)
		defer tearDown(t)

		t.Run("Write file with json", func(t *testing.T) {
			hotel, hotel2 := createHotels()
			jsonPresenter := &JSON{
				Hotels: model.Hotels{Hotels: []model.Hotel{hotel, hotel2}},
			}
			jsonPresenter.Parser()
			err := jsonPresenter.Present("../../data/data.json")

			if err != nil {
				t.Errorf("Mustn't return err %v", err)
			}
		})
	})
}

func createHotels() (model.Hotel, model.Hotel) {
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
	return hotel, hotel2
}

func setupTest(t *testing.T) func(t *testing.T) {
	t.Log("ensure there is no parsed files")
	os.Remove("../../data/data.json")
	os.Remove("../../data/data.xml")

	return func(t *testing.T) {
		os.Remove("../../data/data.json")
		os.Remove("../../data/data.xml")
		t.Log("removing parsed files")
	}
}
