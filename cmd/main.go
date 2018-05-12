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

	savedOn, err := service.DeliveryFile(string(*format), string(*file))
	if err != nil {
		fmt.Println("Opss:", err.Error())
		fmt.Println("Finished in ", time.Since(start))
	} else {
		fmt.Println("Saved on ", savedOn)
		fmt.Println("Finished in ", time.Since(start))
	}

}
