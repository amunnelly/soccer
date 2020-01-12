package movingfinger

import (
	"log"
	"os"
)

// RecordPageHit takes the page variable, which will be a page reference,
// and writes it to the log file
func RecordPageHit(page string) {
	f, err := os.OpenFile("text.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	logger := log.New(f, "", log.LstdFlags)
	logger.Printf(page)
}
