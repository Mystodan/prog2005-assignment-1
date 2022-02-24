package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

var timer time.Time

func TimerStart() {
	timer = time.Now()
}

func getUptime(inn time.Time) time.Duration {
	return time.Since(inn)
}

/**	checkError logs an error.
 *	@param inn - error value
 */
func checkError(inn error) {
	if inn != nil {
		log.Fatal(inn)
	}
}

/**	Get issues a GET to the specified URL.
 *	@param inn - URL
 */
func getURL(inn string) *http.Response {
	ret, err := http.Get(inn)
	checkError(err)
	return ret
}

/**	Gets bordering countries iso codes to the targeted university.
 *  and returns their iso codes in one string
 *	@param inn - targeted university
 *	@return isocodes f.ex."NO,RU,DE"
 */
func getBorderingIsos(inn []Universities) string {
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

	write := getURL(GET_CNTR + CNRT_REQ + lastIso)
	body, err := io.ReadAll(write.Body)
	checkError(err)
	json.Unmarshal(body, &country)

	for _, s := range country {
		allIso = strings.Join(append(allIsoArr, s.Borders...), ",")
	}

	return allIso
}

/**	takes in alpha/iso codes and gets their info and returns it.
 *	@param inn - iso codes f.ex."NO,RU,DE"
 */
func getBorderingNames(inn string) ([]string, []getCountry) {
	var country []getCountry
	var names []string
	write := getURL(GET_CNTR + CNRT_REQ + inn)
	body, err := io.ReadAll(write.Body)
	checkError(err)
	json.Unmarshal(body, &country)

	for _, s := range country {
		names = append(names, s.Name.Name)
	}
	return names, country
}

/**	takes in a university and returns a list of universities from neighbouring countries.
 *  @see getBorderingNames(inn string)
 *  @see getBorderingIsos(inn []Universities)
 *	@param inn - targeted university
 *	@param limit - the limit for amount of universities to be printed per country.
 *  @return AllBorderingUniversities - all bordering universities
 */
func getBorderingUniversities(target []Universities, limit int) []Universities {
	var AllBorderingUniversities []Universities
	var BorderUnii []getUnii
	BorderingNat, BorderIsos := getBorderingNames(getBorderingIsos(target))

	var isoCode, limCounter int
	var lastA2 string
	for _, s := range BorderingNat {
		var tempBU []getUnii
		writeName := getURL(GET_UNI + UNI_REQ + "&country=" + s)
		bN, err := io.ReadAll(writeName.Body)
		checkError(err)
		json.Unmarshal(bN, &tempBU)
		BorderUnii = append(BorderUnii, tempBU...)
	}

	for i := range BorderUnii {
		if lastA2 != BorderUnii[i].Alpha_2 {
			lastA2 = BorderUnii[i].Alpha_2
			limCounter = 0
		} else {
			limCounter++
		}
		for j := range BorderIsos {
			if BorderIsos[j].Isocode == BorderUnii[i].Alpha_2 {
				isoCode = j
			}
		}
		if limCounter < limit {
			AllBorderingUniversities = append(AllBorderingUniversities, Universities{
				Name:      BorderUnii[i].Name,
				Country:   BorderUnii[i].Country,
				Isocode:   BorderUnii[i].Alpha_2,
				Webpages:  BorderUnii[i].Webpages,
				Languages: BorderIsos[isoCode].Languages,
				Map:       BorderIsos[isoCode].Maps.Map,
			})
		}
		if limit == 0 {
			AllBorderingUniversities = append(AllBorderingUniversities, Universities{
				Name:      BorderUnii[i].Name,
				Country:   BorderUnii[i].Country,
				Isocode:   BorderUnii[i].Alpha_2,
				Webpages:  BorderUnii[i].Webpages,
				Languages: BorderIsos[isoCode].Languages,
				Map:       BorderIsos[isoCode].Maps.Map,
			})

		}
	}
	return AllBorderingUniversities
}

/**	takes incomplete data for universities, and returns a list of universities.
 *	Compiled data from GET_UNI and GET_CNTR APIs
 *	@param inn - data from GET_UNI
 *  @return universities - a list of universities compiled from two APIs
 */
func setUniversity(inn []getUnii) []Universities {
	var lastCountry string
	var universities []Universities
	var country []getCountry
	var isoCode int
	for _, s := range inn {
		var tempCoun []getCountry
		if s.Alpha_2 != lastCountry {
			lastCountry = s.Alpha_2
			write := getURL(GET_CNTR + CNRT_REQ + lastCountry)
			ret, err := io.ReadAll(write.Body)
			checkError(err)
			json.Unmarshal(ret, &tempCoun)
			country = append(country, tempCoun...)
		}
		for j := range country {
			if country[j].Isocode == s.Alpha_2 {
				isoCode = j
			}
		}
		// Create instance of content (could be read from DB, file, etc.)
		universities = append(universities, Universities{
			Name:      s.Name,
			Country:   s.Country,
			Isocode:   s.Alpha_2,
			Webpages:  s.Webpages,
			Languages: country[isoCode].Languages,
			Map:       country[isoCode].Maps.Map,
		})
	}
	return universities
}
