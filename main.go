package main

import (
	"github.com/quotebots/quotesbot"
	"github.com/wetrycode/tegenaria"
)

func main() {
	engine := tegenaria.NewSpiderEngine(tegenaria.EngineWithUniqueReq(false))
	spider := &quotesbot.QuotesbotSpider{
		Name:     "quote_bot",
		FeedUrls: []string{"http://quotes.toscrape.com/"},
	}
	engine.RegisterSpider(spider)
	pipe1 := quotesbot.QuotesbotItemPipeline{Priority: 1}
	engine.RegisterPipelines(&pipe1)
	engine.RegisterPipelines(&quotesbot.QuotesbotItemPipeline2{Priority: 2})
	engine.RegisterPipelines(&quotesbot.QuotesbotItemPipeline3{Priority: 3})

	engine.Start("quote_bot")
	engine.Close()

}
