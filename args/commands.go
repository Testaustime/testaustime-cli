package args

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
	Info: "show account information",
	SubCommands: map[string]SubCommand{
		"register": {
			Name: "register",
			Info: "create a new account to testaustime",
		},
		"login": {
			Name: "login",
			Info: "login to your account with username and password",
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
	Info: "show coding statistics",
	SubCommands: map[string]SubCommand{
		"top": {
			Name:        "top",
			Info:        "show top languages and projects",
			SubCommands: blaa,
		},
	},
}

var FriendsCommand = Command{
	Name: "friends",
	Info: "show friends' statistics",
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
	Info: "show specific friend's statistics",
	SubCommands: map[string]SubCommand{
		"<user>": {
			Name: "<user>",
			Info: "show data for specific user",
			SubCommands: map[string]SubCommand{
				"top": {
					Name:        "top",
					Info:        "show top statistics",
					SubCommands: blaa,
				},
			},
		},
	},
}

var Commands = []Command{
	AccountCommand,
	StatisticsCommand,
	FriendsCommand,
	UserCommand,
}

var blaa = map[string]SubCommand{
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
