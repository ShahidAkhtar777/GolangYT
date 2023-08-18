package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	now := time.Now()
	p(now)

	p(now.Year())
	p(now.Month())
	p(now.Day())
	p(now.Hour())
	p(now.Weekday())

	p(now.Format("01-02-2006 Monday"))
}
