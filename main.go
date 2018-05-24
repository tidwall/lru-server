package main

import (
	"flag"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/hashicorp/golang-lru"
	"github.com/tidwall/modern-server"
)

func main() {
	var size int
	var cache *lru.Cache
	opts := &server.Options{
		Version:     "0.0.1",
		Name:        "lru-server",
		Flags:       func() { flag.IntVar(&size, "s", 1000000, "") },
		FlagsParsed: func() { cache, _ = lru.New(size) },
		Usage: func(s string) string {
			return strings.Replace(s, "{{USAGE}}",
				"  -s size      : size of lru (default: 1000000)\n", -1)
		},
	}
	server.Main(
		func(w http.ResponseWriter, r *http.Request) {
			key := strings.Split(r.URL.Path, "/")[1]
			switch r.Method {
			case "GET":
				if val, ok := cache.Get(key); !ok {
					w.WriteHeader(404)
				} else {
					w.Write([]byte(val.(string)))
				}
			case "DELETE":
				cache.Remove(key)
			case "PUT", "POST":
				if data, err := ioutil.ReadAll(r.Body); err != nil {
					w.WriteHeader(500)
				} else {
					cache.Add(key, string(data))
				}
			default:
				w.WriteHeader(405)
			}
		}, opts,
	)
}
