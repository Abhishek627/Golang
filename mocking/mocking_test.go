package main

import (
	"log"
	"reflect"
	"testing"
	"time"
)

// Creating a fake/static time variable to override value of now
var fakenow = func() time.Time {
	t := "2006-01-02T15:04:05Z"
	curr, err := time.Parse(TimeLayout, t)
	if err != nil {
		log.Println("Error in parsing time ", err)
	}
	return curr
}

func Test_dateTime_updateTime(t *testing.T) {
	now = fakenow
	log.Println("now is ", now())
	tests := []struct {
		name string
		dt   *dateTime
		want *dateTime
	}{
		{
			name: "success case",
			dt:   &dateTime{},
			want: &dateTime{"2006-01-02 15:04:05 +0000 UTC"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dt.updateTime(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dateTime.updateTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
