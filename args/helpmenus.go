package args

import (
	"fmt"

	"github.com/romeq/testaustime-cli/utils"
)

func CommandUsage(command Command) {
	fmt.Print(
		formatUsage(command.Name, ""),
		"\n",
		flags(),
		"\n",
		formatSubCommands(&command.SubCommands),
	)
}

func UserUsage() {
	fmt.Print(
		formatUsage(UserCommand.Name, "<user>"),
		"\n",
		flags(),
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

func formatSubCommands(c *map[string]SubCommand) (result string) {
	if len(*c) == 0 {
		return result
	}

	result += fmt.Sprintf("%s:\n", coloredString("subcommands", 33))
	for r, i := range *c {
		result += fmt.Sprintf("  %s: %s\n", coloredString(r, 37), i.Info)
	}

	return result + "\n"
}

func formatUsage(command, subcommand string) string {
	return fmt.Sprintln(
		fmt.Sprint(coloredString("usage", 32), ":"),
		"./testaustime-cli [flags]",
		command,
		utils.StringOr(subcommand, "[subcommand]"),
	)
}

func header(command string) string {
	return fmt.Sprint(
		formatUsage(command, ""),
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
