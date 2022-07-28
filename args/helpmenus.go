package args

import (
	"fmt"

	"github.com/romeq/testaustime-cli/utils"
)

// secondaryColor declares what color should be used when highlighting
// specific things in help menu.
var secondaryColor int = 35

// lightColor declares what color should be used when something wants
// to be given less attention. By default it's white.
var lightColor = 37

// CommandUsage prints command usage for a specific command
func CommandUsage(command Command) {
	fmt.Print(
		formatUsage(command.Name, ""),
		"\n", flags(), "\n",
		formatSubCommands(&command.SubCommands),
	)
}

// SubCommandUsage prints command usage for a specific command
func SubCommandUsage(command Command, subcommand map[string]SubCommand) {
	fmt.Print(
		formatUsage(fmt.Sprintf("%s %s", command.Name, subcommand), "[options]"),
		"\n", flags(), "\n",
		formatSubCommands(&subcommand),
	)
}

// UserUsage prints usage for user command.
func UserUsage() {
	fmt.Print(formatUsage(UserCommand.Name, "<user>"), "\n", flags())
}

// Usage prints program's general usage.
func Usage() {
	fmt.Print(
		header("<command>"), "\n",
		fmt.Sprintf("%s \t \n", coloredString("commands", secondaryColor)),
		formatCommands(&Commands),
	)
}

// formatCommands returns every command specified in commands list
func formatCommands(commands *[]Command) (result string) {
	for _, i := range *commands {
		r := i.Name
		result += fmt.Sprintf("  %s \t %s\n", r, coloredString(i.Info, lightColor))
	}
	return result + "\n"
}

// formatSubCommands prints subcommands header and usage of command's every subcommand.
// If no subcommands are given, it will result in an empty string.
func formatSubCommands(c *map[string]SubCommand) (result string) {
	if len(*c) == 0 {
		return result
	}

	result += fmt.Sprintf("%s\n", coloredString("subcommands", secondaryColor))
	for _, i := range *c {
		r := i.Name
		result += fmt.Sprintf("  %s \t %s\n", r, coloredString(i.Info, lightColor))
		formatSubCommands(&i.SubCommands)
	}
	return result + "\n"
}

// formatUsage returns colored usage of program.
// if subcommand is an empty string, it will be replaced with "[subcommand]"
func formatUsage(command, subcommand string) string {
	return fmt.Sprintln(
		fmt.Sprint(coloredString("usage", secondaryColor)),
		"./testaustime",
		coloredString("[flags]", lightColor),
		coloredString(command, lightColor),
		coloredString(utils.StringOr(subcommand, "[subcommand]"), lightColor),
	)
}

// header returns a string including a colored header
func header(command string) string {
	return fmt.Sprint(
		formatUsage(command, ""),
		"\n",
		flags(),
	)
}

// flags returns all program flags and their usage.
func flags() string {
	return fmt.Sprint(
		fmt.Sprintf("%s\n", coloredString("flags", secondaryColor)),
		"  -c file \t ", coloredString("read configuration from file \n", lightColor),
		"  -no-colors \t ", coloredString("don't include colors in output \n", lightColor),
		"  -time  \t ", coloredString("measure time taken in request\n", lightColor),
		"  -h, -help  \t ", coloredString("show help menu\n", lightColor),
	)
}
