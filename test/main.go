package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var (
	// Compile regular expressions once to improve performance
	dateRe  = regexp.MustCompile(`\d+`)
	timeRe  = regexp.MustCompile(`(\d{2}:\d{2})-(\d{2}:\d{2})`)
)

func getSchedule(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	var data []string
	doc.Find(".data-grey, .data-white").Each(func(index int, row *goquery.Selection) {
		dateText := row.Find("td").Eq(2).Text()
		timeText := row.Find("td").Eq(3).Text()

		dateMatch := dateRe.FindString(dateText)
		if len(dateMatch) == 1 {
			dateMatch = "0" + dateMatch
		}

		timeMatches := timeRe.FindStringSubmatch(timeText)
		if timeMatches != nil {
			// Use time.Parse to directly work with time.Duration if applicable
			startTime, errStart := time.Parse("15:04", timeMatches[1])
			endTime, errEnd := time.Parse("15:04", timeMatches[2])
			if errStart == nil && errEnd == nil {
				duration := endTime.Sub(startTime)
				hours := duration.Hours()
				formattedData := fmt.Sprintf("%s,%.2f", dateMatch, hours)
				data = append(data, formattedData)
			}
		}
	})
	return data, nil
}

func main() {
	url := "https://schema.mau.se/setup/jsp/Schema.jsp?slutDatum=2024-03-31&sprak=SV&sokMedAND=true&startDatum=2024-02-29&moment=philip&resurser=k.BIT%20-%20IT"
	start := time.Now()

	schedule, err := getSchedule(url)
	if err != nil {
		log.Fatalf("Failed to get schedule: %v", err)
	}

	fmt.Println(schedule)
	fmt.Printf("Execution time: %v\n", time.Since(start))
}
