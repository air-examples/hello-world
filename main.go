package main

import "github.com/aofei/air"

func main() {
	air.GET("/", func(req *air.Request, res *air.Response) error {
		return res.WriteString("Hello, 世界")
	})
	air.Serve()
}
