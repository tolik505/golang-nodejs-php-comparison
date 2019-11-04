package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/fzzy/radix/redis"
)

func errHndlr(err error) {
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}

func main() {
	c, err := redis.DialTimeout("tcp", "172.17.0.1:6379", time.Duration(10)*time.Second)
	defer c.Close()

	errHndlr(err)

	c.Cmd("select", 8)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		s, _ := c.Cmd("get", "count").Int()

		s++
		fmt.Fprintf(w, "%n", s)

		c.Cmd("set", "count", s)
	})
	http.ListenAndServe(":8080", nil)
}
