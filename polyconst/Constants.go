package polyconst

import (
	"time"
	_ "time/tzdata"
)

var NYCTime *time.Location

func init() {
	nyc, err := time.LoadLocation("America/New_York")
	if err != nil {
		panic(err)
	}
	NYCTime = nyc
}
