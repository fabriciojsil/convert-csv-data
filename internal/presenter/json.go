package presenter

import (
	"encoding/json"
	"io/ioutil"

	"github.com/fabriciojsil/convert-csv-data/internal/model"
)

type JSON struct {
	Hotels model.Hotels
	data   []byte
}

func (j *JSON) Parser() {
	j.data, _ = json.Marshal(j.Hotels)
}

func (j *JSON) Present(pathToSave string) error {
	return ioutil.WriteFile(pathToSave, j.data, 0644)
}
