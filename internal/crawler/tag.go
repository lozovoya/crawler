package crawler

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type crawServ struct {
	inputChan chan string
	outChan chan Title
}

type Title struct {
	Url string
	Title string
}

func NewCrawlerService () Tag {
	var in = make(chan string)
	var out = make(chan Title)
	return &crawServ{inputChan: in, outChan: out}
}

func (c *crawServ) Crawler (urls []string) ([]Title, error) {
	var list = make([]Title, 0)
	go func() {
		for _, item := range urls {
			c.inputChan <- item
		}
		close(c.inputChan)
	}()
	for result := range c.outChan {
		list = append(list, result)
		if len(list) == len(urls) {
			close(c.outChan)
			break
		}
	}
	return list, nil
}

func (c *crawServ) CrawlerWorker () {
	for url := range c.inputChan {
		title, err := c.getTag(url, "title")
		if err != nil {
			log.Println(err)
		}
		var result = Title{
			Url:   url,
			Title: title,
		}
		c.outChan <- result
	}
}

func (c *crawServ) getTag (url string, tag string) (string, error) {
	var result string
	client := http.Client{
		Timeout: time.Second*30,
	}
	response, err := client.Get(url)
	if err != nil {
		return result, err
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return result, err
	}
	dataInString := string(data)
	titleStart := strings.Index(dataInString, "<"+tag+">")
	if titleStart == -1 {
		return result, errors.New("no tag")
	}
	titleStart += len(tag)+2
	titleEnd := strings.Index(dataInString, "</"+tag+">")
	if titleStart == -1 {
		return result, errors.New("closing tag not found")
	}
	result = dataInString[titleStart:titleEnd]
	return result, nil
}


