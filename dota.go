package steam

import (
	"bytes"
	"fmt"
)

type DotaMatch struct {
	Id            uint64            `json:"match_id"`
	SeqNum        uint64            `json:"match_seq_num"`
	StartTime     uint64            `json:"start_time"`
	LobbyType     int               `json:"lobby_type"`
	RadiantTeamId int               `json:"radian_team_id"`
	DireTeamId    int               `json:"dire_team_id"`
	Players       []DotaMatchPlayer `json:"players"`
}

func (d DotaMatch) Oneline() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "%d\t%d\t%d\t%d\t%d\t%d\n", d.Id, d.SeqNum, d.StartTime, d.LobbyType, d.RadiantTeamId, d.DireTeamId)
	for _, player := range d.Players {
		fmt.Fprintf(&buf, "-\t-\t-\t%d\t%d\t%d\n", player.AccountId, player.PlayerSlot, player.HeroId)
	}
	return buf.String()
}

type DotaMatchPlayer struct {
	AccountId  uint64 `json:"account_id"`
	PlayerSlot int    `json:"player_slot"`
	HeroId     int    `json:"hero_id"`
}

/*
{
    "match_id": 1680503925,
    "match_seq_num": 1497692382,
    "start_time": 1438511639,
    "lobby_type": 0,
    "radiant_team_id": 0,
    "dire_team_id": 0,
    "players": [
        {
            "account_id": 4294967295,
            "player_slot": 0,
            "hero_id": 61
        },
        {
            "account_id": 4294967295,
            "player_slot": 1,
            "hero_id": 56
        },
        {
            "account_id": 4294967295,
            "player_slot": 2,
            "hero_id": 41
        },
        {
            "account_id": 264612072,
            "player_slot": 3,
            "hero_id": 48
        },
        {
            "account_id": 4294967295,
            "player_slot": 4,
            "hero_id": 95
        },
        {
            "account_id": 176784521,
            "player_slot": 128,
            "hero_id": 99
        },
        {
            "account_id": 4294967295,
            "player_slot": 129,
            "hero_id": 5
        },
        {
            "account_id": 4294967295,
            "player_slot": 130,
            "hero_id": 17
        },
        {
            "account_id": 62572906,
            "player_slot": 131,
            "hero_id": 46
        },
        {
            "account_id": 111126659,
            "player_slot": 132,
            "hero_id": 66
        }
    ]

},

*/
