package main

// import (
// 	"fmt"
// 	"sync"

// 	queue "github.com/yireyun/go-queue"
// )

import (
	"fmt"
	"os"
	"runtime/pprof"
	"time"

	"github.com/geebytes/tegenaria"
	"github.com/quotebots/quotesbot"
)

func main() {
	f, _ := os.OpenFile("cpu.pprof", os.O_CREATE|os.O_RDWR, 0644)
	defer f.Close()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	// go func() {
	// 	http.ListenAndServe("0.0.0.0:8899", nil)
	// }()
	engine := tegenaria.NewSpiderEngine(tegenaria.EngineWithUniqueReq(false), tegenaria.EngineWithConcurrencyNum(32))
	spider := &quotesbot.QuotesbotSpider{
		Name:     "quote_bot",
		FeedUrls: []string{"http://quotes.toscrape.com/"},
	}
	engine.RegisterSpider(spider)
	pipe1 := quotesbot.QuotesbotItemPipeline{Priority: 1}
	engine.RegisterPipelines(&pipe1)
	engine.RegisterPipelines(&quotesbot.QuotesbotItemPipeline2{Priority: 2})
	engine.RegisterPipelines(&quotesbot.QuotesbotItemPipeline3{Priority: 3})
	now := time.Now()

	engine.Start("quote_bot")
	runTime := time.Since(now).Seconds()
	fmt.Printf("任务执行时间%f\n", runTime)
	engine.Close()

}
