package steam

import (
	"fmt"
)

type PlayerSummary struct {
	SteamId        uint64 `json:"steamid,string"`
	Visibility     int    `json:"communityvisibilitystate"`
	ProfileState   int    `json:"profilestate"`
	PersonaName    string `json:"personaname"`
	LastLogOff     int64  `json:"lastlogoff"`
	ProfileUrl     string `json:"profileurl"`
	Avatar         string `json:"avatar"`
	AvatarMedium   string `json:"avatarmedium"`
	AvatarFull     string `json:"avatarfull"`
	PersonaState   int    `json:"personastate"`
	LocCountryCode string `json:"loccountrycode"`
	LocStateCode   string `json:"locstatecode"`
	LocCityID      int    `json:"loccityid"`
}

func (p PlayerSummary) Oneline() string {
	return fmt.Sprintf("%d\t%s\t%s", p.SteamId, p.PersonaName, p.ProfileUrl)
}

type PlayerFriend struct {
	SteamId      uint64 `json:"steamid,string"`
	Relationship string `json:"relationship"`
	FriendSince  int    `json:"friend_since"`
}

func (p PlayerFriend) Oneline() string {
	return fmt.Sprintf("%d\t%s\t%d", p.SteamId, p.Relationship, p.FriendSince)
}
