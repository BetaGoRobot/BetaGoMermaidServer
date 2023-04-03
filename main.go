package main

import (
	"context"
	"flag"
	"log"

	mermaid_go "github.com/dreampuf/mermaid.go"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

var addr = flag.String("addr", ":8088", "TCP address to listen to")

func generateGraph(content string) []byte {
	re, err := mermaid_go.NewRenderEngine(context.Background())
	if err != nil {
		log.Panicln(err.Error())
	}
	defer re.Cancel()

	// content := `pie title pie_graph
	// "dimension 1": 60
	// "dimension 2": 40
	// "dimension 3": 100`

	pngInBytes, _, err := re.RenderAsPng(content)
	if err != nil {
		log.Panicln(err.Error())
	}
	return pngInBytes
}

func requestHandler(ctx *fasthttp.RequestCtx) {
	input := ctx.Request.Body()

	ctx.Response.AppendBody(generateGraph(string(input)))
}

func main() {
	r := router.New()
	r.POST("/make-charts", requestHandler)

	log.Fatal(fasthttp.ListenAndServe(*addr, r.Handler))
}
