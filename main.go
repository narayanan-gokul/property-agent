package main

import (
	"fmt"
	"net/http"
	"time"
	"io"
	"strings"
	"strconv"
	"math"
	"os"
	"bufio"
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Suburbs []string
	Availability time.Time
	MaxDistance float64
	TempFolder string
}

type Listing struct {
	address string
	price float64
	longitude float64
	latitude float64
	availability time.Time
	link string
}

func (listing Listing) prettyPrint() {
	fmt.Printf(
		"%s\n%f\nAvailable: %s\nLocation: longitude: %f latitude: %f\n",
		listing.address,
		listing.price,
		listing.availability.Format(time.DateOnly),
		listing.longitude,
		listing.latitude,
	)
}

func (listing Listing) filePrintString() string {
	return fmt.Sprintf(
		"- [%s](%s)\n\t- Price: $%.0f\n\t- Available: %s\n\t- Location: %f, %f\n",
		listing.address,
		listing.link,
		listing.price,
		listing.availability.Format(time.DateOnly),
		listing.latitude,
		listing.longitude,
	)
}


const BUFFER_SIZE int = 4096

func substringSearch(input []byte, target []byte, startIndex int) int {
	result := -1
	lenInput := len(input)
	lenTarget := len(target)
	if lenInput - startIndex < lenTarget {
		return result
	}
	for i := startIndex; i < lenInput; i++ {
		result = i
		j := 0
		for i < lenInput && j < lenTarget && input[i] == target[j] {
			i++
			j++
		}
		if j == lenTarget {
			return result
		} else {
			j = 0
			result = -1
		}
	}
	return result
}

func extractLinkFromChunk(chunk []byte) string {
	lineHasCSSString := -1
	lineHashref := substringSearch(chunk, []byte("href"), 0)
	if lineHashref != -1 {
		lineHasCSSString = substringSearch(
			chunk,
			[]byte("css-1y2bib4"),
			lineHashref,
		)
	}
	if lineHasCSSString != -1 {
		startIndex := lineHashref + 6
		endIndex := substringSearch(
			chunk,
			[]byte("\""),
			startIndex,
		)
		return string(chunk[startIndex:endIndex])
	}
	return ""
}

func makeRequest(client *http.Client, URL string) *http.Response {
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		log.Println("Request creation error:", err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64)")
	req.Header.Set(
		"Accept",
		"text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
	)
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Request failed:", err)
	}
	log.Println("Request status:", resp.Status)
	return resp
}

func getListings(client *http.Client, suburbs []string, page int) (listings []string) {
	URL := fmt.Sprintf(
		"https://www.domain.com.au/rent/?%s&bedrooms=2-any&bathrooms=2-any&price=0-1000&availableto=2025-07-14&excludedeposittaken=1&page=%d",
		strings.Join(suburbs, ","),
		page,
	) 
	resp := makeRequest(client, URL)
	defer resp.Body.Close()
	for {
		buffer := make([]byte, BUFFER_SIZE)
		bytesRead, err := io.ReadAtLeast(
			resp.Body, buffer, BUFFER_SIZE)
		if err != nil && err != io.EOF && err != io.ErrUnexpectedEOF {
			log.Println("Error getting listings: ", err)
		}
		if bytesRead == 0 {
			break
		}
		link := extractLinkFromChunk(buffer[:bytesRead])
		if len(link) != 0 {
			listings = append(listings, link)
		}
	}
	return listings
}

func (listing Listing) distanceFrom(lat float64, lng float64) float64 {
	const PI float64 = 3.141592653589793
	
	radlat1 := float64(PI * listing.latitude / 180)
	radlat2 := float64(PI * lat / 180)
	
	theta := float64(listing.longitude - lng)
	radtheta := float64(PI * theta / 180)
	
	dist := math.Sin(radlat1) * math.Sin(radlat2) + math.Cos(radlat1) * math.Cos(radlat2) * math.Cos(radtheta)
	
	if dist > 1 {
		dist = 1
	}
	
	dist = math.Acos(dist)
	dist = dist * 180 / PI
	dist = dist * 60 * 1.1515
	
	return dist * 1.609344
}

