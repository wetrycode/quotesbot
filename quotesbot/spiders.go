package quotesbot

import (
	"log"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/wetrycode/tegenaria"
)

type QuotesbotSpider struct {
	Name     string
	FeedUrls []string
}
type QuotesbotItem struct {
	Text   string
	Author string
	Tags   string
}

func (d *QuotesbotSpider) StartRequest(req chan<- *tegenaria.Context) {
	for i := 0; i < 10000; i++ {
		for _, url := range d.FeedUrls {
			request := tegenaria.NewRequest(url, tegenaria.GET, d.Parser)
			ctx := tegenaria.NewContext(request)
			req <- ctx
		}
	}

}
func (d *QuotesbotSpider) Parser(resp *tegenaria.Context, item chan<- *tegenaria.ItemMeta, req chan<- *tegenaria.Context) {
	text := resp.DownloadResult.Response.String()

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(text))

	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".quote").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		qText := s.Find(".text").Text()
		author := s.Find(".author").Text()
		tags := make([]string, 0)
		s.Find("a.tag").Each(func(i int, s *goquery.Selection) {
			tags = append(tags, s.Text())
		})
		var quoteItem = QuotesbotItem{
			Text:   qText,
			Author: author,
			Tags:   strings.Join(tags, ","),
		}
		itemCtx := tegenaria.NewItem(resp, &quoteItem)
		item <- itemCtx
	})
	doamin_url := resp.Request.Url
	next := doc.Find("li.next")
	if next != nil {
		nextUrl, ok := next.Find("a").Attr("href")
		if ok {
			u, _ := url.Parse(doamin_url)

			nextInfo, _ := url.Parse(nextUrl)
			s := u.ResolveReference(nextInfo).String()
			newRequest := tegenaria.NewRequest(s, tegenaria.GET, d.Parser)
			newCtx := tegenaria.NewContext(newRequest)
			req <- newCtx
		}
	}

}
func (d *QuotesbotSpider) ErrorHandler() {

}
func (d *QuotesbotSpider) GetName() string {
	return d.Name
}

type QuotesbotItemPipeline struct {
	Priority int
}
type QuotesbotItemPipeline2 struct {
	Priority int
}
type QuotesbotItemPipeline3 struct {
	Priority int
}

func (p *QuotesbotItemPipeline) ProcessItem(spider tegenaria.SpiderInterface, item *tegenaria.ItemMeta) error {
	// fmt.Printf("Spider %s run QuotesbotItemPipeline priority is %d\n", spider.GetName(), p.GetPriority())
	return nil

}
func (p *QuotesbotItemPipeline) GetPriority() int {
	return p.Priority
}
func (p *QuotesbotItemPipeline2) ProcessItem(spider tegenaria.SpiderInterface, item *tegenaria.ItemMeta) error {
	// fmt.Printf("Spider %s run QuotesbotItemPipeline2 priority is %d\n", spider.GetName(), p.GetPriority())
	return nil
}
func (p *QuotesbotItemPipeline2) GetPriority() int {
	return p.Priority
}

func (p *QuotesbotItemPipeline3) ProcessItem(spider tegenaria.SpiderInterface, item *tegenaria.ItemMeta) error {
	// fmt.Printf("Spider %s run QuotesbotItemPipeline3 priority is %d\n", spider.GetName(), p.GetPriority())
	return nil

}
func (p *QuotesbotItemPipeline3) GetPriority() int {
	return p.Priority
}
