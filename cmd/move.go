package cmd

import (
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"yext/m4/confcode/cmd/sud/filesystem"
	"yext/m4/confcode/cmd/sud/github"
)

// moveCmd represents the move command
var moveCmd = &cobra.Command{
	Use:   "move FILE_PATH DIRs...",
	Short: "moves a file, or a path in a file, in the given list of repositories",
	Long: `Moves a file, or a path in a file, in the given repositories :

Moves the given file in all the given repositories, if the file is present.
Moves within a file if the path flag is provided.

Examples:

Moving a File:

sud move km/entity-type/atm.json
  --value km/entity-type-extension/atm.json
  ~/repos/ans150-front-end-overview
  ~/repos/kg*

will move km/entity-type/atm.json to km/entity-type-extension/atm.json in the ans150-front-end-overview repo as well as any kg* repos.

Moving within a File:

sud move
  --path /a/b/c
  --value /a/d
  ~/repos/ans150-front-end-overview

will move the JSON element specified by the JSON pointer /a/b/c to /a/d in the provided directories:

{
  "a": {
    "b": {
	  ...
	  c: {...}
	}
  }
}

becomes:

{
  "a": {
	"b": {
	  ...
	},
	"d:" {...}
  }
}`,
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fs := afero.NewOsFs()
		fileName := args[0]
		repos := args[1:]

		dirs, err := github.CloneToDirs(repos)
		if err != nil {
			panic(err)
		}

		filesystem.Move(fileName, value, dirs, path, fs)

		if !push {
			return
		}
		err = github.Push(dirs)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(moveCmd)

	moveCmd.Flags().StringVarP(
		&value,
		"value",
		"v",
		"",
		"The desired value")
	err := cobra.MarkFlagRequired(moveCmd.Flags(), "value")
	if err != nil {
		panic("unable to mark required for flag named value")
	}

	moveCmd.Flags().StringVarP(
		&path,
		"path",
		"p",
		"",
		"The path to the json pointer in the file as per https://tools.ietf.org/html/rfc6901")
}
