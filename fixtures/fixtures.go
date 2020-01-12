package fixtures

import (
	"strconv"
	"strings"
)

// Fixture holds match data in a JSON-type container
type Fixture struct {
	Date     string
	HomeTeam string
	AwayTeam string
	FTHG     int32
	FTAG     int32
	FTR      int32
	HTHG     int32
	HTAG     int32
	HTR      int32
	Referee  string
}

// ProcessFixtures iterates over each line of fixtures, leaves out extraneous
// data and returns the essential details of each game as a struct
func ProcessFixtures(fixtures []string) []*Fixture {
	var cleanFixtures []*Fixture

	for _, f := range fixtures {
		t := strings.Split(f, ",")
		details := Fixture{
			Date:     t[1],
			HomeTeam: t[2],
			AwayTeam: t[3],
			FTHG:     strconv.Atoi(t[4]),
			FTAG:     strconv.Atoi(t[5]),
			FTR:      strconv.Atoi(t[6]),
			HTHG:     strconv.Atoi(t[7]),
			HTAG:     strconv.Atoi(t[8]),
			HTR:      strconv.Atoi(t[9]),
			Referee:  t[10],
		}
		cleanFixtures.append(details)

	}
	return cleanFixtures
}
