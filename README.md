# sud
![Test](https://github.com/yext/sud/workflows/Test/badge.svg)
[![Release][s0]][l0]

[s0]: https://github.com/yext/sud/workflows/Release/badge.svg
[l0]: https://github.com/yext/sud/actions


sud OR solution updater updates solutions by running the command on the file in the given repositories.

Repositories are space separated list of directories or github repository URLs.

A repository can be one of the following:
 1. A directory on your file system such as ~/repos/my-solution-dir
 2. A url to github repository such as https://github.com/MyOrgIsHere/asbTest
 3. A wildcard such as https://github.com/MyOrgIsHere/asbTest* that will be expanded to matched repositories

For example:

```bash
sud replace default/km/*.json \
  --value "\"yext/atm\"" \
  --path /primaryEntityType \
  "https://github.com/MyOrgIsHere/asb*" \
  --push
```

will make the change in the matching files in all the matching repositories and push to github.

Note: To avoid bash expansions, please quote the arguments. To consider bash expansions, please keep the arguments with wildcards unquoted.

Usage:
  sud [command]

Available Commands:
  add         adds a file, or a path in a file, in the given list of repositories
  help        Help about any command
  move        moves a file, or a path in a file, in the given list of repositories
  remove      removes a file, or a path in a file, in the given list of repositories
  rename      renames a file, or a path in a file, in the given list of repositories
  replace     replaces a file, or a path in a file, in the given list of repositories

Flags:
  -h, --help   help for sud
      --push   Whether to push the repo to github

Use "sud [command] --help" for more information about a command.

## Table of Contents

* [Install](#install)  
* [Update](#update)
* [Access](#access)
* [Run](#run)

## Install

Sud requires Mac OS or Go 1.7 or higher.

    brew install yext/tap/sud
    go get github.com/yext/sud

## Update

To update an existing install to the latest version of Sud, run:

    brew upgrade yext/tap/sud
    go get -u github.com/yext/sud
## Access
Please make sure you have access to push to the repositories you are trying to update using Sud.
To update Github repos in specific orgs, you may need to create a personal token that has permissions to push to the repos: https://docs.github.com/en/github/authenticating-to-github/keeping-your-account-and-data-secure/creating-a-personal-access-token

## Run
### `sud rename` command

sud rename --help
Renames a file, or a path in a file, in the given repositories:

Renames the given file in all the given repositories, if the file is present.
Renames within a file if the path flag is provided.

Examples:

Renaming a File:

```bash
sud rename dependencies.json \
  --value foo.json \
  ~/repos/ans*
```
will rename dependencies.json to foo.json in all ~/repos/ans* directories.

Renaming a Path:

```bash
sud rename km/*/*.json \
  --path /apiName \
  --value "\$id" \
  ~/repos/ans*
```
will replace "apiName" with "$id" in km/*/*.json files in all ~/repos/ans* directories:

{ "apiName": "123" } => { "$id": "123" }

Usage:
  sud rename FILE_PATH DIRs... [flags]

Flags:
  -h, --help           help for rename
  -p, --path string    The path to the json pointer in the file as per https://tools.ietf.org/html/rfc6901
  -v, --value string   The desired value

Global Flags:
      --push   Whether to push the repo to github

### `sud move` command

Moves a file, or a path in a file, in the given repositories :

Moves the given file in all the given repositories, if the file is present.
Moves within a file if the path flag is provided.

Examples:

Moving a File:

```bash
sud move km/entity-type/atm.json \
  --value km/entity-type-extension/atm.json \
  ~/repos/ans150-front-end-overview \
  ~/repos/kg*
```
will move km/entity-type/atm.json to km/entity-type-extension/atm.json in the ans150-front-end-overview repo as well as any kg* repos.

Moving within a File:

```bash
sud move \
  --path /a/b/c \
  --value /a/d \
  ~/repos/ans150-front-end-overview
```
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
}

Usage:
  sud move FILE_PATH DIRs... [flags]

Flags:
  -h, --help           help for move
  -p, --path string    The path to the json pointer in the file as per https://tools.ietf.org/html/rfc6901
  -v, --value string   The desired value

Global Flags:
      --push   Whether to push the repo to github

### `sud add` command

Adds a file, or a path in a file, in the given repositories :

 - Adds the given file in all the given repositories. It adds only if the file is not already present.
 - Adds within a file if the path flag is provided.

The value provided should conform to the intended type.
For example,
 1. to add a number would be --value=304
 2. to add string would be --value=\"my-string\" (escape the quote in bash)
 3. to add an array would be --value="[1, 2]"

Instead of providing a value you can also provide a file path of a file that has the desired value.

Examples:

Adding a File:

```bash
sud add foo.json \
  --file bar.json \
  ~/repos/ans150-front-end-overview
```
will add a file foo.json with the contents of the provided bar.json to ~/repos/ans150-front-end-overview.

Adding a Path:

```bash
sud add dependencies.json \
  --path /productFeatureIds/- \
  --values 304,305 \
  ~/repos/ans150-front-end-overview
```
will add product features 304 and 305 to ~repos/ans150-front-end-overview/dependencies.json:

{ "productFeatureIds": [1, 2] } => { "productFeatureIds": [1, 2, 304, 305] }

Usage:
  sud add FILE_PATH DIRs... [flags]

Flags:
  -f, --file string      The file with the desired value
  -h, --help             help for add
  -p, --path string      The path to the json pointer in the file as per https://tools.ietf.org/html/rfc6901
  -v, --value string     The desired value
  -l, --values strings   The desired list of values

Global Flags:
      --push   Whether to push the repo to github
      
### `sud remove` command

Removes a file, or a path in a file, in the given repositories :

Removes the given file in all the given repositories, if the file is present.
Removes within a file if the path flag is provided.
If value is provided along with the path, then removes only the provided value, if present at the path. (Currently, only removing a number from an array is supported.)

Examples:

Remove a File:

```bash
sud remove dependencies.json \
  ~/repos/ans150-front-end-overview \
  ~/repos/kg*
```
Removes dependencies.json from all of the specified directories.

Removing a Value in a File:

```bash
sud remove dependencies.json \
  --path /productFeatureIds \
  --value 304 \
  ~/repos/ans150-front-end-overview \
  ~/repos/kg*
```
will remove product feature 304 from dependencies.json in all the specified directories:

{ "productFeatureIds": [1, 2, 304] } => { "productFeatureIds": [1 ,2] }


Removing fields from entities:

```bash
sud remove default/km/entity/*.json \
  --path /content/facebookParentPageId \
  ~/repos/kg*
```
will remove the given field from all the entities in the matched directories.

Usage:
  sud remove FILE_PATH DIRs... [flags]

Flags:
  -f, --file string      The file with the desired value
  -h, --help             help for remove
  -p, --path string      The path to the json pointer in the file as per https://tools.ietf.org/html/rfc6901
  -v, --value string     The desired value
  -l, --values strings   The desired list of values

Global Flags:
      --push   Whether to push the repo to github
      
### `sud replace` command

Replaces a file, or a path in a file, in the given repositories :

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

```bash
sud replace pages/page-builder/template/page.json \
  --file new-page.json \
  ~/repos/pgs*
```

will replace pages/page-builder/template/page.json with the contents of new-page.json in all ~/repos/pgs* directories.

Replacing a Path:

```bash
sud replace km/*/hotel.json \
  --path "/\$id" \
  --value "\"my_hotel\"" \
  ~/repos/kg122-intro-to-fields
```

will replace the value at path /$id with "my_hotel" in ~/repos/kg122-intro-to-fields/km/entity-type-extension/hotel.json

```bash
sud replace km/*/hotel.json \
  --path /enabled \
  --value true \
  ~/repos/kg122-intro-to-fields
```

will replace the value at path /enabled with true in ~/repos/kg122-intro-to-fields/km/entity-type-extension/hotel.json

Usage:
  sud replace FILE_PATH DIRs... [flags]

Flags:
  -f, --file string    The file with the desired value
  -h, --help           help for replace
  -p, --path string    The path to the json pointer in the file as per https://tools.ietf.org/html/rfc6901
  -v, --value string   The desired value

Global Flags:
      --push   Whether to push the repo to github

