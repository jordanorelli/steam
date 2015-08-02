package steam

import (
	"bytes"
	"fmt"
	"io"
	"strings"
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
	return strings.TrimSpace(buf.String())
}

type DotaMatchPlayer struct {
	AccountId  uint64 `json:"account_id"`
	PlayerSlot int    `json:"player_slot"`
	HeroId     int    `json:"hero_id"`
}

type DotaMatchDetails struct {
	RadiantWin         bool                     `json:"radian_win"`
	Duration           int                      `json:"duration"`
	StartTime          int                      `json:"start_time"`
	Id                 uint64                   `json:"match_id"`
	SeqNum             uint64                   `json:"match_seq_num"`
	TowerStatusRadiant int                      `json:"tower_status_radiant"`
	TowerStatusDire    int                      `json:"tower_status_dire"`
	Cluster            int                      `json:"cluster"`
	FirstBloodTime     int                      `json:"first_blood_time"`
	LobbyTime          int                      `json:"lobby_time"`
	HumanPlayers       int                      `json:"human_players"`
	LeagueId           int                      `json:"leagueid"`
	PositiveVotes      int                      `json:"positive_votes"`
	NegativeVotes      int                      `json:"negative_votes"`
	GameMode           int                      `json:"game_mode"`
	Engine             int                      `json:"engine"`
	Players            []DotaMatchPlayerDetails `json:"players"`
}

func (d DotaMatchDetails) Display(w io.Writer) {
	if d.RadiantWin {
		fmt.Fprintln(w, "Radiant Victory")
	} else {
		fmt.Fprintln(w, "Dire Victory")
	}
	fmt.Fprintf(w, "Duration: %d\n", d.Duration)
	fmt.Fprintf(w, "StartTime: %d\n", d.StartTime)
	fmt.Fprintf(w, "Id: %d\n", d.Id)
	fmt.Fprintf(w, "SeqNum: %d\n", d.SeqNum)
	fmt.Fprintf(w, "TowerStatusRadiant: %d\n", d.TowerStatusRadiant)
	fmt.Fprintf(w, "TowerStatusDire: %d\n", d.TowerStatusDire)
	fmt.Fprintf(w, "Cluster: %d\n", d.Cluster)
	fmt.Fprintf(w, "FirstBloodTime: %d\n", d.FirstBloodTime)
	fmt.Fprintf(w, "LobbyTime: %d\n", d.LobbyTime)
	fmt.Fprintf(w, "HumanPlayers: %d\n", d.HumanPlayers)
	fmt.Fprintf(w, "LeagueId: %d\n", d.LeagueId)
	fmt.Fprintf(w, "PositiveVotes: %d\n", d.PositiveVotes)
	fmt.Fprintf(w, "NegativeVotes: %d\n", d.NegativeVotes)
	fmt.Fprintf(w, "GameMode: %d\n", d.GameMode)
	fmt.Fprintf(w, "Engine: %d\n", d.Engine)
	for _, player := range d.Players {
		player.Display(w)
	}
}

type DotaMatchPlayerDetails struct {
	AccountId       uint64                `json:"account_id"`
	PlayerSlot      int                   `json:"player_slot"`
	HeroId          int                   `json:"hero_id"`
	Item0           int                   `json:"item_0"`
	Item1           int                   `json:"item_1"`
	Item2           int                   `json:"item_2"`
	Item3           int                   `json:"item_3"`
	Item4           int                   `json:"item_4"`
	Item5           int                   `json:"item_5"`
	Kills           int                   `json:"kills"`
	Deaths          int                   `json:"deaths"`
	Assists         int                   `json:"assists"`
	LeaverStatus    int                   `json:"leaver_status"`
	Gold            int                   `json:"gold"`
	LastHits        int                   `json:"last_hits"`
	Denies          int                   `json:"denies"`
	GoldPerMinute   int                   `json:"gold_per_min"`
	XPPerMinute     int                   `json:"xp_per_min"`
	GoldSpent       int                   `json:"gold_spent"`
	HeroDamage      int                   `json:"hero_damage"`
	TowerDamage     int                   `json:"tower_damage"`
	HeroHealing     int                   `json:"hero_healing"`
	Level           int                   `json:"level"`
	AbilityUpgrades []DotaAbilityUpgrades `json:"ability_upgrades"`
}

func (p DotaMatchPlayerDetails) Display(w io.Writer) {
	fmt.Fprintf(w, "AccountId: %d\n", p.AccountId)
	fmt.Fprintf(w, "PlayerSlot: %d\n", p.PlayerSlot)
	fmt.Fprintf(w, "HeroId: %d\n", p.HeroId)
	fmt.Fprintf(w, "Item0: %d\n", p.Item0)
	fmt.Fprintf(w, "Item1: %d\n", p.Item1)
	fmt.Fprintf(w, "Item2: %d\n", p.Item2)
	fmt.Fprintf(w, "Item3: %d\n", p.Item3)
	fmt.Fprintf(w, "Item4: %d\n", p.Item4)
	fmt.Fprintf(w, "Item5: %d\n", p.Item5)
	fmt.Fprintf(w, "Kills: %d\n", p.Kills)
	fmt.Fprintf(w, "Deaths: %d\n", p.Deaths)
	fmt.Fprintf(w, "Assists: %d\n", p.Assists)
	fmt.Fprintf(w, "LeaverStatus: %d\n", p.LeaverStatus)
	fmt.Fprintf(w, "Gold: %d\n", p.Gold)
	fmt.Fprintf(w, "LastHits: %d\n", p.LastHits)
	fmt.Fprintf(w, "Denies: %d\n", p.Denies)
	fmt.Fprintf(w, "GoldPerMinute: %d\n", p.GoldPerMinute)
	fmt.Fprintf(w, "XPPerMinute: %d\n", p.XPPerMinute)
	fmt.Fprintf(w, "GoldSpent: %d\n", p.GoldSpent)
	fmt.Fprintf(w, "HeroDamage: %d\n", p.HeroDamage)
	fmt.Fprintf(w, "TowerDamage: %d\n", p.TowerDamage)
	fmt.Fprintf(w, "HeroHealing: %d\n", p.HeroHealing)
	fmt.Fprintf(w, "Level: %d\n", p.Level)
}

type DotaAbilityUpgrades struct {
	Ability int `json:"ability"`
	Time    int `json:"time"`
	Level   int `json:"level"`
}
