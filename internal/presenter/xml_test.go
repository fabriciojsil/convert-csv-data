package presenter

import (
	"testing"

	"github.com/fabriciojsil/convert-csv-data/internal/model"
)

func TestXMLPresenter(t *testing.T) {

	t.Run("Parse", func(t *testing.T) {
		t.Run("Parse Hotels to XML", func(t *testing.T) {
			expected := `<Hotels><hotel><name>The Gibson</name><address>63847 Lowe Knoll, East Maxine, WA 97030-4876</address><stars>5</stars><contact>Dr. Sinda Wyman</contact><phone>1-270-665-9933x1626</phone><url>http://www.paucek.com/search.htm</url></hotel><hotel><name>Martini Cattaneo</name><address>Stretto Bernardi 004, Quarto Mietta nell&#39;emilia, 07958 Torino (OG)</address><stars>0</stars><contact>Rosalino Marchetti</contact><phone>+39 627 68225719</phone><url>http://www.farina.org/blog/categories/tags/about.html</url></hotel></Hotels>`
			hotel, hotel2 := createHotels()
			hotels := model.Hotels{Hotels: []model.Hotel{hotel, hotel2}}
			xmlPresenter := XML{
				Hotels: hotels,
			}
			xmlPresenter.Parser()
			if expected != string(xmlPresenter.data) {
				t.Errorf("The Struct is diferent, expected %v Actual %v", expected, string(xmlPresenter.data))
			}
		})

		t.Run("Parse empty Hotels to XML", func(t *testing.T) {
			expected := `<Hotels></Hotels>`
			hotels := model.Hotels{Hotels: []model.Hotel{}}
			xmlPresenter := XML{
				Hotels: hotels,
			}
			xmlPresenter.Parser()
			if expected != string(xmlPresenter.data) {
				t.Errorf("The Struct is diferent, expected %v Actual %v", expected, string(xmlPresenter.data))
			}
		})

		t.Run("Parse nil Hotels to XML", func(t *testing.T) {
			expected := `<Hotels></Hotels>`
			xmlPresenter := XML{}
			xmlPresenter.Parser()
			if expected != string(xmlPresenter.data) {
				t.Errorf("The Struct is diferent, expected %v Actual %v", expected, string(xmlPresenter.data))
			}
		})
	})

	t.Run("Present", func(t *testing.T) {
		tearDown := setupTest(t)
		defer tearDown(t)

		t.Run("Write file with xml", func(t *testing.T) {
			hotel, hotel2 := createHotels()
			jsonPresenter := &XML{
				Hotels: model.Hotels{Hotels: []model.Hotel{hotel, hotel2}},
			}
			jsonPresenter.Parser()
			err := jsonPresenter.Present("../../data/data.xml")

			if err != nil {
				t.Errorf("Mustn't return err %v", err)
			}
		})
	})

}
