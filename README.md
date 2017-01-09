# Rogue Wave API Platform CLI


```
rwctl is a CLI tool to manage the Rogue Wave API Platform, including
 APIs, Policies, and API platform and portal settings.

Usage:
  rwctl [command]

Available Commands:
  apis        api commands
  portal      portal commands
  profile     profile information
  version     show the current version
  zip         convenience compression

Flags:
  -h, --help   help for rwctl

Use "rwctl [command] --help" for more information about a command.
```



## Install

### Option 1 - Use a package manager (preferred)


On OS X, with homebrew

```
brew update
brew install ghchinoy/roguewave/rwctl
```

On Windows, with scoop

(coming soon)

### Option 2 - Download from GitHub release

View the [releases](/releases) page for the `rwctl` GitHub project, and find the appropriate archive for your operating system and architecture. (For OS X systems, remember to use the darwin archive.)

### Option 3 - From source

```
go get github.com/ghchinoy/rwctl
go install github.com/ghchinoy/rwctl
```



## Config file

rwctl expects a valid configuration file in [TOML](https://github.com/toml-lang/toml) format. The configuration file's default location is `$HOME/.config/roguewave/rwctl.toml` but a file location can be specified with the `--config` flag.

Example config file

```
[default]
url = "http://portal.roguewave.dev:9980"
email = "administrator@roguewave.com"
password = "password"
theme = "hermosa"

[portal2]
url = "http://partners.roguewave.dev:9980"
email = "administrator@roguewave.com"
password = "password"
console-username = "partners-HussainChinoy"
```

You may define multiple API Platform targets with different TOML blocks.

Optional elements are `theme` and `console-username`.


## License

rwctl is released uner the Apache 2.0 License, see the [LICENSE](LICENSE) file for a full version of the license.