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

type CSVReader struct {
	headers map[string]int
}

// Convert will convert to Hotel Model
func (c *CSVReader) Convert(file *os.File) (hotels []model.Hotel) {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		if c.headers == nil {
			c.getHeaders(scanner.Text())
			continue
		}
		hotel, err := c.parseLineToHotel(scanner.Text())
		if err == nil {
			hotels = append(hotels, hotel)
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
