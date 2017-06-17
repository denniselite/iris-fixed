package main

import (
	"github.com/denniselite/iris-fixed"
	"github.com/denniselite/iris-fixed/context"
)

func main() {
	app := iris.New()
	app.Handle("GET", "/", func(ctx context.Context) {
		ctx.HTML("<b> Hello world! </b>")
	})
	app.Run(iris.Addr(":8080"))
}
