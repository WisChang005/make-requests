package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"gopkg.in/ini.v1"
)

func main() {
	sungrilNewsDomain := "https://www.sungirlbaby.com/news/"
	newsIDs := strings.Split(getKeyFromConfig("Sungirl", "NewsIDs"), ",")
	repeats, err := strconv.Atoi(getKeyFromConfig("Sungirl", "Repeats"))
	errorCheck(err, "Parse repeats to int failed")

	for i := 1; i <= repeats; i++ {
		fmt.Printf("================== Round %d ==================\n", i)
		for _, newsID := range newsIDs {
			resp := requestsUrl(sungrilNewsDomain + newsID)
			showPageTitle(resp)
		}
	}
	fmt.Println("=============================================")
	fmt.Printf("\nPages add [ %d ] views, 五秒後自動關閉...\n", repeats)

	pauseTime := time.Duration(5) * time.Second
	time.Sleep(pauseTime)
}

func errorCheck(err error, errMsg string) {
	if err != nil {
		if errMsg != "" {
			panic(errMsg)
		}
	}
}

func requestsUrl(theUrl string) *http.Response {
	resp, err := http.Get(theUrl)
	if resp.StatusCode != 200 {
		fmt.Printf("Status Code: [%d]", resp.StatusCode)
	}
	errorCheck(err, "Request URL Failed")
	return resp
}

func getKeyFromConfig(sectionName, keyName string) string {
	conf, err := ini.Load("config.ini")
	errorCheck(err, "Read config.ini failed")
	return conf.Section(sectionName).Key(keyName).String()
}

func showPageTitle(resp *http.Response) {
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	errorCheck(err, "Html parse failed")

	doc.Find("meta").Each(func(i int, s *goquery.Selection) {
		if property, _ := s.Attr("property"); property == "og:title" {
			pageTitle, _ := s.Attr("content")
			if pageTitle != "" {
				fmt.Printf("[%s] - Done!\n", pageTitle)
			}
		}
	})

}
