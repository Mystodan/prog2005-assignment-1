package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

/**	checkError logs an error.
 *	@param inn - error value
 */
func checkError(inn error) {
	if inn != nil {
		log.Fatal(inn)
	}
}

func setUniversity(inn []getUnii) []Universities {
	var lastCountry string
	var universities []Universities
	var country []getCountry
	for _, s := range inn {
		var body []byte
		if s.Country != lastCountry {
			lastCountry = s.Country
			write, err := http.Get("https://restcountries.com/v3.1/name/" + lastCountry)
			checkError(err)
			ret, err := io.ReadAll(write.Body)
			checkError(err)
			body = ret
		}
		json.Unmarshal(body, &country)

		// Create instance of content (could be read from DB, file, etc.)
		universities = append(universities, Universities{
			Name:      s.Name,
			Country:   s.Country,
			Isocode:   s.Alpha_2,
			Webpages:  s.Webpages,
			Languages: country[0].Languages,
			Map:       country[0].Maps.Map,
		})

	}
	return universities
}
