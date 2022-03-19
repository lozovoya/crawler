package v1

import (
	"crawler/internal/crawler"
	"encoding/json"
	"log"
	"net/http"
)

type Crawler struct {
	CrawlerServ crawler.Tag
}

func NewCrawlerController (CrawlerServ crawler.Tag) *Crawler {
	return &Crawler{CrawlerServ: CrawlerServ}
}

func (c *Crawler) GetTitles (writer http.ResponseWriter, request *http.Request) {
	var data URLListDTO
	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		log.Println(err)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
	list, err := c.CrawlerServ.Crawler(data.List)
	writer.Header().Set("Content-Type", "application/json")

	var Reply = TitleListDTO{List: list}
	err = json.NewEncoder(writer).Encode(Reply)
	if err != nil {
		log.Println(err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}