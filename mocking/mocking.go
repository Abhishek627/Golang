package main

import (
	"fmt"
	"time"
)

var now = time.Now

const (
	//TimeLayout is standard RFC time format to be used here.
	TimeLayout = "2006-01-02T15:04:05Z"
)

//Normally using time.Now() will be harder to unit test
//We can create a var here and overrride it in tests
type dateTime struct {
	currTime string
}

// updateTime updates the struct time to current UTC time
func (dt *dateTime) updateTime() *dateTime {
	//Doing something on the
	currTime := now().UTC().String()
	dt.currTime = currTime
	return dt
}

func main() {
	dt := &dateTime{}
	dt.updateTime()
	fmt.Println("Current time is ", dt.currTime)

}
