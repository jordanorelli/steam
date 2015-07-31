package main

import (
	"github.com/jordanorelli/steam"
	"io"
	"net/http"
	"os"
)

var commands = map[string]command{
	"api-list": command{
		handler: func(c *steam.Client) {
			dump(c.Get("ISteamWebAPIUtil", "GetSupportedAPIList", "v0001"))
		},
	},
}

type command struct {
	handler func(*steam.Client)
}

func dump(r *http.Response, e error) {
	if e != nil {
		bail(1, e.Error())
	}
	io.Copy(os.Stdout, r.Body)
}
