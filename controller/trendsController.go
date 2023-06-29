package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func getGoogleTrends() *http.Response {
	res, err := http.Get("https://trends.google.com/trends/trendingsearches/daily/rss?geo=US")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return res
}

func ReadGoogleTrends() []byte {
	results := getGoogleTrends()
	data, err := ioutil.ReadAll(results.Body)
	if err != nil {
		return nil
	}
	return data

}
