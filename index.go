package main

import (
	"log"
	"os"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"html/template"
)

var respString string = ""
var cachedNews news

const url string = "https://newsapi.org/v1/articles?source=bbc-news&sortBy=top&apiKey=87295cd4d5d4415c9cb4896590ccf2c3"

func handleRoot(resp http.ResponseWriter, req *http.Request) {
	
	renderTemplate("templates/home.html", cachedNews, resp)
}


func renderTemplate(fileName string, data interface{}, resp http.ResponseWriter) {
	tmpl, err := template.ParseFiles(fileName)

	if err != nil {
		log.Fatalln("Error in Reading Template %s %s", fileName, err)
	}
	e := tmpl.Execute(resp, data)
	if e != nil {
		log.Fatalln("Failed to Execute Tempalte")
	}

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

	resp.Body.Close()
	cachedNews = newsResults
	
	
	
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handleRoot)
	port := ""

	if os.Getenv("PORT") != "" {
		port = ":"+os.Getenv("PORT")
	} else {
		port = ":9080"
	}

	fmt.Println("Hello we are here", port)

	fmt.Printf("Server has started, browse http://localhost%s to check out news", port)
	http.ListenAndServe(port, nil)

}
