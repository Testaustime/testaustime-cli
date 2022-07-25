package args

var AccountCommand = Command{
	Name: "account",
	Info: "show account information",
	SubCommands: []SubCommand{
		{
			Name: "token",
			Info: "shows your authtoken",
		},
		{
			Name: "newToken",
			Info: "regenerates your authorization token",
		},
		{
			Name: "newFriendcode",
			Info: "regenerates your friend code",
		},
	},
}

var StatisticsCommand = Command{
	Name: "statistics",
	Info: "show coding statistics",
	SubCommands: []SubCommand{
		{
			Name: "latest",
			Info: "show latest languages and projects",
		},
	},
}

var FriendsCommand = Command{
	Name: "friends",
	Info: "show friends' statistics",
	SubCommands: []SubCommand{
		{
			Name: "pastMonth",
			Info: "show friends' coding time during past month",
		},
		{
			Name: "pastWeek",
			Info: "show friends' coding time during past week",
		},
	},
}

var UserCommand = Command{
	Name:        "user",
	Info:        "show specific friend's statistics",
	SubCommands: []SubCommand{},
}

var Commands = []Command{
	AccountCommand,
	StatisticsCommand,
	FriendsCommand,
	UserCommand,
}
