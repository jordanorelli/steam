package steam

import (
	"encoding/json"
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

func (c *Client) GetFriendList(userid uint64) ([]PlayerFriend, error) {
	url := fmt.Sprintf("https://api.steampowered.com/ISteamUser/GetFriendList/v1/?key=%s&steamid=%d", c.key, userid)
	res, err := http.Get(url)
	if err != nil {
		return nil, errorf(err, "unable to get friend list")
	}
	var response struct {
		V struct {
			Friends []PlayerFriend `json:"friends"`
		} `json:"friendslist"`
	}
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, errorf(err, "unable to parse friends list response")
	}
	return response.V.Friends, nil
}

func (c *Client) ResolveVanityUrl(vanity string) (uint64, error) {
	var v struct {
		V struct {
			Id      uint64 `json:"steamid,string"`
			Success int    `json:"success"`
		} `json:"response"`
	}
	url := fmt.Sprintf("https://api.steampowered.com/ISteamUser/ResolveVanityURL/v0001/?key=%s&vanityurl=%s", c.key, vanity)
	res, err := http.Get(url)
	if err != nil {
		return 0, errorf(err, "unable to resolve vanity url")
	}
	if err := json.NewDecoder(res.Body).Decode(&v); err != nil {
		return 0, errorf(err, "unable to decode vanity url response")
	}
	if v.V.Success != 1 {
		return 0, errorf(err, "resolving vanity url returned non-1 status")
	}
	return v.V.Id, nil
}

func (c *Client) GetPlayerSummaries(steamids ...uint64) ([]PlayerSummary, error) {
	if len(steamids) > 100 {
		return nil, errorf(nil, "GetPlayerSummaries accepts a max of 100 ids, saw %d", len(steamids))
	}
	ids_s := make([]string, len(steamids))
	for i := range steamids {
		ids_s[i] = strconv.FormatUint(steamids[i], 10)
	}
	ids := strings.Join(ids_s, ",")
	url := fmt.Sprintf("https://api.steampowered.com/ISteamUser/GetPlayerSummaries/v0002/?key=%s&steamids=%s", c.key, ids)
	res, err := http.Get(url)
	if err != nil {
		return nil, errorf(err, "unable to call GetPlayerSummaries API")
	}
	var response struct {
		V struct {
			Players []PlayerSummary `json:"players"`
		} `json:"response"`
	}
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, errorf(err, "unable to parse GetPlayerSummaries response")
	}
	return response.V.Players, nil
}

func (c *Client) DotaMatchSequence(lastId uint64, n int) ([]DotaMatch, error) {
	// http://api.steampowered.com/IDOTA2Match_<ID>/GetMatchHistoryBySequenceNum/v1
	url := fmt.Sprintf("https://api.steampowered.com/IDOTA2Match_570/GetMatchHistoryBySequenceNum/v1/?key=%s", c.key)
	if lastId > 0 {
		url = fmt.Sprintf("%s&start_at_match_seq_num=%d", url, lastId)
	}
	if n > 0 {
		url = fmt.Sprintf("%s&matches_requested=%d", url, n)
	}
	fmt.Println(url)
	var response struct {
		V struct {
			Status     int         `json:"status"`
			NumResults int         `json:"num_results"`
			Total      int         `json:"total_results"`
			Remaining  int         `json:"results_remaining"`
			Matches    []DotaMatch `json:"matches"`
		} `json:"result"`
	}
	res, err := http.Get(url)
	if err != nil {
		return nil, errorf(err, "unable to get match history")
	}
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, errorf(err, "unable to parse match history response")
	}
	return response.V.Matches, nil
}

func (c *Client) DotaMatchHistory(lastId uint64, n int) ([]DotaMatch, error) {
	url := fmt.Sprintf("https://api.steampowered.com/IDOTA2Match_570/GetMatchHistory/v0001/?key=%s", c.key)
	if lastId > 0 {
		url = fmt.Sprintf("%s&last_match_id=%d", url, lastId)
	}
	if n > 0 {
		url = fmt.Sprintf("%s&matches_requested=%d", url, n)
	}
	fmt.Println(url)
	var response struct {
		V struct {
			Status     int         `json:"status"`
			NumResults int         `json:"num_results"`
			Total      int         `json:"total_results"`
			Remaining  int         `json:"results_remaining"`
			Matches    []DotaMatch `json:"matches"`
		} `json:"result"`
	}
	res, err := http.Get(url)
	if err != nil {
		return nil, errorf(err, "unable to get match history")
	}
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, errorf(err, "unable to parse match history response")
	}
	return response.V.Matches, nil
}

func (c *Client) DotaMatchDetails(id uint64) (*DotaMatchDetails, error) {
	url := fmt.Sprintf("https://api.steampowered.com/IDOTA2Match_570/GetMatchDetails/v0001/?key=%s&match_id=%d", c.key, id)
	res, err := http.Get(url)
	if err != nil {
		return nil, errorf(err, "unable to get match details")
	}
	var result struct {
		V DotaMatchDetails `json:"result"`
	}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, errorf(err, "unable to parse match details")
	}
	return &result.V, nil
}
