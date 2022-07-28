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

func ShowFriend(friend apiengine.Friend) {
	printField("Username", friend.Username, 37)
	printField("All time", rawTimeToHumanReadable(friend.CodingTime.AllTime/60), 37)
	printField("Past week", rawTimeToHumanReadable(friend.CodingTime.PastWeek/60), 37)
	printField("Past month", rawTimeToHumanReadable(friend.CodingTime.PastMonth/60), 37)
}
