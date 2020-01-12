package dbconnect

import (
	"github.com/amunnelly/dbconnect/connector"
	"database/sql"
	"fmt"

	// postgres driver
	_ "github.com/lib/pq"
)

// Db is a type that extends sql.DB. Is there anything to be done with this?
type Db sql.DB

// Team is a struct with one property, Team, a string
type Team struct {
	Team string
}

// Teams returns a lot of stuff
type Teams struct {
	Date     string
	HomeTeam string
	AwayTeam string
	FTHG     int32
	FTAG     int32
	FTR      rune
	Season   string
}

// Points is a struct for holding a row from the points table of the epl
// database
type Points struct {
	Season         string
	Date           string
	Team           string
	Venue          string
	Opposition     string
	Points         int32
	RunningPoints int32
	GoalsFor       int32
	GoalsAgainst   int32
	GoalDifference int32
}

// PointsGdTable is a struct for holding the values returned by the
// `points_v_goal_difference.sql` query.
type PointsGdTable struct {
	Team         string
	Points       int32
	GD           int32
	GoalsFor     int32
	GoalsAgainst int32
}


func connectToDb() *sql.DB {
	// check if program is in heroku
	if os.Getenv("DATABASE_URL") != "" {
		db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
if err != nil {
    log.Fatal(err)
	}
	
} else {


	deets := connector.LoadDetails()

	db, err := sql.Open("postgres", deets)
	if err != nil {
		fmt.Println("Failed to connect")
		panic(err)
	}

	return db
}
}

func runQuery(q string) *sql.Rows {
	db := connectToDb()

	defer db.Close()

	rows, err := db.Query(q)

	if err != nil {
		panic(err)
	}
	return rows
}

// FindUniqueTeams returns an array of the distinct teams in the points table
// of the database.
func FindUniqueTeams() []Team {
	q := "./queries/distinctTeams.sql"
	rows := runQuery(q)
	defer rows.Close()

	holder := []Team{}

	for rows.Next() {
		var temp Team
		err := rows.Scan(&temp.Team)
		if err != nil {
			panic(err)
		}
		holder = append(holder, temp)
	}
	return holder
}

// SeasonQuery returns all the fields from the points table for a given season.
func SeasonQuery(q string) []Points {
	fmt.Println(q)
	rows := runQuery(q)
	defer rows.Close()

	holder := []Points{}

	for rows.Next() {
		var temp Points
		err := rows.Scan(&temp.Season,
			&temp.Date,
			&temp.Team,
			&temp.Venue,
			&temp.Opposition,
			&temp.Points,
			&temp.GoalsFor,
			&temp.GoalsAgainst,
			&temp.GoalDifference)
		if err != nil {
			panic(err)
		}
		holder = append(holder, temp)
	}
	return holder
}

// TeamQuery returns the season details for a given team for every season
// that team has spent in the Premier League
func TeamQuery(q string) []Points {
	fmt.Println(q)
	rows := runQuery(q)
	defer rows.Close()

	holder := []Points{}

	for rows.Next() {
		var temp Points
		err := rows.Scan(&temp.Season,
			&temp.Date,
			&temp.Venue,
			&temp.Opposition,
			&temp.Points,
			&temp.GoalsFor,
			&temp.GoalsAgainst,
			&temp.GoalDifference)
		if err != nil {
			panic(err)
		}
		holder = append(holder, temp)
	}
	return holder
}

// PointsGdTableQuery returns the final table for a particular season,
// listing teams, goals for, goals against, goal difference and points.
func PointsGdTableQuery(q string) []PointsGdTable {

	rows := runQuery(q)
	defer rows.Close()

	holder := []PointsGdTable{}

	for rows.Next() {
		var temp PointsGdTable
		err := rows.Scan(&temp.Team,
			&temp.Points,
			&temp.GD,
			&temp.GoalsFor,
			&temp.GoalsAgainst)
		if err != nil {
			panic(err)
		}
		holder = append(holder, temp)
	}
	return holder
}
