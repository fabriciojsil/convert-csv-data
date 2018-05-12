package presenter

import (
	"encoding/xml"
	"io/ioutil"

	"github.com/fabriciojsil/convert-csv-data/internal/model"
)

type XML struct {
	Hotels model.Hotels
	data   []byte
}

func (x *XML) Parser() {
	x.data, _ = xml.Marshal(x.Hotels)
}
func (x *XML) Present(pathToSave string) error {
	return ioutil.WriteFile(pathToSave, x.data, 0644)
}
