package filesystem

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/spf13/afero"
	"github.com/yext/sud/json"
)

// Move moves the fileName to the value in the given directories.
// If path is non-empty, then it moves the path within the file.
func Move(fileName string, value string, dirs []string, path string, fs afero.Fs) {
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
				moveFile(existingFile, dir, value, fs)
			} else {
				moveWithinFile(existingFile, path, value, fs)
			}
		}
	}
}

func moveWithinFile(existingFile string, path string, value string, fs afero.Fs) {
	jsonData, err := afero.ReadFile(fs, existingFile)
	if err != nil {
		fmt.Printf("unable to read file "+existingFile+" %v\n", err)
		return
	}
	if jsonData == nil || string(jsonData) == "" {
		fmt.Println("empty file " + existingFile)
		return
	}

	fmt.Println("moving path " + path + " to " + value + " in " + existingFile)
	movedJson, err := json.Move(path, value, jsonData)
	if err != nil {
		fmt.Printf("unable to move "+path+" to value "+value+" in "+existingFile+"%v\n", err)
		return
	}
	err = afero.WriteFile(fs, existingFile, movedJson, 0644)
	if err != nil {
		fmt.Printf("unable to write file "+existingFile+"%v\n", err)
		return
	}
}

func moveFile(existingFile string, dir string, value string, fs afero.Fs) {
	toFile := filepath.Join(dir, value)
	if strings.HasSuffix(value, "/") {
		err := fs.MkdirAll(toFile, 0777)
		if err != nil {
			fmt.Printf("unable to make dir "+toFile+" %v\n", err)
			return
		}
		toFile = filepath.Join(toFile, filepath.Base(existingFile))
	}
	fmt.Println("moving from " + existingFile + " to " + toFile)
	err := fs.Rename(existingFile, toFile)
	if err != nil {
		fmt.Printf("unable to move "+existingFile+" to "+toFile+" %v\n", err)
		return
	}
}
