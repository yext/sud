package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestHelp(t *testing.T) {
	os.Args = []string{"sud", "--help"}
	err := Execute()
	if err != nil {
		t.Fatal(err)
	}

	os.Args = []string{"sud", "rename", "--help"}
	err = Execute()
	if err != nil {
		t.Fatal(err)
	}
}

func TestSolutionUpdater(t *testing.T) {
	// create tmp dir
	tmpDir, err := ioutil.TempDir("", "sud-test")
	fmt.Println("created " + tmpDir)
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		err := os.RemoveAll(tmpDir)
		if err != nil {
			t.Fatal(err)
		}
	}()

	// copy from testdata
	fs := afero.NewOsFs()
	err = copyAllFiles("testdata/source", tmpDir, fs, t)
	if err != nil {
		t.Fatal(err)
	}

	// rename
	// (note that the flag help has to be set to false to revert the help flag of true)
	os.Args = []string{"sud", "rename", "km/settings.json", "--value=km/newSettings.json", "--help=false", tmpDir + "/repo1", tmpDir + "/repo2"}
	err = Execute()
	if err != nil {
		t.Fatal(err)
	}

	// verify renamed
	err = assertSameFiles("testdata/renamed", tmpDir, fs, t)
	if err != nil {
		t.Fatal(err)
	}

	// rename within the file
	os.Args = []string{"sud", "rename", "km/newSettings.json", "--path=/primaryEntityType", "--value=p", "--help=false", tmpDir + "/repo1", tmpDir + "/repo2"}
	err = Execute()
	if err != nil {
		t.Fatal(err)
	}

	// verify renamed within the file
	err = assertSameFiles("testdata/renamed_within_file", tmpDir, fs, t)
	if err != nil {
		t.Fatal(err)
	}

	// move within the file
	os.Args = []string{"sud", "move", "km/newSettings.json", "--path=/p", "--value=/primaryEntityType", "--help=false", tmpDir + "/repo1", tmpDir + "/repo2"}
	err = Execute()
	if err != nil {
		t.Fatal(err)
	}

	// verify moved within the file
	err = assertSameFiles("testdata/renamed", tmpDir, fs, t)
	if err != nil {
		t.Fatal(err)
	}

	// move the file
	os.Args = []string{"sud", "move", "km/newSettings.json", "--path=", "--value=km/settings.json", "--help=false", tmpDir + "/repo1", tmpDir + "/repo2"}
	err = Execute()
	if err != nil {
		t.Fatal(err)
	}

	// verify moved
	err = assertSameFiles("testdata/source", tmpDir, fs, t)
	if err != nil {
		t.Fatal(err)
	}

	// add a file
	os.Args = []string{"sud", "add", "dependencies.json", "--path=", "--value={\n  \"productFeatureIds\": [\n    1,\n    2\n  ]\n}", "--help=false", tmpDir + "/repo1", tmpDir + "/repo2"}
	err = Execute()
	if err != nil {
		t.Fatal(err)
	}

	// verify added
	err = assertSameFiles("testdata/added", tmpDir, fs, t)
	if err != nil {
		t.Fatal(err)
	}

	// add within a file
	values = nil
	os.Args = []string{"sud", "add", "dependencies.json", "--path=/productFeatureIds/-", "--value=304", "--help=false", tmpDir + "/repo1", tmpDir + "/repo2"}
	err = Execute()
	if err != nil {
		t.Fatal(err)
	}

	// verify added with the file
	err = assertSameFiles("testdata/added_within_file", tmpDir, fs, t)
	if err != nil {
		t.Fatal(err)
	}

	// remove within file
	values = nil
	os.Args = []string{"sud", "remove", "dependencies.json", "--path=/productFeatureIds/2", "--value=", "--help=false", tmpDir + "/repo1", tmpDir + "/repo2"}
	err = Execute()
	if err != nil {
		t.Fatal(err)
	}

	// verify removed with the file
	err = assertSameFiles("testdata/added", tmpDir, fs, t)
	if err != nil {
		t.Fatal(err)
	}

	// remove file
	os.Args = []string{"sud", "remove", "dependencies.json", "--path=", "--help=false", tmpDir + "/repo1", tmpDir + "/repo2"}
	err = Execute()
	if err != nil {
		t.Fatal(err)
	}

	// verify removed
	err = assertSameFiles("testdata/source", tmpDir, fs, t)
	if err != nil {
		t.Fatal(err)
	}

	// replace file
	os.Args = []string{"sud", "replace", "km/settings.json", "--value={\n  \"$schema\": \"https://schema.yext.com/config/km/settings/v1\",\n  \"p\": \"yext/restaurant\"\n}", "--path=", "--help=false", tmpDir + "/repo1", tmpDir + "/repo2"}
	err = Execute()
	if err != nil {
		t.Fatal(err)
	}

	// verify replaced
	err = assertSameFiles("testdata/replaced", tmpDir, fs, t)
	if err != nil {
		t.Fatal(err)
	}

	// add using file flag
	values = nil
	os.Args = []string{"sud", "add", "dependencies.json", "--file=testdata/dependencies.json", "--value=", "--path=", "--help=false", tmpDir + "/repo1", tmpDir + "/repo2"}
	err = Execute()
	if err != nil {
		t.Fatal(err)
	}

	// verify added
	err = assertSameFiles("testdata/added_using_file", tmpDir, fs, t)
	if err != nil {
		t.Fatal(err)
	}

	// replace using file flag
	os.Args = []string{"sud", "replace", "dependencies.json", "--file=testdata/dependencies_new.json", "--value=", "--path=", "--help=false", tmpDir + "/repo1", tmpDir + "/repo2"}
	err = Execute()
	if err != nil {
		t.Fatal(err)
	}

	// verify replaced using file flag
	err = assertSameFiles("testdata/replaced_using_file_flag", tmpDir, fs, t)
	if err != nil {
		t.Fatal(err)
	}

	// add using values flag
	values = nil
	os.Args = []string{"sud", "add", "dependencies.json", "--file=", "--values=304,405", "--value=", "--path=/productFeatureIds/-", "--help=false", tmpDir + "/repo1", tmpDir + "/repo2"}
	err = Execute()
	if err != nil {
		t.Fatal(err)
	}

	// verify added multiple values
	err = assertSameFiles("testdata/added_multiple_values", tmpDir, fs, t)
	if err != nil {
		t.Fatal(err)
	}

	// remove matched value 304
	values = nil
	os.Args = []string{"sud", "remove", "dependencies.json", "--value=304", "--path=/productFeatureIds", "--help=false", tmpDir + "/repo1", tmpDir + "/repo2"}
	err = Execute()
	if err != nil {
		t.Fatal(err)
	}

	// remove matched value 405
	values = nil
	os.Args = []string{"sud", "remove", "dependencies.json", "--value=405", "--path=/productFeatureIds", "--help=false", tmpDir + "/repo1", tmpDir + "/repo2"}
	err = Execute()
	if err != nil {
		t.Fatal(err)
	}

	// verify removed the matched values
	err = assertSameFiles("testdata/replaced_using_file_flag", tmpDir, fs, t)
	if err != nil {
		t.Fatal(err)
	}

	// remove using values flag
	os.Args = []string{"sud", "remove", "dependencies.json", "--values=2,3", "--value=", "--path=/productFeatureIds", "--help=false", tmpDir + "/repo1", tmpDir + "/repo2"}
	err = Execute()
	if err != nil {
		t.Fatal(err)
	}

	// verify removed the matched values
	err = assertSameFiles("testdata/removed_using_values_flag", tmpDir, fs, t)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGlob(t *testing.T) {
	// create tmp dir
	tmpDir, err := ioutil.TempDir("", "glob-test")
	fmt.Println("created " + tmpDir)
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		err := os.RemoveAll(tmpDir)
		if err != nil {
			t.Fatal(err)
		}
	}()

	// copy from testdata
	fs := afero.NewOsFs()
	err = copyAllFiles("testdata/glob_source", tmpDir, fs, t)
	if err != nil {
		t.Fatal(err)
	}

	// rename within the file
	os.Args = []string{"sud", "rename", "km/*/*.json", "--path=/apiName", "--value=$id", "--help=false", tmpDir + "/repo1", tmpDir + "/repo2"}
	err = Execute()
	if err != nil {
		t.Fatal(err)
	}

	// verify renamed within the file
	err = assertSameFiles("testdata/glob_renamed", tmpDir, fs, t)
	if err != nil {
		t.Fatal(err)
	}

	// replace within the files
	os.Args = []string{"sud", "replace", "km/*/*.json", "--path=/group", "--value=\"ABC\"", "--file=", "--help=false", tmpDir + "/repo1", tmpDir + "/repo2"}
	err = Execute()
	if err != nil {
		t.Fatal(err)
	}

	// verify replaced within the files
	err = assertSameFiles("testdata/glob_replaced", tmpDir, fs, t)
	if err != nil {
		t.Fatal(err)
	}

	// add within the files
	os.Args = []string{"sud", "add", "km/field/*.json", "--path=/description", "--value=\"description for documentation\"", "--file=", "--help=false", tmpDir + "/repo1", tmpDir + "/repo2"}
	err = Execute()
	if err != nil {
		t.Fatal(err)
	}

	// verify added within the files
	err = assertSameFiles("testdata/glob_added", tmpDir, fs, t)
	if err != nil {
		t.Fatal(err)
	}

	// move within the files
	os.Args = []string{"sud", "move", "km/field/*.json", "--path=/description", "--value=/type/desc", "--help=false", tmpDir + "/repo1", tmpDir + "/repo2"}
	err = Execute()
	if err != nil {
		t.Fatal(err)
	}

	// verify moved within the files
	err = assertSameFiles("testdata/glob_moved", tmpDir, fs, t)
	if err != nil {
		t.Fatal(err)
	}

	// remove within the files
	values = nil
	os.Args = []string{"sud", "remove", "km/field/*.json", "--path=/type/desc", "--value=", "--file=", "--help=false", tmpDir + "/repo1", tmpDir + "/repo2"}
	err = Execute()
	if err != nil {
		t.Fatal(err)
	}

	// verify removed within the files
	err = assertSameFiles("testdata/glob_removed", tmpDir, fs, t)
	if err != nil {
		t.Fatal(err)
	}

	// add within using file flag
	values = nil
	os.Args = []string{"sud", "add", "km/field/*.json", "--path=/type/desc", "--file=testdata/to_add.json", "--value=", "--help=false", tmpDir + "/repo1", tmpDir + "/repo2"}
	err = Execute()
	if err != nil {
		t.Fatal(err)
	}

	// verify added using file flag
	err = assertSameFiles("testdata/glob_moved", tmpDir, fs, t)
	if err != nil {
		t.Fatal(err)
	}

	// remove within using file flag
	values = nil
	os.Args = []string{"sud", "remove", "km/field/*.json", "--path=/type/desc", "--file=testdata/to_add.json", "--value=", "--help=false", tmpDir + "/repo1", tmpDir + "/repo2"}
	err = Execute()
	if err != nil {
		t.Fatal(err)
	}

	// verify removed using file flag
	err = assertSameFiles("testdata/glob_removed", tmpDir, fs, t)
	if err != nil {
		t.Fatal(err)
	}
}

