package steam

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// a steam API client, not tied to any particular game
type Client struct {
	key string
}

func NewClient(key string) *Client {
	return &Client{key: key}
}

func (c *Client) Get(iface, method, version string) (*http.Response, error) {
	url := fmt.Sprintf("https://api.steampowered.com/%s/%s/%s/?key=%s", iface, method, version, c.key)
	return http.Get(url)
}

func (c *Client) GetFriendList(userid uint64) (*http.Response, error) {
	url := fmt.Sprintf("https://api.steampowered.com/ISteamUser/GetFriendList/v1/?key=%s&steamid=%d", c.key, userid)
	return http.Get(url)
}

func (c *Client) ResolveVanityUrl(vanity string) (*http.Response, error) {
	url := fmt.Sprintf("https://api.steampowered.com/ISteamUser/ResolveVanityURL/v0001/?key=%s&vanityurl=%s", c.key, vanity)
	return http.Get(url)
}

func (c *Client) GetPlayerSummaries(steamids ...uint64) (*http.Response, error) {
	if len(steamids) > 100 {
		return nil, fmt.Errorf("GetPlayerSummaries accepts a max of 100 ids, saw %d", len(steamids))
	}
	ids_s := make([]string, len(steamids))
	for i := range steamids {
		ids_s[i] = strconv.FormatUint(steamids[i], 10)
	}
	ids := strings.Join(ids_s, ",")
	url := fmt.Sprintf("https://api.steampowered.com/ISteamUser/GetPlayerSummaries/v0002/?key=%s&steamids=%s", c.key, ids)
	return http.Get(url)
}

/*
       "name": "ISteamUser",
       "methods": [
           {
               "name": "GetPlayerBans",
               "version": 1,
               "httpmethod": "GET",
               "parameters": [
                   {
                       "name": "key",
                       "type": "string",
                       "optional": false,
                       "description": "access key"
                   },
                   {
                       "name": "steamids",
                       "type": "string",
                       "optional": false,
                       "description": "Comma-delimited list of SteamIDs"
                   }
               ]

           },
           {
               "name": "GetPlayerSummaries",
               "version": 1,
               "httpmethod": "GET",
               "parameters": [
                   {
                       "name": "key",
                       "type": "string",
                       "optional": false,
                       "description": "access key"
                   },
                   {
                       "name": "steamids",
                       "type": "string",
                       "optional": false,
                       "description": "Comma-delimited list of SteamIDs"
                   }
               ]

           },
           {
               "name": "GetPlayerSummaries",
               "version": 2,
               "httpmethod": "GET",
               "parameters": [
                   {
                       "name": "key",
                       "type": "string",
                       "optional": false,
                       "description": "access key"
                   },
                   {
                       "name": "steamids",
                       "type": "string",
                       "optional": false,
                       "description": "Comma-delimited list of SteamIDs (max: 100)"
                   }
               ]

           },
           {
               "name": "GetUserGroupList",
               "version": 1,
               "httpmethod": "GET",
               "parameters": [
                   {
                       "name": "key",
                       "type": "string",
                       "optional": false,
                       "description": "access key"
                   },
                   {
                       "name": "steamid",
                       "type": "uint64",
                       "optional": false,
                       "description": "SteamID of user"
                   }
               ]

           },
           {
               "name": "ResolveVanityURL",
               "version": 1,
               "httpmethod": "GET",
               "parameters": [
                   {
                       "name": "key",
                       "type": "string",
                       "optional": false,
                       "description": "access key"
                   },
                   {
                       "name": "vanityurl",
                       "type": "string",
                       "optional": false,
                       "description": "The vanity URL to get a SteamID for"
                   },
                   {
                       "name": "url_type",
                       "type": "int32",
                       "optional": true,
                       "description": "The type of vanity URL. 1 (default): Individual profile, 2: Group, 3: Official game group"
                   }
               ]

           }
       ]

   },
*/
