package main

import (
	"fmt"
	"time"
)

func PrintTime(label string, t *time.Time) {
	fmt.Println(label, t.Format(time.RFC822Z))
}

func writeToChannel(nameChannel chan<- string) {

	names := []string{"Alice", "Bob", "Charlie", "Dora"}

	ticker := time.NewTicker(time.Second / 10)
	index := 0

	for {
		<-ticker.C
		nameChannel <- names[index]
		index++
		if index == len(names) {
			ticker.Stop()
			close(nameChannel)
			break
		}
	}
}

func main() {
	Label()
	//nameChannel := make(chan string)
	//
	//go writeToChannel(nameChannel)
	//
	//for name := range nameChannel {
	//	Printfln("Read name: %v", name)
	//}
}

func parseDUration() {
	d, err := time.ParseDuration("1h30m")
	if err == nil {
		Printfln("Hours: %v", d.Hours())
		Printfln("Mins: %v", d.Minutes())
		Printfln("Seconds: %v", d.Seconds())
		Printfln("Millseconds: %v", d.Milliseconds())
	} else {
		fmt.Println(err.Error())
	}
}

func parseLocations() {
	layout := "02 Jan 06 15:04"
	date := "09 Jun 95 19:30"

	//london, lonerr := time.LoadLocation("Europe/London")
	//newyork, nycerr := time.LoadLocation("America/New_York")
	//local, _ := time.LoadLocation("Local")

	london := time.FixedZone("BST", 1*60*60)
	newyork := time.FixedZone("EDT", -4*60*60)
	local := time.FixedZone("Local", 0)

	//if lonerr == nil && nycerr == nil {
	nolocation, _ := time.Parse(layout, date)
	londonTime, _ := time.ParseInLocation(layout, date, london)
	newyorkTime, _ := time.ParseInLocation(layout, date, newyork)
	localTime, _ := time.ParseInLocation(layout, date, local)
	PrintTime("No location:", &nolocation)
	PrintTime("London:", &londonTime)
	PrintTime("New York:", &newyorkTime)
	PrintTime("Local:", &localTime)

	//} else {
	//	fmt.Println(lonerr.Error(), nycerr.Error())
	//}
}

func parseStringToTime() {
	layout := "2006-Jan-02"
	dates := []string{
		"1995-Jun-09",
		"2015-Jun-02",
	}
	for _, d := range dates {
		tim, err := time.Parse(layout, d)
		if err == nil {
			PrintTime("Parsed", &tim)
		} else {
			Printfln("Error: %s", err.Error())
		}
	}
}

func Label() {
	current := time.Now()
	specific := time.Date(1995, time.June, 15, 0, 0, 0, 0, time.Local)
	unix := time.Unix(1433228090, 0)

	fmt.Println("current", current.Format(time.RFC850))
	//PrintTimeLabel("Current", &current)
	PrintTimeLabel("Specific", &specific)
	PrintTimeLabel("UNIX", &unix)

}

func PrintTimeLabel(label string, t *time.Time) {
	//layout := "Day: 01 Month: Jan Year: 2006"
	fmt.Println(label, t.Format(time.RFC822))
}
