package commas

import (
	"bufio"
	"log"
	"os"
	"github.com/amunnelly/soccer/dbconnect"
	"strconv"
	"strings"
)



// CompileCsv takes a slice of Table objects and creates a slice of
// strings, where each string is make up from the individual Table
// fields, seperated by commas. `WriteToFile()` is then called with this slice
// of strings as its paramater.
func CompileCsv(season []connectdb.PointsGdTable) {
	var holder = []string{}
	headings := "Team,Points,GD,GF,GA\n"
	// hbytes := []byte(headings)
	holder = append(holder, headings)

	for s := range season {
		t := season[s].Team
		p := season[s].Points
		gd := season[s].GD
		gf := season[s].GoalsFor
		ga := season[s].GoalsAgainst

		p1 := strconv.Itoa(int(p))
		gd1 := strconv.Itoa(int(gd))
		gf1 := strconv.Itoa(int(gf))
		ga1 := strconv.Itoa(int(ga))

		line := []string{t, p1, gd1, gf1, ga1}
		lineString := strings.Join(line, ",")
		lineString = strings.TrimSuffix(lineString, ",")
		lineString = lineString + "\n"

		holder = append(holder, lineString)

	}

	WriteToFile(holder)

}

// WriteToFile takes an array of strings and writes them to the location parameter.
func WriteToFile(data []string) {
	f, err := os.Create("./static/data/data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	f.Sync()

	w := bufio.NewWriter(f)
	for _, line := range data {
		_, err := w.WriteString(line)
		if err != nil {
			log.Fatal(err)
		}
	}
	w.Flush()
}
