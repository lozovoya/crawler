package crawler

type Tag interface  {
	Crawler (urls []string) ([]Title, error)
	CrawlerWorker ()
}
