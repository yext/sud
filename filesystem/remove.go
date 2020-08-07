package filesystem

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/afero"
	"github.com/yext/sud/json"
)

func Remove(fileName string, dirs []string, path string, values []string, fs afero.Fs) {
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
				removeFile(existingFile, fs)
			} else {
				removeWithinFile(existingFile, path, values, fs)
			}
		}
	}
}

func removeWithinFile(existingFile string, path string, values []string, fs afero.Fs) {
	jsonData, err := afero.ReadFile(fs, existingFile)
	if err != nil {
		fmt.Printf("unable to read file "+existingFile+" %v\n", err)
		return
	}
	if jsonData == nil || string(jsonData) == "" {
		fmt.Println("empty file " + existingFile)
		return
	}

	movedJson := jsonData
	if values != nil {
		for _, value := range values {
			fmt.Println("removing path " + path + " with value " + value + " in " + existingFile)
			movedJson, err = json.RemoveByValue(path, value, movedJson)
			if err != nil {
				fmt.Printf("unable to remove "+path+" with "+value+" in "+existingFile+" %v\n", err)
				return
			}
		}
	} else {
		fmt.Println("removing path " + path + " in " + existingFile)
		movedJson, err = json.Remove(path, jsonData)
		if err != nil {
			fmt.Printf("unable to remove "+path+" in "+existingFile+" %v\n", err)
			return
		}
	}
	err = afero.WriteFile(fs, existingFile, movedJson, 0644)
	if err != nil {
		fmt.Printf("unable to write file "+existingFile+" %v\n", err)
		return
	}
}

func removeFile(existingFile string, fs afero.Fs) {
	fmt.Println("removing file " + existingFile)
	err := fs.Remove(existingFile)
	if err != nil {
		fmt.Printf("unable to remove "+existingFile+" %v", err)
		return
	}
}