func copyAllFiles(srcDir string, destDir string, fs afero.Fs, t *testing.T) error {
	return afero.Walk(fs, srcDir,
		func(srcPath string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}

			testFile, err := afero.ReadFile(fs, srcPath)
			if err != nil {
				t.Fatal(err)
			}
			relativePath := strings.TrimPrefix(srcPath, srcDir+"/")

			destPath := filepath.Join(destDir, relativePath)
			err = fs.MkdirAll(filepath.Dir(destPath), 0777)
			if err != nil {
				t.Fatal(err)
			}
			err = afero.WriteFile(fs, destPath, testFile, 0644)
			if err != nil {
				t.Fatal(err)
			}
			return nil
		})
}

func assertSameFiles(srcDir string, destDir string, fs afero.Fs, t *testing.T) error {
	err := walk(srcDir, destDir, fs, t)
	if err != nil {
		return err
	}
	err = walk(destDir, srcDir, fs, t)
	if err != nil {
		return err
	}
	return nil
}

func walk(srcDir string, destDir string, fs afero.Fs, t *testing.T) error {
	return afero.Walk(fs, srcDir,
		func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}

			testFile, err := afero.ReadFile(fs, path)
			if err != nil {
				t.Fatal(err)
			}

			relativePath := strings.TrimPrefix(path, srcDir+"/")
			renamedFile, err := afero.ReadFile(fs, filepath.Join(destDir, relativePath))
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, string(testFile), string(renamedFile))
			return nil
		})
}
