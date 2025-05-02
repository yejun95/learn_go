package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type extractedJob struct {
	id          string
	badge       string
	title       string
	corporation string
}

var baseURL string = "https://www.saramin.co.kr/zf_user/search/recruit?&searchword=backend"
var baseDetailURL string = "https://www.saramin.co.kr/zf_user/jobs/relay/view?isMypage=no"

func main() {
	var jobs []extractedJob
	totalPages := getPages()
	fmt.Println("totalPages :", totalPages)

	for i := 1; i <= totalPages; i++ {
		extractedJobs := getPage(i)
		jobs = append(jobs, extractedJobs...) // 배열의 콘텐츠를 모아 다시 배열로 묶음 -> [] + [] + [] + [] = []
	}
	fmt.Println(jobs)
}

func getPage(page int) []extractedJob {
	var jobs []extractedJob
	pageURL := baseURL + "&recruitPage=" + strconv.Itoa(page)
	fmt.Println("Requesting", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".item_recruit").Each(func(i int, card *goquery.Selection) {
		job := extractJob(card)
		jobs = append(jobs, job)
	})

	return jobs
}

func extractJob(card *goquery.Selection) extractedJob {
	//detailURL := baseDetailURL + "&rec_idx=" + id
	id, _ := card.Attr("value")
	badge := cleanString(card.Find(".area_badge>span").Text())
	title := cleanString(card.Find(".job_tit>a").Text())
	corporation := cleanString(card.Find(".area_corp>strong").Text())
	return extractedJob{id, badge, title, corporation}
}

func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

func getPages() int {
	pages := 0

	res, err := http.Get(baseURL)

	checkErr(err)
	checkCode(res)
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
		//fmt.Println(pages)
	})

	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status : ", res.StatusCode)
	}
}
