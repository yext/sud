$ sud
sud OR solution updater updates solutions by running the command on the file in the given repositories.

Repositories are space separated list of directories or github repository urls.

A repository can be one of the following:
 1. A directory on your file system such as ~/repos/my-solution-dir
 2. A url to github repository such as https://github.com/YextHHChallenges/asbTest
 3. A wildcard such as https://github.com/YextHHChallenges/asbTest* that will be expanded to matched repositories

For example:

sud replace default/km/*.json --value "\"yext/atm\"" --path /primaryEntityType https://github.com/YextHHChallenges/asb* --push

will make the change in the matching files in all the matching repositories and push to github.

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




