package args

type SubCommand struct {
	Name string
	Info string
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
			Info: "regenerates your authorization token",
		},
		"newFriendcode": {
			Name: "newFriendcode",
			Info: "regenerates your friend code",
		},
		"changePassword": {
			Name: "changePassword",
			Info: "change your password",
		},
	},
}

var StatisticsCommand = Command{
	Name: "statistics",
	Info: "show coding statistics",
	SubCommands: map[string]SubCommand{
		"latest": {
			Name: "latest",
			Info: "show latest languages and projects",
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
			Name: "add",
			Info: "Add a new friend",
		},
		"remove": {
			Name: "remove",
			Info: "Remove a friend",
		},
	},
}

var UserCommand = Command{
	Name:        "user",
	Info:        "show specific friend's statistics",
	SubCommands: map[string]SubCommand{},
}

var Commands = []Command{
	AccountCommand,
	StatisticsCommand,
	FriendsCommand,
	UserCommand,
}
