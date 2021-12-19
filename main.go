package main

import (
	"github.com/geebytes/tegenaria"
	"github.com/quotebots/quotesbot"
)

func main() {
	enginer := tegenaria.NewSpiderEnginer()
	spider := &quotesbot.QuotesbotSpider{
		Name:     "quote_bot",
		FeedUrls: []string{"http://quotes.toscrape.com/"},
	}
	enginer.RegisterSpider(spider)
	pipe1 := quotesbot.QuotesbotItemPipeline{Priority: 1}
	enginer.RegisterPipelines(&pipe1)
	enginer.RegisterPipelines(&quotesbot.QuotesbotItemPipeline2{Priority: 2})
	enginer.RegisterPipelines(&quotesbot.QuotesbotItemPipeline3{Priority: 3})
	enginer.Start("quote_bot")
	enginer.Close()

}
