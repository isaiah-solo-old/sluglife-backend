// dining provides the ability to get the current menus from the dining hall.
package dining

import (
	"fmt"
	"../scraper"
	"log"
)

const (
	website_format = "http://nutrition.sa.ucsc.edu/menuSamp.asp?locationNum=%s"
)

var (
	dhalls = []diningId{
		{"Cowell & Stevenson", "05"},
		{"Crown & Merrill", "20"},
		{"Porter & Kresge", "25"},
		{"College Eight & Oakes", "30"},
		{"College Nine & College Ten", "40"},
	}
)

type diningId struct {
	name  string
	locId string
}

type Location struct {
	Name string `json:"name"`
	Menu Menu   `json:"items"`
}

func handleUrlError(err error, url string) {
	if err != nil {
		log.Printf("Unable to create scraper for url: %s: %s", url, err.Error())
	}
}

func ParseAll() []Location {
	menus := make([]Location, len(dhalls))
	for i, v := range dhalls {
		url := fmt.Sprintf(website_format, v.locId)
		menu, err := scraper.NewFromURL(url)
		handleUrlError(err, url)
		menus[i] = Location{v.name, menuDoc{menu}.Parse()}
	}
	return menus
}
