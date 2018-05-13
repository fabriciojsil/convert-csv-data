package service

import (
	"os"
	"strings"

	"github.com/fabriciojsil/convert-csv-data/internal/model"
	"github.com/fabriciojsil/convert-csv-data/internal/parser"
	"github.com/fabriciojsil/convert-csv-data/internal/presenter"
)

const JSON = "json"
const XML = "xml"
const CSV = "csv"

func DeliveryFile(format, filePath, field, order string) (string, error) {

	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}

	hotelsParser := parser.CSVReader{}
	hotels := hotelsParser.Run(file, field, order)

	p, definedFormat := factoryPresenter(format, hotels)
	pathToSave := strings.Replace(filePath, CSV, definedFormat, 1)
	p.Parser()
	p.Present(pathToSave)

	if err != nil {
		return "", err
	}
	return strings.Replace(pathToSave, CSV, format, 1), nil
}

func factoryPresenter(format string, hotels model.Hotels) (presenter.Presenter, string) {
	var (
		p             presenter.Presenter
		definedFormat string
	)

	switch format {
	case JSON:
		p = &presenter.JSON{Hotels: hotels}
		definedFormat = JSON
	case XML:
		p = &presenter.XML{Hotels: hotels}
		definedFormat = XML
	default:
		p = &presenter.JSON{Hotels: hotels}
		definedFormat = JSON
	}
	return p, definedFormat
}
