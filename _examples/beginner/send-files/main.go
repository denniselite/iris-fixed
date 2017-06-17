package main

import (
	"github.com/denniselite/iris-fixed"
	"github.com/denniselite/iris-fixed/context"
)

func main() {
	app := iris.New()

	app.Get("/", func(ctx context.Context) {
		file := "./files/first.zip"
		ctx.SendFile(file, "c.zip")
	})

	app.Run(iris.Addr(":8080"))
}
