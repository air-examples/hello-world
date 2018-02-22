package main

import "github.com/sheng/air"

func main() {
	air.GET("/", func(req *air.Request, res *air.Response) error {
		return res.String("Hello, 世界")
	})
	air.Serve()
}
