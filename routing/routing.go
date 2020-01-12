package routing

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"github.com/amunnelly/soccer/movingfinger"
	"net/http"
	"github.com/amunnelly/soccer/carpenter"
	"github.com/amunnelly/soccer/commas"
	"github.com/amunnelly/soccer/dbconnect"
	"strings"
	"time"
	"os"

	// _ "github.com/gorilla/mux"
)

var t *template.Template
var i interface{}

// SeasonGraphDetail is a struct that supplies the title of a particular graph (the
// season itsself, invariably.)
type SeasonGraphDetail struct {
	Title string
	Results []connectdb.PointsGdTable
}

// TeamGraphDetail is a struct that supplies the title of a particular graph (the
// team itsself, invariably.)
type TeamGraphDetail struct {
	Title string
	Results []connectdb.Points
}

// ContactName records the first and last names of anyone who filled out and 
// submitted the contact form.
type ContactName struct {
	Firstname string
	Lastname string
}

// ServeHome returns the home page template
func ServeHome(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseGlob("./templates/*")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("\nHome.")
	movingfinger.RecordPageHit("Home.")
	t.ExecuteTemplate(w, "index.html", i)
}

// ServeTeams returns the teams template
func ServeTeams(w http.ResponseWriter, r *http.Request) {
	// Find the teams
	fileBytes, err := ioutil.ReadFile("./static/teams.txt")
	if err != nil {
		log.Fatal(err)
	}
	s := string(fileBytes)
	prototeams := strings.Split(s, "\n")

	teams := carpenter.Nest(prototeams, 8, 5)
	t, err := template.ParseGlob("./templates/*")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("\nTeams.")
	movingfinger.RecordPageHit("Teams.")

	t.ExecuteTemplate(w, "chooseTeams.html", teams)
}

// ServeSeasonTable returns the season table template
func ServeSeasonTable(w http.ResponseWriter, r *http.Request) {
	// Find the seasons
	fileBytes, err := ioutil.ReadFile("./static/seasons.txt")
	if err != nil {
		log.Fatal(err)
	}
	s := string(fileBytes)
	protoseasons := strings.Split(s, "\n")

	t, err := template.ParseGlob("./templates/*")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("\nSeason Table.")
	movingfinger.RecordPageHit("Season Table.")

	seasons := carpenter.Nest(protoseasons, 5, 3)
	t.ExecuteTemplate(w, "chooseSeasonTable.html", seasons)
}

// ServeSeasonGraph returns the season graph template
func ServeSeasonGraph(w http.ResponseWriter, r *http.Request) {
	// Find the seasons
	fileBytes, err := ioutil.ReadFile("./static/seasons.txt")
	if err != nil {
		log.Fatal(err)
	}
	s := string(fileBytes)
	protoseasons := strings.Split(s, "\n")

	t, err := template.ParseGlob("./templates/*")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("\nSeason Graph.")
	movingfinger.RecordPageHit("Season Graph.")

	seasons := carpenter.Nest(protoseasons, 5, 3)
	t.ExecuteTemplate(w, "chooseSeasonGraph.html", seasons)
}

// ServeAbout returns the about page
func ServeAbout(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseGlob("./templates/*")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("\nAbout.")

	t.ExecuteTemplate(w, "about.html", i)
}

// ServeContact returns the contact page
func ServeContact(w http.ResponseWriter, r *http.Request){
	t, err := template.ParseGlob("./templates/*")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("\nContact.")
	movingfinger.RecordPageHit("Contact.")

	t.ExecuteTemplate(w, "contact.html", i)

}

// SeasonTable serves the details of the season requested by the user. First, it
// checks that the request URL has parameters. If it doesn’t, home is served.
// If it does, the season is identified, a query string is prepared and a
// request is made to the database. The `season_table` template is then served
// with the results slice as its parameter.
func SeasonTable(w http.ResponseWriter, r *http.Request) {
	var g SeasonGraphDetail

	if len(r.URL.Query()) == 0 {
		log.Printf("\nNo Season Provided")
		t.ExecuteTemplate(w, "index.html", i)
	} else {
		// Find the query
		fileBytes, err := ioutil.ReadFile("./queries/points_v_goal_difference.sql")
		if err != nil {
			log.Fatal(err)
		}
		season := r.URL.Query()["season"][0]
		q := string(fileBytes)
		q = fmt.Sprintf(q, season)

		results := connectdb.PointsGdTableQuery(q)

		t, err := template.ParseGlob("./templates/*")
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("\nSeason %s.", season)
		pageHit := fmt.Sprintf("Season %s", season)
		movingfinger.RecordPageHit(pageHit)

		g.Title = season
		g.Results = results
		commas.CompileCsv(results)
		t.ExecuteTemplate(w, "season_table.html", g)
	}
}

