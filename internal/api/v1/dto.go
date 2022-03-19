package v1

import "crawler/internal/crawler"

type URLListDTO struct {
	List []string `json:"list"`
}

type TitleListDTO struct {
	List []crawler.Title `json:"list"`
}
