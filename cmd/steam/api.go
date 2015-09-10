package main

import (
	"encoding/json"
	"fmt"
	"github.com/jordanorelli/steam"
	"os"
	"text/tabwriter"
)

type ApiList struct {
	Interfaces []Interface `json:"interfaces"`
}

type Interface struct {
	Name    string   `json:"name"`
	Methods []Method `json:"methods"`
}

type Method struct {
	Name       string  `json:"name"`
	Version    int     `json:"version"`
	HttpMethod string  `json:"httpmethod"`
	Params     []Param `json:"parameters"`
}

type Param struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Optional    bool   `json:"optional"`
	Description string `json:"description"`
}

var cmd_api_list = command{
	help: `
retrieves the list of currently supported api endpoints from steam and dumps
out the raw json
`,
	handler: func(c *steam.Client, args ...string) {
		dump(c.Get("ISteamWebAPIUtil", "GetSupportedAPIList", "v0001"))
	},
}

var cmd_api_interfaces = command{
	help: `
retrieves the list of currently supported api interfaces from steam
`,
	handler: func(c *steam.Client, args ...string) {
		res, err := c.Get("ISteamWebAPIUtil", "GetSupportedAPIList", "v0001")
		if err != nil {
			bail(1, "error: %s", err)
		}
		var response struct {
			ApiList ApiList `json:"apilist"`
		}
		if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
			bail(1, "error parsing response: %s", err)
		}
		for _, i := range response.ApiList.Interfaces {
			fmt.Println(i.Name)
		}
	},
}

var cmd_api_methods = command{
	help: `
`,
	handler: func(c *steam.Client, args ...string) {
		var filter map[string]bool
		if len(args) > 0 {
			filter = make(map[string]bool, len(args))
			for _, name := range args {
				filter[name] = true
			}
		}
		res, err := c.Get("ISteamWebAPIUtil", "GetSupportedAPIList", "v0001")
		if err != nil {
			bail(1, "error: %s", err)
		}
		var response struct {
			ApiList ApiList `json:"apilist"`
		}
		if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
			bail(1, "error parsing response: %s", err)
		}
		w := tabwriter.NewWriter(os.Stdout, 0, 8, 0, '\t', 0)
		defer w.Flush()
		for _, i := range response.ApiList.Interfaces {
			if filter != nil && !filter[i.Name] {
				continue
			}
			for _, m := range i.Methods {
				fmt.Fprintf(w, "%s\t%s\n", i.Name, m.Name)
			}
		}
	},
}
