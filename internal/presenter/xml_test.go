package presenter

import (
	"testing"

	"github.com/fabriciojsil/convert-csv-data/internal/model"
)

func TestXMLPresenter(t *testing.T) {

	t.Run("Parse Hotels to XML", func(t *testing.T) {
		expected := `<Hotels><hotel><name>The Gibson</name><address>63847 Lowe Knoll, East Maxine, WA 97030-4876</address><stars>5</stars><contact>Dr. Sinda Wyman</contact><phone>1-270-665-9933x1626</phone><url>http://www.paucek.com/search.htm</url></hotel><hotel><name>Martini Cattaneo</name><address>Stretto Bernardi 004, Quarto Mietta nell&#39;emilia, 07958 Torino (OG)</address><stars>0</stars><contact>Rosalino Marchetti</contact><phone>+39 627 68225719</phone><url>http://www.farina.org/blog/categories/tags/about.html</url></hotel></Hotels>`
		hotel, hotel2 := createHotels()
		hotels := model.Hotels{Hotels: []model.Hotel{hotel, hotel2}}
		xmlPresenter := XML{
			Hotels: hotels,
		}
		parsed := xmlPresenter.Parser()
		if expected != string(parsed) {
			t.Errorf("The Struct is diferent, expected %v Actual %v", expected, string(parsed))
		}
	})

	t.Run("Parse empty Hotels to XML", func(t *testing.T) {
		expected := `<Hotels></Hotels>`
		hotels := model.Hotels{Hotels: []model.Hotel{}}
		xmlPresenter := XML{
			Hotels: hotels,
		}
		parsed := xmlPresenter.Parser()
		if expected != string(parsed) {
			t.Errorf("The Struct is diferent, expected %v Actual %v", expected, string(parsed))
		}
	})

	t.Run("Parse nil Hotels to XML", func(t *testing.T) {
		expected := `<Hotels></Hotels>`
		xmlPresenter := XML{}
		parsed := xmlPresenter.Parser()
		if expected != string(parsed) {
			t.Errorf("The Struct is diferent, expected %v Actual %v", expected, string(parsed))
		}
	})
}
