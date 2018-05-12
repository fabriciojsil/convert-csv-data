package presenter

import (
	"encoding/xml"

	"github.com/fabriciojsil/convert-csv-data/internal/model"
)

type XML struct {
	Hotels model.Hotels
}

func (x XML) Parser() []byte {
	data, err := xml.Marshal(x.Hotels)
	if err != nil {
		return []byte{}
	}
	return data
}
