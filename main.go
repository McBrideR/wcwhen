package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/McBrideR/wcwhen/scraper"
)

//to set up go
//GOROOT=/usr/local/Cellar/go/1.12.9/libexec #gosetup
//GOPATH=/Users/rossmcbride/go #gosetup

//to build package with different executable name
//go build -o wcwhen

//talk to scott about vendoring

var country string
var groupName string
var venue string
var dateTime time.Time
var team1 string

var team2 string

func main() {
	flag.StringVar(&country, "country", "Ireland", "Country you want to see matches for.")
	flag.StringVar(&groupName, "groupName", "A", "Group Name you want to see matches for.")
	flag.Parse()
	fmt.Println("Hello, World")
	fmt.Println(country)
	scraper.GetMatchStats()
}
