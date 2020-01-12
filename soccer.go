package main

import (
	"fmt"
	"github.com/amunnelly/soccer/routing"
	"log"
	"net/http"
	"os"
)

func main() {
	// TO-DO: Check about swapping prefixes
	http.Handle("/static/", http.FileServer(http.Dir(".")))

	http.HandleFunc("/", routing.ServeHome)
	http.HandleFunc("/chooseTeams", routing.ServeTeams)
	http.HandleFunc("/chooseSeasonGraph", routing.ServeSeasonGraph)
	http.HandleFunc("/chooseSeasonTable", routing.ServeSeasonTable)
	http.HandleFunc("/about", routing.ServeAbout)
	http.HandleFunc("/contact", routing.ServeContact)
	http.HandleFunc("/logContact", routing.RecordContact)

	// TO-DO: Fix these rascals
	http.HandleFunc("/teams/{team}", routing.Team)
	http.HandleFunc("/table", routing.SeasonTable)
	http.HandleFunc("/graph", routing.SeasonGraph)

	fmt.Println("Incipio - I begin.")
	fmt.Println(os.Getenv("DATABASE_URL"))

	if len(os.Getenv("PORT")) > 0 {
		log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
	} else {
		log.Fatal(http.ListenAndServe(":8080", nil))
	}
}
