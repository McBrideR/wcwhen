package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

//to set up go
//GOROOT=/usr/local/Cellar/go/1.12.9/libexec #gosetup
//GOPATH=/Users/rossmcbride/go #gosetup

//to build package with different executable name
//go build -o wcwhen

//talk to scott about vendoring

type time struct {
	millis int64
	gmtOffset int
	label string
}

type event struct {
	id int `json:"id"`
	altId int
	label string
	sport string `json:"sport"`
	start time
	end time
	rankingsWeight int
	abbr string
	winningTeam string
	impactPlayers string
}

type venue struct {
	id int
	name string
	city string
	country string
}

type team struct {
	id int
	altId int
	name string
	abbreviation string
	annotations string
}

type weather struct {
	matchWeather string
	matchMinTemperature int
	matchMaxTemperature int
	matchWindConditions float32
	matchPitchConditions float32
}

type match struct {
	matchId int
	altId int
	description string
	eventPhase string
	venue venue
	time time
	attendance int
	teams []team
	scores []int
	kc string
	status string
	clock string
	outcome string
	events event
	sport string
	competition string
	weather weather
}

type matchResponseObject struct {
	event event `json:"event"`
	matches []match `json:"matches"`
}

var reqCountry string
var reqGroupName string
var reqVenue string
var reqDateTime string
var reqTeam string

var matchStatsObject matchResponseObject

func main() {
	addFlags()
	makeCalls()
}

func addFlags() {
	flag.StringVar(&reqCountry, "country", "Ireland", "Country you want to see matches for.")
	flag.StringVar(&reqGroupName, "groupName", "A", "Group Name you want to see matches for.")

	flag.Parse()
}

func makeCalls()  {
	getMatches()

	if(reqCountry != "") {
		checkMatchesForCountry()
	} else if(reqGroupName != "") {
		checkMatchesForGroupName()
	} else if(reqVenue != "") {
		checkMatchesForCountry()
	} else if(reqDateTime != "") {
		checkMatchesForDateTime()
	} else if(reqTeam != "") {
		checkMatchesForTeam()
	} else {
		fmt.Println("add some parameters fool...")
	}
}

func checkMatchesForCountry()  {
	fmt.Println(matchStatsObject.event.sport)
	fmt.Println("i did get here at least!")
}

func checkMatchesForGroupName()  {

}

func checkMatchesForDateTime()  {

}

func checkMatchesForTeam()  {

}

func getMatches() {
	matchStats, _ := http.Get("https://cmsapi.pulselive.com/rugby/event/1558/schedule")
	body, err := ioutil.ReadAll(matchStats.Body)
	if err != nil {
		fmt.Println("meh")
	}
	err = json.Unmarshal(body, &matchStatsObject)
	if err != nil {
		fmt.Println("ot oh scotty")
	}
	fmt.Println(matchStatsObject)
}