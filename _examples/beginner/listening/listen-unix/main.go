package main

import (
	"github.com/denniselite/iris-fixed"
	"github.com/denniselite/iris-fixed/core/nettools"
)

func main() {
	app := iris.New()

	l, err := nettools.UNIX("/tmpl/srv.sock", 0666) // see its code to see how you can manually create a new file listener, it's easy.
	if err != nil {
		panic(err)
	}

	app.Run(iris.Listener(l))
}
