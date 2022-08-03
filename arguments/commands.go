package arguments

type SubCommand struct {
	Name        string
	Info        string
	SubCommands map[string]SubCommand
}

type Command struct {
	Name        string
	Info        string
	SubCommands map[string]SubCommand
}

var AccountCommand = Command{
	Name: "account",
	Info: "manage accounts",
	SubCommands: map[string]SubCommand{
		"register": {
			Name: "register",
			Info: "create a new account to testaustime",
		},
		"login": {
			Name: "login",
			Info: "login to your account with username and password",
			SubCommands: map[string]SubCommand{
				"<username>": {
					Name: "<username>",
					Info: "Specify username for login",
				},
			},
		},
		"token": {
			Name: "token",
			Info: "shows your authtoken",
		},
		"newToken": {
			Name: "newToken",
			Info: "regenerates your authentication token",
		},
		"newFriendcode": {
			Name: "newCode",
			Info: "regenerates your friend code",
		},
		"changePassword": {
			Name: "newPassword",
			Info: "change your password",
		},
	},
}

var StatisticsCommand = Command{
	Name: "statistics",
	Info: "get coding statistics",
	SubCommands: map[string]SubCommand{
		"top": topSubCommand,
	},
}

var LeaderboardCommand = Command{
	Name: "leaderboards",
	Info: "get leaderboards data",
	SubCommands: map[string]SubCommand{
		"<name>": {
			Name:        "<name>",
			Info:        "show a leaderboard",
			SubCommands: map[string]SubCommand{},
		},
		"create": {
			Name:        "create",
			Info:        "create a leaderboard",
			SubCommands: map[string]SubCommand{},
		},
		"delete": {
			Name:        "delete",
			Info:        "delete a leaderboard",
			SubCommands: map[string]SubCommand{},
		},
		"join": {
			Name:        "joinboard",
			Info:        "join a leaderboard",
			SubCommands: map[string]SubCommand{},
		},
		"leave": {
			Name:        "leave",
			Info:        "leave a leaderboard",
			SubCommands: map[string]SubCommand{},
		},
		"regenerate": {
			Name:        "regenerate",
			Info:        "regenerate a leaderboard invite token",
			SubCommands: map[string]SubCommand{},
		},
		"kick": {
			Name: "kick",
			Info: "Kick a member from a leaderboard",
			SubCommands: map[string]SubCommand{
				"<name>": {
					Name: "<name>",
					Info: "Leaderboard name",
					SubCommands: map[string]SubCommand{
						"<username>": {
							Name:        "<username>",
							Info:        "Member name",
							SubCommands: map[string]SubCommand{},
						},
					},
				},
			},
		},
	},
}

var FriendsCommand = Command{
	Name: "friends",
	Info: "get friends' coding statistics",
	SubCommands: map[string]SubCommand{
		"pastMonth": {
			Name: "pastMonth",
			Info: "show friends' coding time during past month",
		},
		"pastWeek": {
			Name: "pastWeek",
			Info: "show friends' coding time during past week",
		},
		"add": {
			Name: "addFriend",
			Info: "Add a new friend",
		},
		"remove": {
			Name: "removeFriend",
			Info: "Remove a friend",
		},
	},
}

var UserCommand = Command{
	Name: "getuser",
	Info: "get specific friend's statistics",
	SubCommands: map[string]SubCommand{
		"<user>": {
			Name: "<user>",
			Info: "show data for specific user",
			SubCommands: map[string]SubCommand{
				"top": topSubCommand,
			},
		},
	},
}

var HelpCommand = Command{
	Name: "helpcmd",
	Info: "show help menus for specific commands",
	SubCommands: map[string]SubCommand{
		"<command>": {
			Name:        "<command>",
			Info:        "show help menu for given command",
			SubCommands: map[string]SubCommand{},
		},
	},
}

var Commands = []Command{
	HelpCommand,
	AccountCommand,
	StatisticsCommand,
	LeaderboardCommand,
	FriendsCommand,
	UserCommand,
}

var topSubCommand = SubCommand{
	Name:        "top",
	Info:        "show top languages and projects",
	SubCommands: topSubCommands,
}

var topSubCommands = map[string]SubCommand{
	"pastWeek": {
		Name:        "pastWeek",
		Info:        "show past week's top statistics",
		SubCommands: nil,
	},
	"pastMonth": {
		Name:        "pastMonth",
		Info:        "show past month's top statistics",
		SubCommands: nil,
	},
}
