package filesystem

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/spf13/afero"
	"github.com/yext/sud/json"
)

// Add adds the fileName to the value in the given directories.
// If path is non-empty, then it adds the path within the file.
func Add(fileName string, value []string, dirs []string, path string, fs afero.Fs) {
	for _, dir := range dirs {
		isDir, err := afero.IsDir(fs, dir)
		if !isDir || err != nil {
			continue
		}
		newFile := filepath.Join(dir, fileName)
		exists, err := afero.Exists(fs, newFile)
		if err != nil {
			fmt.Printf("unable to check if "+newFile+" exists %v\n", err)
			continue
		}
		if exists && path == "" {
			fmt.Println("file " + newFile + " already exists, please use replace command instead")
			continue
		}
		if path == "" {
			addFile(newFile, value, fs)
			continue
		}
		if exists {
			addWithinFile(newFile, path, value, fs)
			continue
		}
		files, err := filepath.Glob(newFile)
		if err != nil {
			fmt.Printf("unable to glob "+newFile+" %v\n", err)
			continue
		}
		for _, existingFile := range files {
			addWithinFile(existingFile, path, value, fs)
		}
	}
}

func addWithinFile(existingFile string, path string, values []string, fs afero.Fs) {
	jsonData, err := afero.ReadFile(fs, existingFile)
	if err != nil {
		fmt.Printf("unable to read file "+existingFile+" %v\n", err)
		return
	}
	if jsonData == nil || string(jsonData) == "" {
		fmt.Println("empty file " + existingFile)
		return
	}

	addedJson := jsonData
	for _, value := range values {
		fmt.Println("adding path " + path + " with " + value + " in " + existingFile)
		addedJson, err = json.Add(path, value, addedJson)
		if err != nil {
			fmt.Printf("unable to add "+path+" with value "+value+" in "+existingFile+"%v\n", err)
			return
		}
	}
	err = afero.WriteFile(fs, existingFile, addedJson, 0644)
	if err != nil {
		fmt.Printf("unable to write file "+existingFile+"%v\n", err)
		return
	}
}

func addFile(newFile string, values []string, fs afero.Fs) {
	value := strings.Join(values, "\n")
	fmt.Println("adding " + newFile + " with value " + value)
	err := afero.WriteFile(fs, newFile, []byte(value), 0644)
	if err != nil {
		fmt.Printf("unable to add "+newFile+" with value "+value+" %v\n", err)
		return
	}
}
