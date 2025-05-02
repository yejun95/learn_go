package scrapper

import (
	"encoding/csv"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type extractedJob struct {
	id          string
	badge       string
	title       string
	corporation string
}

func Scrape(term string) {
	baseURL := "https://www.saramin.co.kr/zf_user/search/recruit?&searchword=" + term + ""
	c := make(chan []extractedJob)
	var jobs []extractedJob
	totalPages := getPages(baseURL)
	fmt.Println("totalPages :", totalPages)

	for i := 1; i <= totalPages; i++ {
		go getPage(i, baseURL, c) // 페이지별 공고가 담김
	}

	for i := 0; i < totalPages; i++ {
		extractedJobs := <-c
		jobs = append(jobs, extractedJobs...) // 배열의 콘텐츠를 모아 다시 배열로 묶음 -> [] + [] + [] + [] = []
	}
	writeJobs(jobs)
	fmt.Println("Done, extracted :", len(jobs))
}

func getPage(page int, baseURL string, mainC chan<- []extractedJob) {
	var jobs []extractedJob
	c := make(chan extractedJob)

	pageURL := baseURL + "&recruitPage=" + strconv.Itoa(page)
	fmt.Println("Requesting", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".item_recruit")
	searchCards.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, c)
	})

	for i := 0; i < searchCards.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}

	mainC <- jobs
}

func extractJob(card *goquery.Selection, c chan<- extractedJob) {
	id, _ := card.Attr("value")
	badge := CleanString(card.Find(".area_badge>span").Text())
	title := CleanString(card.Find(".job_tit>a").Text())
	corporation := CleanString(card.Find(".area_corp>strong").Text())
	c <- extractedJob{id, badge, title, corporation}
}

func CleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

func getPages(baseURL string) int {
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

func writeJobs(jobs []extractedJob) {
	baseDetailURL := "https://www.saramin.co.kr/zf_user/jobs/relay/view?isMypage=no&rec_idx="

	file, err := os.Create("jobs.csv")
	checkErr(err)

	// Excel에서 인식 가능한 UTF-8 with BOM 추가
	file.Write([]byte{0xEF, 0xBB, 0xBF})

	w := csv.NewWriter(file)
	defer w.Flush() // 함수가 끝나는 시점에 파일에 데이터를 입력하는 함수

	headers := []string{"Link", "Badge", "Title", "Corporation"}

	wErr := w.Write(headers)
	checkErr(wErr)

	// loop가 끝나면 defer 함수 실행을 통해 데이터가 파일에 입력
	for _, job := range jobs {
		jobSlice := []string{baseDetailURL + job.id, job.badge, job.title, job.corporation}
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)
	}
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
