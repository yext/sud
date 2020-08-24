package cmd

import (
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/yext/sud/filesystem"
	"github.com/yext/sud/github"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add FILE_PATH DIRs...",
	Short: "adds a file, or a path in a file, in the given list of repositories",
	Long: `Adds a file, or a path in a file, in the given repositories :

 - Adds the given file in all the given repositories. It adds only if the file is not already present.
 - Adds within a file if the path flag is provided.

The value provided should conform to the intended type.
For example,
 1. to add a number would be --value=304
 2. to add string would be --value=\"my-string\" (escape the quote in bash)
 3. to add an array would be --value="[1, 2]"

Instead of providing a value, you can also provide a file path of a file that has the desired value.

Examples:

Adding a File:

sud add foo.json
  --file bar.json
  ~/repos/ans150-front-end-overview

will add a file foo.json with the contents of the provided bar.json to ~/repos/ans150-front-end-overview.

Adding a Path:

sud add dependencies.json
  --path /productFeatureIds/-
  --values 304,305
  ~/repos/ans150-front-end-overview

will add product features 304 and 305 to ~repos/ans150-front-end-overview/dependencies.json:

{ "productFeatureIds": [1, 2] } => { "productFeatureIds": [1, 2, 304, 305] }`,
	Args:   cobra.MinimumNArgs(2),
	PreRun: loadFileToValues,
	Run: func(cmd *cobra.Command, args []string) {
		fs := afero.NewOsFs()
		fileName := args[0]
		repos := args[1:]

		dirs, err := github.CloneToDirs(repos)
		if err != nil {
			panic(err)
		}

		filesystem.Add(fileName, values, dirs, path, fs)

		if !push {
			return
		}
		err = github.Push(dirs)
		if err != nil {
			panic(err)
		}
	},
}

var values []string

func loadFileToValues(cmd *cobra.Command, args []string) {
	if value != "" {
		values = append(values, value)
	}
	if values == nil && file == "" {
		panic("either value flag or file flag should be set")
	}
	if values != nil && file != "" {
		panic("only one of value flag or file flag should be set")
	}
	if file != "" {
		valueBytes, err := afero.ReadFile(afero.NewOsFs(), file)
		if err != nil {
			panic("unable to read file " + file)
		}
		values = append(values, string(valueBytes))
	}
}

var file string

func init() {
	RootCmd.AddCommand(addCmd)

	addCmd.Flags().StringSliceVarP(
		&values,
		"values",
		"l",
		nil,
		"The desired list of values")

	addCmd.Flags().StringVarP(
		&value,
		"value",
		"v",
		"",
		"The desired value")

	addCmd.Flags().StringVarP(
		&file,
		"file",
		"f",
		"",
		"The file with the desired value")

	addCmd.Flags().StringVarP(
		&path,
		"path",
		"p",
		"",
		"The path to the json pointer in the file as per https://tools.ietf.org/html/rfc6901")
}
