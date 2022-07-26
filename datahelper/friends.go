package datahelper

import "github.com/romeq/testaustime-cli/apiengine"

func ShowFriends(friends []apiengine.FriendsCodingTime) {
	for _, friend := range friends {
		color := 37
		if friend.Username == "@me" {
			color = 32
		}
		printField(
			friend.Username,
			rawTimeToHumanReadable(float32(friend.CodingTime)/60.0),
			color,
		)
	}
}
