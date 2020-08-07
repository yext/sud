package cmd

import (
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/yext/sud/filesystem"
	"github.com/yext/sud/github"
)

// replaceCmd represents the replace command
var replaceCmd = &cobra.Command{
	Use:   "replace FILE_PATH DIRs...",
	Short: "replaces a file, or a path in a file, in the given list of repositories",
	Long: `Replaces a file, or a path in a file, in the given repositories :

 - Replaces the given file in all the given repositories. It does nothing if the file is not present.
 - Replaces within a file if the path flag is provided.

The value provided should conform to the intended type.
For example,
 1. to replace a number would be --value=304
 2. to replace string would be --value=\"my-string\" (escape the quote in bash)
 3. to replace an array would be --value="[1, 2]"

Instead of providing a value you can also provide a file path of a file that has the desired value.

Examples:

Replacing a File:

sud replace pages/page-builder/template/page.json
  --file new-page.json
  ~/repos/pgs*

will replace pages/page-builder/template/page.json with the contents of new-page.json in all ~/repos/pgs* directories.

Replacing a Path:

sud replace km/*/hotel.json
  --path "/\$id"
  --value "\"my_hotel\""
  ~/repos/kg122-intro-to-fields

will replace the value at path /$id with "my_hotel" in ~/repos/kg122-intro-to-fields/km/entity-type-extension/hotel.json

sud replace km/*/hotel.json
  --path /enabled
  --value true
  ~/repos/kg122-intro-to-fields

will replace the value at path /enabled with true in ~/repos/kg122-intro-to-fields/km/entity-type-extension/hotel.json`,
	Args:   cobra.MinimumNArgs(2),
	PreRun: loadFileToValue,
	Run: func(cmd *cobra.Command, args []string) {
		fs := afero.NewOsFs()
		fileName := args[0]
		repos := args[1:]

		dirs, err := github.CloneToDirs(repos)
		if err != nil {
			panic(err)
		}

		filesystem.Replace(fileName, value, dirs, path, fs)

		if !push {
			return
		}
		err = github.Push(dirs)
		if err != nil {
			panic(err)
		}
	},
}

func loadFileToValue(cmd *cobra.Command, args []string) {
	if value == "" && file == "" {
		panic("either value flag or file flag should be set")
	}
	if value != "" && file != "" {
		panic("only one of value flag or file flag should be set")
	}
	if file != "" {
		valueBytes, err := afero.ReadFile(afero.NewOsFs(), file)
		if err != nil {
			panic("unable to read file " + file)
		}
		value = string(valueBytes)
	}
}

func init() {
	RootCmd.AddCommand(replaceCmd)

	replaceCmd.Flags().StringVarP(
		&value,
		"value",
		"v",
		"",
		"The desired value")

	replaceCmd.Flags().StringVarP(
		&file,
		"file",
		"f",
		"",
		"The file with the desired value")

	replaceCmd.Flags().StringVarP(
		&path,
		"path",
		"p",
		"",
		"The path to the json pointer in the file as per https://tools.ietf.org/html/rfc6901")
}
