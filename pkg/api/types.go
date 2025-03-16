package api

type SteamUserNotFound struct{}

func (SteamUserNotFound) Error() string {
	return "Could not find steam user"
}

type SteamFriendsListResult struct {
	FriendsList SteamFriendsList `json:"friendslist"`
}

type SteamFriendsList struct {
	Friends []SteamFriend `json:"friends"`
}

type SteamFriend struct {
	SteamId string `json:"steamid"`
}

type SteamUsersResult struct {
	Response SteamUserList `json:"response"`
}

type SteamUserList struct {
	Players []SteamUser `json:"players"`
}

type SteamUser struct {
	SteamId string `json:"steamid"`
	Name    string `json:"personaname"`
	Url     string `json:"profileurl"`
}
