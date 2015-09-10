package main

import (
	"fmt"
	"github.com/jordanorelli/steam"
	"io"
	"net/http"
	"os"
	"sort"
	"strconv"
	"text/tabwriter"
)

var commands map[string]command

func init() {
	commands = map[string]command{
		"api-list":           cmd_api_list,
		"user-friends":       cmd_user_friends,
		"user-id":            cmd_user_id,
		"user-details":       cmd_user_details,
		"dota-match-history": cmd_dota_match_history,
		"dota-match-details": cmd_dota_match_details,
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

var cmd_api_list = command{
	handler: func(c *steam.Client, args ...string) {
		dump(c.Get("ISteamWebAPIUtil", "GetSupportedAPIList", "v0001"))
	},
}

var cmd_user_friends = command{
	handler: func(c *steam.Client, args ...string) {
		if len(args) < 1 {
			bail(1, "please provide a user id")
		}
		userid, err := strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			bail(1, "bad user id: %s", err)
		}
		friends, err := c.GetFriendList(userid)
		if err != nil {
			bail(1, "%v", err)
		}
		w := tabwriter.NewWriter(os.Stdout, 0, 8, 0, '\t', 0)
		defer w.Flush()
		for _, friend := range friends {
			fmt.Fprintln(w, friend.Oneline())
		}
	},
}

var cmd_user_id = command{
	handler: func(c *steam.Client, args ...string) {
		userid, err := c.ResolveVanityUrl(args[0])
		if err != nil {
			bail(1, err.Error())
		}
		fmt.Println(userid)
	},
}

var cmd_user_details = command{
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
		players, err := c.GetPlayerSummaries(ids...)
		if err != nil {
			bail(1, "%v", err)
		}
		w := tabwriter.NewWriter(os.Stdout, 0, 8, 0, '\t', 0)
		defer w.Flush()
		for _, player := range players {
			fmt.Fprintln(w, player.Oneline())
		}
	},
}

var cmd_dota_match_history = command{
	handler: func(c *steam.Client, args ...string) {
		matches, err := c.DotaMatchHistory(0, 0)
		if err != nil {
			bail(1, "%v", err)
		}
		w := tabwriter.NewWriter(os.Stdout, 0, 8, 0, '\t', 0)
		defer w.Flush()
		for _, match := range matches {
			fmt.Fprintln(w, match.Oneline())
		}
	},
}

var cmd_dota_match_details = command{
	handler: func(c *steam.Client, args ...string) {
		if len(args) != 1 {
			bail(1, "please provide exactly one match id")
		}
		id, err := strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			bail(1, "bad match id: %s", err)
		}
		details, err := c.DotaMatchDetails(id)
		if err != nil {
			bail(1, "%v", err)
		}
		details.Display(os.Stdout)
	},
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
