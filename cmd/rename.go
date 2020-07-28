package cmd

import (
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"yext/m4/confcode/cmd/sud/filesystem"
	"yext/m4/confcode/cmd/sud/github"
)

// renameCmd represents the rename command
var renameCmd = &cobra.Command{
	Use:   "rename FILE_PATH DIRs...",
	Short: "renames a file, or a path in a file, in the given list of repositories",
	Long: `Renames a file, or a path in a file, in the given repositories:

Renames the given file in all the given repositories, if the file is present.
Renames within a file if the path flag is provided.

Examples:

Renaming a File:

sud rename dependencies.json
  --value foo.json
  ~/repos/ans*

will rename dependencies.json to foo.json in all ~/repos/ans* directories.

Renaming a Path:

sud rename km/*/*.json
  --path /apiName
  --value "\$id"
  ~/repos/ans*

will replace "apiName" with "$id" in km/*/*.json files in all ~/repos/ans* directories:

{ "apiName": "123" } => { "$id": "123" }`,
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fs := afero.NewOsFs()
		fileName := args[0]
		repos := args[1:]

		dirs, err := github.CloneToDirs(repos)
		if err != nil {
			panic(err)
		}

		filesystem.Rename(fileName, value, dirs, path, fs)

		if !push {
			return
		}
		err = github.Push(dirs)
		if err != nil {
			panic(err)
		}
	},
}

var value string
var path string

func init() {
	RootCmd.AddCommand(renameCmd)

	renameCmd.Flags().StringVarP(
		&value,
		"value",
		"v",
		"",
		"The desired value")
	err := cobra.MarkFlagRequired(renameCmd.Flags(), "value")
	if err != nil {
		panic("unable to mark required for flag named value")
	}

	renameCmd.Flags().StringVarP(
		&path,
		"path",
		"p",
		"",
		"The path to the json pointer in the file as per https://tools.ietf.org/html/rfc6901")
}
