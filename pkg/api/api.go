package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var Key = "Set via CLI flags"
var baseUrl = "https://api.steampowered.com/ISteamUser"

func GetFriends(steamId string) ([]SteamFriend, error) {
	url := baseUrl + fmt.Sprintf("/GetFriendList/v1/?key=%s&steamid=%s&relationship=friend", Key, steamId)

	resp, err := http.Get(url)

	if err != nil || resp.StatusCode != http.StatusOK {
		return []SteamFriend{}, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return []SteamFriend{}, err
	}

	var result SteamFriendsListResult
	err = json.Unmarshal(data, &result)

	if err != nil {
		return []SteamFriend{}, err
	}

	return result.FriendsList.Friends, nil
}

func GetUser(steamId string) (SteamUser, error) {
	url := baseUrl + fmt.Sprintf("/GetPlayerSummaries/v0002/?key=%s&steamids=%s", Key, steamId)

	resp, err := http.Get(url)

	if err != nil {
		return SteamUser{}, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return SteamUser{}, err
	}

	var result SteamUsersResult
	err = json.Unmarshal(data, &result)

	if err != nil || len(result.Response.Players) == 0 {
		return SteamUser{}, SteamUserNotFound{}
	}

	return result.Response.Players[0], nil
}
