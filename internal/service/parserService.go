package service

import (
	"os"
	"strings"

	"github.com/fabriciojsil/convert-csv-data/internal/model"
	"github.com/fabriciojsil/convert-csv-data/internal/parser"
	"github.com/fabriciojsil/convert-csv-data/internal/presenter"
)

func DeliveryFile(format, filePath string) (string, error) {

	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}

	hotelsParser := parser.CSVReader{}
	hotels := hotelsParser.Convert(file)

	p, definedFormat := factoryPresenter(format, hotels)
	pathToSave := strings.Replace(filePath, "csv", definedFormat, 1)
	p.Parser()
	p.Present(pathToSave)

	if err != nil {
		return "", err
	}
	return strings.Replace(pathToSave, "csv", format, 1), nil
}

func factoryPresenter(format string, hotels model.Hotels) (presenter.Presenter, string) {
	var (
		p             presenter.Presenter
		definedFormat string
	)

	switch format {
	case "json":
		p = &presenter.JSON{Hotels: hotels}
		definedFormat = "json"
	case "xml":
		p = &presenter.XML{Hotels: hotels}
		definedFormat = "xml"
	default:
		p = &presenter.JSON{Hotels: hotels}
		definedFormat = "json"
	}
	return p, definedFormat
}
