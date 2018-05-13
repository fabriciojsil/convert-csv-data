package parser

import (
	"bufio"
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/fabriciojsil/convert-csv-data/internal/model"
)

const ASC = "asc"
const DESC = "desc"

type CSVReader struct {
	headers map[string]int
}

func validateOrdering(field, order string) (string, string, bool) {
	if len(field) == 0 || (order != ASC && order != DESC) {
		return "", "", false
	}
	return field, order, true
}

// Convert will convert to Hotel Model
func (c *CSVReader) Run(file *os.File, field, order string) (hotels model.Hotels) {
	field, order, toOrder := validateOrdering(field, order)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		if c.headers == nil {
			c.getHeaders(scanner.Text())
			continue
		}
		hotel, err := c.parseLineToHotel(scanner.Text())
		if err == nil {
			if toOrder {
				hotels.InsertSortedBy(hotel, field, order)
			} else {
				hotels.InsertHotel(hotel)
			}
		}
	}
	return hotels
}

func (c *CSVReader) parseLineToHotel(line string) (model.Hotel, error) {
	return c.createHotel(extractCSV(line))
}

func extractCSV(line string) []string {
	r := csv.NewReader(strings.NewReader(line))
	r.TrimLeadingSpace = true
	r.Comma = ','
	r.Comment = '#'
	record, err := r.Read()

	if err != nil {
		log.Fatal(err)
	}

	return record
}

func (c *CSVReader) getHeaders(header string) {
	headers := make(map[string]int)
	record := extractCSV(header)
	for rec := range record {
		headers[record[rec]] = rec
	}
	c.headers = headers
}

func (c *CSVReader) createHotel(record []string) (model.Hotel, error) {
	name := record[c.headers["name"]]
	address := record[c.headers["address"]]
	starsRaw := record[c.headers["stars"]]
	contact := record[c.headers["contact"]]
	phone := record[c.headers["phone"]]
	uri := record[c.headers["uri"]]

	stars, err := strconv.Atoi(starsRaw)
	if err != nil {
		stars = 0
	}
	return model.HotelFactory(name, address, contact, phone, uri, stars)
}