func extractListing(client *http.Client, link string) *Listing {
	var listing Listing
	page := make([]byte, 0)
	resp := makeRequest(client, link)
	defer resp.Body.Close()
	for {
		buffer := make([]byte, BUFFER_SIZE)
		bytesRead, err := io.ReadAtLeast(resp.Body, buffer, BUFFER_SIZE)
		if err != nil && err != io.EOF && err != io.ErrUnexpectedEOF {
			log.Println("Error reading listing %s:", link, err)
		}
		if bytesRead == 0 {
			break
		}
		page = append(page, buffer...)
	}

	priceStart := substringSearch(
		page,
		[]byte("<div data-testid=\"listing-details__summary-title\" class=\"css-twgrok\">"),
		0,)
	priceStart = substringSearch(page, []byte("$"), priceStart) + 1
	priceEnd := substringSearch(page, []byte(" "), priceStart)
	priceString := string(page[priceStart:priceEnd])

	addressStart := substringSearch(
		page,
		[]byte("<h1 class=\"css-hkh81z\">"),
		priceEnd,
	) + 23
	addressEnd := substringSearch(page, []byte("</h1>"), addressStart)
	listing.address = string(page[addressStart:addressEnd])

	availabilityDateStart := substringSearch(
		page,
		[]byte("Available from<!"),
		addressEnd,
	)
	var availabilityDateEnd int
	if availabilityDateStart == -1 {
		listing.availability = time.Now()
		availabilityDateEnd = addressEnd
	} else {
		availabilityDateStart = substringSearch(page, []byte("<strong>"), availabilityDateStart) + 8
		availabilityDateEnd = substringSearch(page, []byte("</strong>"), availabilityDateStart)
		availabilityDateString := string(page[availabilityDateStart:availabilityDateEnd])
		temp := strings.Split(availabilityDateString, " ")
		
		day := temp[1][:len(temp[1]) - 2]
		month := temp[2][:3]
		year := temp[3][2:4]

		if len(day) == 1 {
			day = fmt.Sprintf("0%s", day)
		}

		availabilityDate, err := time.Parse(time.RFC822, fmt.Sprintf("%s %s %s 00:00 IST", day, month, year))
		
		if err != nil {
			log.Println("Error parsing date:", err)
			log.Println(link)
		}
		listing.availability = availabilityDate
	}

	locationStart := substringSearch(page, []byte("maps.googleapis.com/maps/api/staticmap?center="), availabilityDateEnd)
	locationStart = substringSearch(page, []byte("center="), locationStart) + 7
	locationEnd := substringSearch(page, []byte("\\u0026"), locationStart)
	locationString := string(page[locationStart:locationEnd])
	locationStringSplit := strings.Split(locationString, ",")
	longString := locationStringSplit[1]
	latString := locationStringSplit[0]

	long, err := strconv.ParseFloat(longString, 64)
	if err != nil {
		log.Println("Error parsing longitude:", err)
	}
	lat, err:= strconv.ParseFloat(latString, 64)
	if err != nil {
		log.Println("Error parsing latitude:", err)
	}
	listing.latitude = lat
	listing.longitude = long

	price, err := strconv.ParseFloat(priceString, 64)
	if err != nil {
		priceString = priceString[:substringSearch([]byte(priceString), []byte("<"), 0)]
	}
	price, err = strconv.ParseFloat(priceString, 64)
	if err != nil {
		log.Println("Error parsing price:", err)
		log.Println(link)
	}
	listing.price = price
	listing.link = link

	return &listing
}

func filterListings(client *http.Client, listings []string, availableDate time.Time, distance float64) []*Listing{
	filteredListings := make([]*Listing, 0)
	for _, listing := range listings {
		extractedListing := extractListing(client, listing)
		if !(extractedListing.distanceFrom(-33.888636, 151.187301) > distance) && availableDate.Before(extractedListing.availability) {
			filteredListings = append(filteredListings, extractedListing)
		}
	}
	return filteredListings
}

func testInclude(client *http.Client) {
	url := "https://www.domain.com.au/105-14-mcgill-street-lewisham-nsw-2049-12487121"
	listing := extractListing(client, url)
	log.Println(listing.filePrintString())
}

func createPropertiesFile(listings []*Listing) {
	fileName := fmt.Sprintf("properties-%s.md", time.Now().Format(time.DateTime))
	propertiesFile, err := os.Create(fmt.Sprintf("%s/%s", os.Args[1], fileName))
	defer propertiesFile.Close()
	if err != nil {
		log.Println("Error creating file properties.md:", err)
	}
	writer := bufio.NewWriter(propertiesFile)
	for _, listing := range listings {
		_, err = writer.WriteString(listing.filePrintString())
		if err != nil {
			log.Println("Error writing to properties file:", err)
		}
	}
	writer.Flush()
}

func main() {
	log.Println("Initializing client")
	tr := &http.Transport{
		ForceAttemptHTTP2: false,
	}

	client := &http.Client{
		Transport: tr,
		Timeout:   10 * time.Second,
	}

	log.Println("Fetching config")
	var config Config
	_, err := toml.DecodeFile("config.toml", &config)
	if err != nil {
		log.Println("Error fetching config file:", err)
	}

	log.Println("Fetching listings")
	var listings []string
	page := 1
	for {
		resultsPerPage := getListings(client, config.Suburbs, page)
		if len(resultsPerPage) == 0 {
			break
		}
		listings = append(listings, resultsPerPage...)
		page++
	}

	filteredListings := filterListings(client, listings, config.Availability, config.MaxDistance)
	createPropertiesFile(filteredListings)
}
