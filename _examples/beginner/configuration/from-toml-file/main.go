package main

import (
	"github.com/denniselite/iris-fixed"
)

func main() {
	app := iris.New()

	// [...]

	// Good when you have two configurations, one for development and a different one for production use.
	app.Run(iris.Addr(":8080"), iris.WithConfiguration(iris.TOML("./configs/iris.tml")))

	// or before run:
	// app.Configure(iris.WithConfiguration(iris.TOML("./configs/iris.tml")))
	// app.Run(iris.Addr(":8080"))
}
