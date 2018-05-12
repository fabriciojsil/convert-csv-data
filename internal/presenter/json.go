package presenter

import (
	"encoding/json"

	"github.com/fabriciojsil/convert-csv-data/internal/model"
)

type JSON struct {
	Hotels model.Hotels
}

func (j JSON) Parser() []byte {
	data, err := json.Marshal(j.Hotels)
	if err != nil {
		return []byte{}
	}
	return data
}
