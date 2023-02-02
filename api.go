package main

import (

	//"net/http"

	g "github.com/gin-gonic/gin"
)

type Data struct {
	Data string `json:"data"`
}

func main() {

	r := g.Default()
	r.GET("myapi", func(ctx *g.Context) {
		dados := Data{
			Data: "hello",
		}
		ctx.JSON(200, g.H{
			"dados": dados.Data,
		})
	})
	r.Run()
}
