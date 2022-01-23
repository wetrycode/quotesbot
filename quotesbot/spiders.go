package quotesbot

import (
	"log"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/wetrycode/tegenaria"
)
// QuotesbotSpider an example of tegenaria spider
// It is an instance of tegenaria.SpiderInterface interface
type QuotesbotSpider struct {
	Name     string
	FeedUrls []string
}
// QuotesbotSpider an example of tegenaria item
type QuotesbotItem struct {
	Text   string
	Author string
	Tags   string
}

// StartRequest funcation of tegenaria.SpiderInterface interface
// send feeds url request to engine
func (d *QuotesbotSpider) StartRequest(req chan<- *tegenaria.Context) {
	for i := 0; i < 10000; i++ {
		for _, url := range d.FeedUrls {
			// get a new request
			request := tegenaria.NewRequest(url, tegenaria.GET, d.Parser)
			// get request context
			ctx := tegenaria.NewContext(request)
			// send request context to engine
			req <- ctx
		}
	}

}
// Parser funcation of tegenaria.SpiderInterface interface
// recvie request download response context
// and it will send parse result as an item to engine
// it also will send a new request context to engine
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
		// ready to send a item to engine
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
			// ready to send a new request context to engine
			newRequest := tegenaria.NewRequest(s, tegenaria.GET, d.Parser)
			newCtx := tegenaria.NewContext(newRequest)
			req <- newCtx
		}
	}

}
// ErrorHandler handler of error
// it will recvie a tegenaria.HandleError and you can do some handler of this error
func (d *QuotesbotSpider)ErrorHandler(err *tegenaria.HandleError){

}
// GetName get spider name
func (d *QuotesbotSpider) GetName() string {
	return d.Name
}

// QuotesbotItemPipeline an example of tegenaria.PipelinesInterface interface
// You need to set priority of each piplines
// These pipeline will handle all item order by priority
// Priority is that the lower the number, the higher the priority
type QuotesbotItemPipeline struct {
	Priority int
}
type QuotesbotItemPipeline2 struct {
	Priority int
}
type QuotesbotItemPipeline3 struct {
	Priority int
}
// ProcessItem funcation of tegenaria.PipelinesInterface,it is used to handle item
func (p *QuotesbotItemPipeline) ProcessItem(spider tegenaria.SpiderInterface, item *tegenaria.ItemMeta) error {
	return nil

}
// GetPriority get priority of pipline
func (p *QuotesbotItemPipeline) GetPriority() int {
	return p.Priority
}
func (p *QuotesbotItemPipeline2) ProcessItem(spider tegenaria.SpiderInterface, item *tegenaria.ItemMeta) error {
	return nil
}
func (p *QuotesbotItemPipeline2) GetPriority() int {
	return p.Priority
}

func (p *QuotesbotItemPipeline3) ProcessItem(spider tegenaria.SpiderInterface, item *tegenaria.ItemMeta) error {
	return nil

}
func (p *QuotesbotItemPipeline3) GetPriority() int {
	return p.Priority
}
