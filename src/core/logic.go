package core

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func aggregateEvents(events []*Event, existingReports []*Report) []*Report {
	reports := make(map[int]map[string]map[string]int)
	for _, v := range existingReports {
		if reports[v.UserId] == nil {
			reports[v.UserId] = make(map[string]map[string]int)
		}
		if reports[v.UserId][v.Date] == nil {
			reports[v.UserId][v.Date] = make(map[string]int)
		}
		reports[v.UserId][v.Date] = v.Counts
	}

	for _, event := range events {
		date := time.Unix(event.Timestamp, 0).UTC().Format("2006-01-02")

		if reports[event.UserId] == nil {
			reports[event.UserId] = make(map[string]map[string]int)
		}
		if reports[event.UserId][date] == nil {
			reports[event.UserId][date] = make(map[string]int)
		}

		reports[event.UserId][date][event.EventType]++
	}

	result := []*Report{}
	for userID, userReports := range reports {
		for date, counts := range userReports {
			report := Report{
				UserId: userID,
				Date:   date,
				Counts: counts,
			}
			result = append(result, &report)
		}
	}
	return result
}

func Runner() {
	inputFile := flag.String("i", "", "Input JSON file")
	outputFile := flag.String("o", "", "Output JSON file")
	update := flag.Bool("update", false, "Update existing reports with new events")
	flag.Parse()

	if *inputFile == "" || *outputFile == "" {
		fmt.Println("error : input/out files args are not passed. Please check the README.md for help")
		os.Exit(1)
	}

	events, err := readEvents(*inputFile)
	if err != nil {
		fmt.Println("error reading input file:", err)
		os.Exit(1)
	}

	reports := []*Report{}
	if *update {
		existingReports, err := readReports(*outputFile)
		if err != nil {
			fmt.Println("error reading existing reports:", err)
			os.Exit(1)
		}
		reports = append(reports, existingReports...)
	}
	newReports := aggregateEvents(events, reports)

	err = writeReports(*outputFile, newReports)
	if err != nil {
		fmt.Println("error in writting files:", err)
		os.Exit(1)
	}
	fmt.Println("reports aggregated successfully.")
}
