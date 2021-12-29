package quotesbot

import (
	"net/url"
	"strings"

	"github.com/antchfx/htmlquery"
	"github.com/geebytes/tegenaria"
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

func (d *QuotesbotSpider) StartRequest(req chan<- *tegenaria.Request) {
	for i := 0; i < 10000; i++ {
		for _, url := range d.FeedUrls {
			request := tegenaria.NewRequest(url, tegenaria.GET, d.Parser)
			req <- request
		}
	}

}
func (d *QuotesbotSpider) Parser(response *tegenaria.Response, item chan<- tegenaria.ItemInterface, req chan<- *tegenaria.Request) {
	text := response.String()
	doc, _ := htmlquery.Parse(strings.NewReader(text))
	list, err := htmlquery.QueryAll(doc, "//div[@class='quote']")
	if err != nil {
		panic(err)
	}
	for _, n := range list {
		t := htmlquery.FindOne(n, "//span[@class='text']")
		var quoteText string = ""
		var quoteAuthor string = ""
		var quoteTags string = ""
		if t != nil {
			quoteText = htmlquery.InnerText(t)
		}
		author := htmlquery.FindOne(n, "//small[@class='author']")
		if author != nil {
			quoteAuthor = htmlquery.InnerText(author)
		}
		tags := htmlquery.FindOne(n, "//div[@class='tags']/a[@class='tag']")
		if tags != nil {
			quoteTags = htmlquery.InnerText(tags)
		}
		var quoteItem = QuotesbotItem{
			Text:   quoteText,
			Author: quoteAuthor,
			Tags:   quoteTags,
		}
		item <- &quoteItem
	}
	doamin_url := response.Req.Url
	u, _ := url.Parse(doamin_url)

	var nextPageUrl string = ""
	nextUrl := htmlquery.FindOne(doc, "//li[@class='next']/a")
	if nextUrl != nil {
		nextPageUrl = htmlquery.SelectAttr(nextUrl, "href")
		next, _ := url.Parse(nextPageUrl)
		s := u.ResolveReference(next).String()
		request := tegenaria.NewRequest(s, tegenaria.GET, d.Parser)
		req <- request
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

func (p *QuotesbotItemPipeline) ProcessItem(spider tegenaria.SpiderInterface, item tegenaria.ItemInterface) error {
	// fmt.Printf("Spider %s run QuotesbotItemPipeline priority is %d\n", spider.GetName(), p.GetPriority())
	return nil

}
func (p *QuotesbotItemPipeline) GetPriority() int {
	return p.Priority
}
func (p *QuotesbotItemPipeline2) ProcessItem(spider tegenaria.SpiderInterface, item tegenaria.ItemInterface) error {
	// fmt.Printf("Spider %s run QuotesbotItemPipeline2 priority is %d\n", spider.GetName(), p.GetPriority())
	return nil
}
func (p *QuotesbotItemPipeline2) GetPriority() int {
	return p.Priority
}

func (p *QuotesbotItemPipeline3) ProcessItem(spider tegenaria.SpiderInterface, item tegenaria.ItemInterface) error {
	// fmt.Printf("Spider %s run QuotesbotItemPipeline3 priority is %d\n", spider.GetName(), p.GetPriority())
	return nil

}
func (p *QuotesbotItemPipeline3) GetPriority() int {
	return p.Priority
}
