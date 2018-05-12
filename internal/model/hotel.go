package model

import (
	"net/url"
)

type Hotel struct {
	Name    string
	Address string
	Stars   int
	Contact string
	Phone   string
	URL     string
}

func fixRating(stars int) (validStars int) {
	validStars = stars
	if stars < 0 {
		return 0
	}
	if stars > 5 {
		return 5
	}
	return validStars
}

func validateURI(uri string) (string, error) {
	_, err := url.ParseRequestURI(uri)
	if err != nil {
		return "", err
	}
	return uri, nil
}

// NewHotel create a Valid Hotel Struct if the validation is OK
func HotelFactory(name, address, contact, phone, uri string, stars int) (Hotel, error) {
	if _, err := validateURI(uri); err != nil {
		return Hotel{}, err
	}

	return Hotel{
		Name:    name,
		Address: address,
		Stars:   fixRating(stars),
		Contact: contact,
		Phone:   phone,
		URL:     uri,
	}, nil
}
