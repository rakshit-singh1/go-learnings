package main

import (
	"fmt"
	"time"
)

func main() {

	currentTime := time.Now()
	fmt.Println(currentTime)                                          // 2026-01-01 12:00:21.876728916 +0530 IST m=+0.000025564
	fmt.Println(currentTime.Format("2006-01-02 15:04:05"))            // 2026-01-01 12:01:32
	fmt.Println(currentTime.Format("2006-01-02 12:04:05 PM"))         // 2026-01-01 11:01:32 PM
	fmt.Println(currentTime.Format("2006-01-02 12:04:05 AM"))         // 2026-01-01 11:01:32 AM
	fmt.Println(currentTime.Format("2006-01-02 13:04:05 AM"))         // 2026-01-01 112:01:32 AM
	fmt.Println(currentTime.Format("2006-01-02 03:04:05 PM"))         // 12:04:16 PM
	fmt.Println(currentTime.Format("Mon Jan 2 12:04:05 AM IST 2006")) // Thu Jan 1 11:23:50 AM IST 2026
	fmt.Println(currentTime.Format("Mon Jan 2 12:04 AM IST 2006"))    // Thu Jan 1 11:23 AM IST 2026
	fmt.Println(currentTime.Format("Mon Jan 2 03:04 PM MST 2006"))    // Thu Jan 1 12:26 PM IST 2026
	fmt.Printf("%T\n", currentTime)

	// currentTime := time.Now()
	// fmt.Println("Current time: ", currentTime)
	// fmt.Printf("Type of currentTime %T\n", currentTime)

	formatted := currentTime.Format("2006/01/02, Monday")
	fmt.Println("Formatted time: ", formatted)

	layout_str := "02/01/2006"
	dateStr := "25/11/2030"
	formatted_time, _ := time.Parse(layout_str, dateStr) // dateStr will in format of layout_str format value will be of dateStr
	fmt.Println("Formatted time 1:", formatted_time)     // Formatted time 1: 2030-11-25 00:00:00 +0000 UTC

	// add 2 more day in currentTime
	new_date := currentTime.Add(48 * time.Hour)
	fmt.Println("new_date time: ", new_date)
	formatted_new_date := new_date.Format("2006/01/02 Monday")
	fmt.Println("formatted_new_date time: ", formatted_new_date)

	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		panic(err)
	}

	t := time.Now().In(loc)
	fmt.Println(t.Format("2006-01-02 03:04:05 PM MST"))

	loc2, err := time.LoadLocation("UTC")
	if err != nil {
		panic(err)
	}

	t2 := time.Now().In(loc2)
	fmt.Println(t2.Format("2006-01-02 03:04:05 PM MST"))

	loc3, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		panic(err)
	}

	t3 := time.Now().In(loc3)
	fmt.Println(t3.Format("2006-01-02 03:04:05 PM MST"))

	new_date1 := currentTime.Add(2 * time.Hour)
	fmt.Println(new_date1)

	// Addbles
	// 	time.Nanosecond
	// time.Microsecond
	// time.Millisecond
	// time.Second
	// time.Minute
	// time.Hour

	//  Formats
	// 	"2006-01-02"        // 2026-01-01
	// "02-01-2006"        // 01-01-2026
	// "02 Jan 2006"       // 01 Jan 2026
	// "Monday, Jan 2 2006"
	// "15:04:05"          // 24-hour
	// "03:04:05 PM"       // 12-hour
	// "3:04 PM"
	// "2006-01-02 15:04:05"
	// "2006-01-02 03:04:05 PM"
	// "02 Jan 2006 15:04"
	// "Mon Jan 2 03:04 PM"
	// "2006-01-02 15:04:05 MST"
	// "2006-01-02 15:04:05 -0700"
	// "2006-01-02T15:04:05Z07:00"

}
