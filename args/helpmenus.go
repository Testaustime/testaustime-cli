package args

import "fmt"

type SubCommand struct {
	Name string
	Info string
}
type Command struct {
	Name        string
	Info        string
	SubCommands []SubCommand
}

func AccountUsage() {
	fmt.Print(
		commandUsage("account"),
		"\n",
		flags(),
		"\n",
		formatSubCommands(&Commands[0].SubCommands),
	)
}

func StatisticsUsage() {
	fmt.Print(
		commandUsage("statistics"),
		"\n",
		flags(),
		"\n",
		formatSubCommands(&Commands[1].SubCommands),
	)
}

func FriendsUsage() {
	fmt.Print(
		commandUsage("statistics"),
		"\n",
		flags(),
		"\n",
		formatSubCommands(&Commands[2].SubCommands),
	)
}

func Usage() {
	fmt.Print(
		header("<command>"),
		"\n",
		fmt.Sprintf("%s:\n", coloredString("commands", 33)),
		formatCommands(&Commands),
	)
}

func formatCommands(c *[]Command) (result string) {
	for _, i := range *c {
		r := i.Name
		result += fmt.Sprintf("  %s: %s\n", coloredString(r, 37), i.Info)
	}

	return result + "\n"
}

func formatSubCommands(c *[]SubCommand) (result string) {
	if len(*c) == 0 {
		return result
	}

	result += fmt.Sprintf("%s:\n", coloredString("subcommands", 33))
	for _, i := range *c {
		r := i.Name
		result += fmt.Sprintf("  %s: %s\n", coloredString(r, 37), i.Info)
	}

	return result + "\n"
}

func commandUsage(command string) string {
	return fmt.Sprintln(
		fmt.Sprint(coloredString("usage", 32), ":"),
		"./testaustime-cli [flags]",
		command,
		"[subcommand]",
	)
}

func header(command string) string {
	return fmt.Sprint(
		commandUsage(command),
		"\n",
		flags(),
	)
}

func flags() string {
	return fmt.Sprint(
		fmt.Sprintf("%s:\n", coloredString("flags", 33)),
		coloredString("  -c file", 37), ": read configuration from file \n",
		coloredString("  -no-colors", 37), ": don't include colors in output \n",
		coloredString("  -h, -help", 37), ": show help menu\n",
	)
}
