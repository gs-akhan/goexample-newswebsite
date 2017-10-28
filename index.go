package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var respString string = ""

const url string = "https://newsapi.org/v1/articles?source=bbc-news&sortBy=top&apiKey=87295cd4d5d4415c9cb4896590ccf2c3"

func handleRoot(resp http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(resp, respString)
}

func processString(item newsItem) string {
	return fmt.Sprintf(`<div>
			<a href = "%s">%s</a>
		</div>`, item.Url, item.Title)
}

type news struct {
	Status   string     `json:"status"`
	Source   string     `json:"source"`
	SortBy   string     `json:"sortBy"`
	Articles []newsItem `json:"articles"`
}

type newsItem struct {
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	UrlToImage  string `json:"urlToImage"`
	PublishedAt string `json:"publishedAt"`
}

func main() {
	var newsResults news
	resp, _ := http.Get(url)
	bytes, _ := ioutil.ReadAll(resp.Body)

	err := json.Unmarshal(bytes, &newsResults)

	if err != nil {
		fmt.Println("We have an error")
		return
	}

	for _, item := range newsResults.Articles {
		respString += processString(item)
	}
	resp.Body.Close()

	fmt.Println("Server has started, browse http://localhost:9080 to check out news")
	http.HandleFunc("/", handleRoot)
	http.ListenAndServe(":9080", nil)

}
