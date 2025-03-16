package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var ApiKey = "Set via CLI flags"

func GetFriends(steamId string) ([]SteamFriend, error) {
	url := fmt.Sprintf("https://api.steampowered.com/ISteamUser/GetFriendList/v1/?key=%s&steamid=%s&relationship=friend", ApiKey, steamId)

	resp, err := http.Get(url)

	if err != nil {
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
	url := fmt.Sprintf("https://api.steampowered.com/ISteamUser/GetPlayerSummaries/v0002/?key=%s&steamids=%s", ApiKey, steamId)

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
