package main

import (
	"fmt"
	"github.com/jordanorelli/steam"
	"io"
	"net/http"
	"os"
	"sort"
	"strconv"
)

var commands map[string]command

func init() {
	commands = map[string]command{
		"api-list": command{
			handler: func(c *steam.Client, args ...string) {
				dump(c.Get("ISteamWebAPIUtil", "GetSupportedAPIList", "v0001"))
			},
		},
		"user-friends": command{
			handler: func(c *steam.Client, args ...string) {
				if len(args) < 1 {
					bail(1, "please provide a user id")
				}
				userid, err := strconv.ParseUint(args[0], 10, 64)
				if err != nil {
					bail(1, "bad user id: %s", err)
				}
				dump(c.GetFriendList(userid))
			},
		},
		"user-id": command{
			handler: func(c *steam.Client, args ...string) {
				dump(c.ResolveVanityUrl(args[0]))
			},
		},
		"user-details": command{
			handler: func(c *steam.Client, args ...string) {
				if len(args) < 1 {
					bail(1, "please provide a user id")
				}
				ids := make([]uint64, 0, len(args))
				for _, arg := range args {
					userid, err := strconv.ParseUint(arg, 10, 64)
					if err != nil {
						bail(1, "bad user id: %s", err)
					}
					ids = append(ids, userid)
				}
				dump(c.GetPlayerSummaries(ids...))
			},
		},
		"commands": command{
			handler: func(c *steam.Client, args ...string) {
				keys := make([]string, 0, len(commands))
				for name, _ := range commands {
					keys = append(keys, name)
				}
				sort.Strings(keys)
				for _, key := range keys {
					fmt.Println(key)
				}
			},
		},
	}
}

type command struct {
	handler func(*steam.Client, ...string)
}

func dump(r *http.Response, e error) {
	if e != nil {
		bail(1, e.Error())
	}
	io.Copy(os.Stdout, r.Body)
}
