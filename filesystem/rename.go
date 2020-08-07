package filesystem

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/afero"
	"github.com/yext/sud/json"
)

// Rename renames the fileName to the value in the given directories.
// If path is non-empty, then it renames the path within the file.
func Rename(fileName string, value string, dirs []string, path string, fs afero.Fs) {
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
				renameFile(existingFile, filepath.Join(dir, value), fs)
			} else {
				renameWithinFile(existingFile, path, value, fs)
			}
		}
	}
}

func renameWithinFile(existingFile string, path string, value string, fs afero.Fs) {
	jsonData, err := afero.ReadFile(fs, existingFile)
	if err != nil {
		fmt.Printf("unable to read file "+existingFile+" %v\n", err)
		return
	}
	if jsonData == nil || string(jsonData) == "" {
		fmt.Println("empty file " + existingFile)
		return
	}

	fmt.Println("renaming path " + path + " to " + value + " in " + existingFile)
	renamedJson, err := json.Rename(path, value, jsonData)
	if err != nil {
		fmt.Printf("unable to rename "+path+" to value "+value+" in "+existingFile+" %v\n", err)
		return
	}
	err = afero.WriteFile(fs, existingFile, renamedJson, 0644)
	if err != nil {
		fmt.Printf("unable to write file "+existingFile+" %v\n", err)
		return
	}
}

func renameFile(existingFile string, toFile string, fs afero.Fs) {
	fmt.Println("renaming from " + existingFile + " to " + toFile)
	err := fs.Rename(existingFile, toFile)
	if err != nil {
		fmt.Printf("unable to rename "+existingFile+" to "+toFile+" %v\n", err)
		return
	}
}
