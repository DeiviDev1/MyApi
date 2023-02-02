package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	g "github.com/gin-gonic/gin"
)

type Coin struct {
	ID      string `json:"id"`
	Content struct {
		Price    float64 `json:"price"`
		Currency string  `json:"currency"`
	} `json:"content"`
	Partial bool `json:"partial"`
}

func main() {
	r := g.Default()

	r.GET("/myapi", func(ctx *g.Context) {
		res, err := http.Get("https://www.coingecko.com/es/api ")
		if err != nil {
			ctx.JSON(http.StatusPartialContent, Coin{Partial: true})
			return
		}

		body := &bytes.Buffer{}
		_, err = io.Copy(body, res.Body)
		if err != nil {
			ctx.JSON(http.StatusPartialContent, Coin{Partial: true})
			return
		}

		var result Coin
		json.Unmarshal(body.Bytes(), &result)

		ctx.JSON(http.StatusOK, result)
	})

	r.Run()
}
