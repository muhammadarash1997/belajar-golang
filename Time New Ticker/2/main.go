package main

import (
	"fmt"
	"time"
)

func main() {
	timeOne()
	timeTwo()
}

func timeOne() {
	myTime := time.Date(2022, time.August, 06, 03, 07, 45, 10, time.UTC).Unix()
	fmt.Println(myTime)

	// Unix() return number of seconds elapsed since January 1, 1970 UTC
	nowTime := time.Now().Local().Unix()
	fmt.Println(nowTime)

	// Because in above we use Unix then it should be multiplied by time.Second
	duration := time.Duration((myTime - nowTime) * int64(time.Second))
	fmt.Println(duration)

	ticker := time.NewTicker(duration)
	fmt.Println(<-ticker.C) // Will be blocked until duration is reached
}

func timeTwo() {
	myTime := time.Date(2022, time.August, 06, 03, 05, 45, 10, time.UTC).UnixMilli()
	fmt.Println(myTime)

	// UnixMilli() return number of seconds elapsed since January 1, 1970 UTC
	nowTime := time.Now().UTC().UnixMilli()
	fmt.Println(nowTime)

	// Because in above we use Unix then it should be multiplied by time.Millisecond
	duration := time.Duration((myTime - nowTime) * int64(time.Millisecond))
	fmt.Println(duration)

	ticker := time.NewTicker(duration)
	fmt.Println(<-ticker.C) // Will be blocked until duration is reached
}

// NOTE
// Unix() return number of seconds elapsed since January 1, 1970 UTC, so the return should be multiplied by time.Second if you want to use with time.Duration
// UnixMilli() return number of milliseconds elapsed since January 1, 1970 UTC, so the return should be multiplied by time.Millisecond if you want to use with time.Duration
// UnixNano() return number of nanoseconds elapsed since January 1, 1970 UTC, so the return should be multiplied by time.Nanosecond if you want to use with time.Duration
