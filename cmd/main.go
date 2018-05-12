package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/fabriciojsil/convert-csv-data/internal/service"
)

func main() {
	start := time.Now()
	file := flag.String("file", ".", "a string")
	format := flag.String("format", "json", "a string")
	flag.Parse()

	service.DeliveryFile(string(*format), string(*file))

	fmt.Println(time.Since(start))
}
