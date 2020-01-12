package carpenter

import (
	"log"
	"strings"
)

// Nest breaks an array of seasons into an array of outer times innter columns
func Nest(a []string, outer int, inner int) [][]string {
	log.Printf("Running the nest() function")
	teams := [][]string{}
	i := 0
	j := 0
	k := 0
	for i < outer {
		temp := []string{}
		for j < inner {
			// An empty string is appended if an index is out of range - an
			// error is thrown otherwise
			if k >= (outer*inner)-1 {
				temp = append(temp, "")
			} else {
				// strings.TrimSpace() is used to eliminate the return character
				// at the end of each line
				temp = append(temp, strings.TrimSpace(a[k]))
			}
			j++
			k++
		} // inner
		teams = append(teams, temp)
		temp = []string{}
		j = 0
		i++
	} // outer
	return teams
} // nest

// CastMapAsString iterates over a map and returns its key-value pairs in a 
// single string
func CastMapAsString(m map[string]string) string {
	var s []string
	for a, b := range m {
		s = append(s, a, b)
	}
	returnString := strings.Join(s, " ")
	return returnString
}