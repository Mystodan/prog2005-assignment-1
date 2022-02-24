package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

/**	checkError logs an error.
 *	@param inn - error value
 */
func checkError(inn error) {
	if inn != nil {
		log.Fatal(inn)
	}
}

func getBorderingIso(inn []Universities) string {
	var lastIso, allIso string
	var allIsoArr []string
	var country []getCountry
	i := 0
	for _, s := range inn {
		if s.Isocode != lastIso {
			lastIso = s.Isocode
			i++
		}
		if i > 1 {
			break
		}
	}

	write, err := http.Get("https://restcountries.com/v3.1/alpha?codes=" + lastIso)
	checkError(err)
	body, err := io.ReadAll(write.Body)
	checkError(err)
	json.Unmarshal(body, &country)

	for _, s := range country {
		allIso = strings.Join(append(allIsoArr, s.Borders...), ",")
	}

	return allIso
}
func getBorderingNames(inn string) []string {
	var country []getCountry
	var names []string
	write, err := http.Get("https://restcountries.com/v3.1/alpha?codes=" + inn)
	checkError(err)
	body, err := io.ReadAll(write.Body)
	checkError(err)
	json.Unmarshal(body, &country)

	for _, s := range country {
		names = append(names, s.Name.Name)
	}
	return names
}

func getBorderingUniversities(target []Universities, limit int) []Universities {
	var AllBorderingUniversities []Universities
	var BorderUnii []getUnii
	var BorderIso []getCountry
	BorderingNat := getBorderingNames(getBorderingIso(target))
	writeIso, err := http.Get("https://restcountries.com/v3.1/alpha?codes=" + getBorderingIso(target))
	checkError(err)
	bI, err := io.ReadAll(writeIso.Body)
	checkError(err)
	json.Unmarshal(bI, &BorderIso)

	var isoCode, limCounter int
	var lastIso string
	var tempBU []getUnii
	for _, s := range BorderingNat {
		writeName, err := http.Get("http://universities.hipolabs.com/search?name" + "&country=" + s)
		checkError(err)
		bN, err := io.ReadAll(writeName.Body)
		checkError(err)
		json.Unmarshal(bN, &tempBU)
		BorderUnii = append(BorderUnii, tempBU...)
	}

	for i := range BorderUnii {
		for j, _ := range BorderIso {
			if BorderIso[j].Isocode == BorderUnii[i].Alpha_2 {
				isoCode = j
			}
		}

		if lastIso != BorderUnii[i].Alpha_2 {
			lastIso = BorderUnii[i].Alpha_2
			limCounter = 0
		} else {
			limCounter++
		}

		if limCounter < limit {
			AllBorderingUniversities = append(AllBorderingUniversities, Universities{
				Name:      BorderUnii[i].Name,
				Country:   BorderUnii[i].Country,
				Isocode:   BorderUnii[i].Alpha_2,
				Webpages:  BorderUnii[i].Webpages,
				Languages: BorderIso[isoCode].Languages,
				Map:       BorderIso[isoCode].Maps.Map,
			})
		}
		if limit == 0 {
			AllBorderingUniversities = append(AllBorderingUniversities, Universities{
				Name:      BorderUnii[i].Name,
				Country:   BorderUnii[i].Country,
				Isocode:   BorderUnii[i].Alpha_2,
				Webpages:  BorderUnii[i].Webpages,
				Languages: BorderIso[isoCode].Languages,
				Map:       BorderIso[isoCode].Maps.Map,
			})

		}
	}
	return AllBorderingUniversities
}

/* func setBorderingUniversity(inn []getUnii) []Universities {
	var lastCountry string
	var universities []Universities
	var country []getCountry
	for _, s := range inn {
		var body []byte
		if s.Alpha_2 != lastCountry {
			lastCountry = s.Alpha_2
			write, err := http.Get("https://restcountries.com/v3.1/alpha?codes=" + lastCountry)
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
} */

func setUniversity(inn []getUnii) []Universities {
	var lastCountry string
	var universities []Universities
	var country []getCountry
	for _, s := range inn {
		var body []byte
		if s.Alpha_2 != lastCountry {
			lastCountry = s.Alpha_2
			write, err := http.Get("https://restcountries.com/v3.1/alpha?codes=" + lastCountry)
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