// TeamGraph serves the details of the season requested by the user. First, it
// checks that the request URL has parameters. If it doesn’t, home is served.
// If it does, the season is identified, a query string is prepared and a
// request is made to the database. The results are found and then written to
// `data.csv`, where `graph.js` can access them. Finally, the `season_graph`
// template is served.
func TeamGraph(w http.ResponseWriter, r *http.Request) {
	var g SeasonGraphDetail
	if len(r.URL.Query()) == 0 {
		log.Printf("\nNo Season Provided")
		t.ExecuteTemplate(w, "index.html", i)
	} else {
		// Find the query
		fileBytes, err := ioutil.ReadFile("./queries/points_v_goal_difference.sql")
		if err != nil {
			log.Fatal(err)
		}
		season := r.URL.Query()["season"][0]
		q := string(fileBytes)
		q = fmt.Sprintf(q, season)

		results := connectdb.PointsGdTableQuery(q)

		t, err := template.ParseGlob("./templates/*")
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("\nSeason %s.", season)
		g.Title = season
		g.Results = results
		commas.CompileCsv(results)
		t.ExecuteTemplate(w, "season_graph.html", g)
	}
}


// SeasonGraph serves the details of the season requested by the user. First, it
// checks that the request URL has parameters. If it doesn’t, home is served.
// If it does, the season is identified, a query string is prepared and a
// request is made to the database. The results are found and then written to
// `data.csv`, where `graph.js` can access them. Finally, the `season_graph`
// template is served.
func SeasonGraph(w http.ResponseWriter, r *http.Request) {
	var g SeasonGraphDetail
	if len(r.URL.Query()) == 0 {
		log.Printf("\nNo Season Provided")
		t.ExecuteTemplate(w, "index.html", i)
	} else {
		// Find the query
		fileBytes, err := ioutil.ReadFile("./queries/points_v_goal_difference.sql")
		if err != nil {
			log.Fatal(err)
		}
		season := r.URL.Query()["season"][0]
		q := string(fileBytes)
		q = fmt.Sprintf(q, season)

		results := connectdb.PointsGdTableQuery(q)

		t, err := template.ParseGlob("./templates/*")
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("\nSeason %s.", season)
		g.Title = season
		g.Results = results
		commas.CompileCsv(results)
		t.ExecuteTemplate(w, "season_graph.html", g)
	}
}

// Team reads the seasons query and returns the every field
// in the points table relative to the season in mux.Vars(r). These
// are passed to the seasons.html template.
func Team(w http.ResponseWriter, r *http.Request) {
	// Find the query
	// fileBytes, err := ioutil.ReadFile("./queries/team.sql")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	vars := map[string]string{"team": "this is only a placeholder, you know"}
	team := vars["team"]
	// q := string(fileBytes)
	// q = fmt.Sprintf(q, team)

	// results := connectdb.SeasonQuery(q)

	results := connectdb.FindUniqueTeams()
	t, err := template.ParseGlob("./templates/*")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("\nTeam %s.", team)

	t.ExecuteTemplate(w, "team.html", results)
}



// RecordContact appends whatever's been written on the contact form to the
// contacts log
func RecordContact(w http.ResponseWriter, r *http.Request) {
	var c ContactName
	log.Printf("Record Contact")
	currentTime := time.Now()
	email := r.FormValue("email")
	header := fmt.Sprintf("%v", r.Header)
	name := r.FormValue("name")
	comment := r.FormValue("comment")
	commentTerminator := "\n********************************************************************************\n"

	formDetails := []string{currentTime.Format(time.RFC3339),
							email,
							name,
							header,
							comment,
							commentTerminator}
							
	fmt.Println(formDetails)
	details := strings.Join(formDetails, "\n")

    f, err := os.OpenFile("comments.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }
    if _, err := f.Write([]byte(details)); err != nil {
        log.Fatal(err)
    }
    if err := f.Close(); err != nil {
        log.Fatal(err)
	}
	
	names := strings.Split(name, " ")
	c.Firstname = names[0]

	t, err := template.ParseGlob("./templates/*")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("\nConfirm Contact.")

	t.ExecuteTemplate(w, "confirm_contact.html", c)

}
