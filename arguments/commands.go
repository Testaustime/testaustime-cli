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

var ExperimentalCommand = Command{
	Name: "experimental",
	Info: "experimental features",
	SubCommands: map[string]SubCommand{
		"summary": {
			Name: "summary",
			Info: "show summary information",
		},
		"settings": {
			Name: "account",
			Info: "edit account settings",
			SubCommands: map[string]SubCommand{
				"public": {
					Name:        "setpublic",
					Info:        "change your account's publicity state",
					SubCommands: trueOrFalse,
				},
			},
		},
	},
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
			Name: "<name>",
			Info: "show a leaderboard",
		},
		"create": {
			Name: "create",
			Info: "create a leaderboard",
		},
		"delete": {
			Name: "delete",
			Info: "delete a leaderboard",
		},
		"join": {
			Name: "joinboard",
			Info: "join a leaderboard",
		},
		"leave": {
			Name: "leave",
			Info: "leave a leaderboard",
		},
		"regenerate": {
			Name: "regenerate",
			Info: "regenerate a leaderboard invite token",
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
							Name: "<username>",
							Info: "Member name",
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
			Name: "<command>",
			Info: "show help menu for given command",
		},
	},
}

var trueOrFalse = map[string]SubCommand{
	"on": {
		Name: "on",
		Info: "switch option on",
	},
	"off": {
		Name: "off",
		Info: "switch option off",
	},
}

var Commands = []Command{
	HelpCommand,
	AccountCommand,
	StatisticsCommand,
	LeaderboardCommand,
	FriendsCommand,
	UserCommand,
	ExperimentalCommand,
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
