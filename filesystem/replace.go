package filesystem

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/afero"
	"yext/m4/confcode/cmd/sud/json"
)

// Replace replaces the fileName with the value in the given directories.
// If the path is non-empty, then it replaces the path within the file.
func Replace(fileName string, value string, dirs []string, path string, fs afero.Fs) {
	for _, dir := range dirs {
		isDir, err := afero.IsDir(fs, dir)
		if !isDir || err != nil {
			continue
		}
		existingGlob := filepath.Join(dir, fileName)
		files, err := filepath.Glob(existingGlob)
		if err != nil {
			fmt.Printf("unable to glob "+existingGlob+" %v\n", err)
			continue
		}
		for _, existingFile := range files {
			exists, err := afero.Exists(fs, existingFile)
			if !exists || err != nil {
				continue
			}
			if path == "" {
				replaceFile(existingFile, value, fs)
			} else {
				replaceWithinFile(existingFile, path, value, fs)
			}
		}
	}
}

func replaceWithinFile(existingFile string, path string, value string, fs afero.Fs) {
	jsonData, err := afero.ReadFile(fs, existingFile)
	if err != nil {
		fmt.Printf("unable to read file "+existingFile+" %v\n", err)
		return
	}
	if jsonData == nil || string(jsonData) == "" {
		fmt.Println("empty file " + existingFile)
		return
	}

	fmt.Println("replacing path " + path + " with " + value + " in " + existingFile)
	replacedJson, err := json.Replace(path, value, jsonData)
	if err != nil {
		fmt.Printf("unable to replace "+path+" with value "+value+" in "+existingFile+" %v\n", err)
		return
	}
	err = afero.WriteFile(fs, existingFile, replacedJson, 0644)
	if err != nil {
		fmt.Printf("unable to write file "+existingFile+" %v\n", err)
		return
	}
}

func replaceFile(existingFile string, value string, fs afero.Fs) {
	fmt.Println("replacing " + existingFile + " with value " + value)
	err := afero.WriteFile(fs, existingFile, []byte(value), 0644)
	if err != nil {
		fmt.Printf("unable to replace "+existingFile+" with value "+value+" %v\n", err)
		return
	}
}
