package cmd

import (
	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "sud COMMAND FILE_PATH DIRs...",
	Short: "sud OR solution updater updates solutions by running the command on the file in the given repositories",
	Long: `sud OR solution updater updates solutions by running the command on the file in the given repositories.

Repositories are space separated list of directories or github repository urls.

A repository can be one of the following:
 1. A directory on your file system such as ~/repos/my-solution-dir
 2. A url to github repository such as https://github.com/YextHHChallenges/asbTest
 3. A wildcard such as https://github.com/YextHHChallenges/asbTest* that will be expanded to matched repositories

For example:

sud replace default/km/*.json --value "\"yext/atm\"" --path /primaryEntityType https://github.com/YextHHChallenges/asb* --push

will make the change in the matching files in all the matching repositories and push to github.

Note: Please consider using quotes around wildcard * to avoid bash expansion, and skip the quotes for bash expansion.
`,
}

func Execute() error {
	return RootCmd.Execute()
}

var push bool

func init() {
	RootCmd.PersistentFlags().BoolVar(
		&push,
		"push",
		false,
		"Whether to push the repo to github")
}
