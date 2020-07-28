package cmd

import (
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"yext/m4/confcode/cmd/sud/filesystem"
	"yext/m4/confcode/cmd/sud/github"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove FILE_PATH DIRs...",
	Short: "removes a file, or a path in a file, in the given list of repositories",
	Long: `Removes a file, or a path in a file, in the given repositories :

Removes the given file in all the given repositories, if the file is present.
Removes within a file if the path flag is provided.
If value is provided along with the path, then removes only the provided value, if present at the path. (Currently, only removing a number from an array is supported.)

Examples:

Remove a File:

sud remove dependencies.json
  ~/repos/ans150-front-end-overview
  ~/repos/kg*

Removes dependencies.json from all of the specified directories.

Removing a Value in a File:

sud remove dependencies.json
  --path /productFeatureIds
  --value 304
  ~/repos/ans150-front-end-overview
  ~/repos/kg*

will remove product feature 304 from dependencies.json in all the specified directories:

{ "productFeatureIds": [1, 2, 304] } => { "productFeatureIds": [1 ,2] }`,
	Args:   cobra.MinimumNArgs(2),
	PreRun: mergeFlagsIntoValues,
	Run: func(cmd *cobra.Command, args []string) {
		fs := afero.NewOsFs()
		fileName := args[0]
		repos := args[1:]

		dirs, err := github.CloneToDirs(repos)
		if err != nil {
			panic(err)
		}

		filesystem.Remove(fileName, dirs, path, values, fs)

		if !push {
			return
		}
		err = github.Push(dirs)
		if err != nil {
			panic(err)
		}
	},
}

func mergeFlagsIntoValues(cmd *cobra.Command, args []string) {
	if value != "" {
		values = append(values, value)
	}
	if file != "" {
		valueBytes, err := afero.ReadFile(afero.NewOsFs(), file)
		if err != nil {
			panic("unable to read file " + file)
		}
		values = append(values, string(valueBytes))
	}
}

func init() {
	RootCmd.AddCommand(removeCmd)

	removeCmd.Flags().StringVarP(
		&path,
		"path",
		"p",
		"",
		"The path to the json pointer in the file as per https://tools.ietf.org/html/rfc6901")

	removeCmd.Flags().StringVarP(
		&value,
		"value",
		"v",
		"",
		"The desired value")

	removeCmd.Flags().StringSliceVarP(
		&values,
		"values",
		"l",
		nil,
		"The desired list of values")

	removeCmd.Flags().StringVarP(
		&file,
		"file",
		"f",
		"",
		"The file with the desired value")
}
