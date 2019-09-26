package scraper

import (
	"fmt"
	"net/http"
)

func GetMatchStats() {
	resp, _ := http.Get("http://example.com/")

	fmt.Println(resp)
}
