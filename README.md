# Rogue Wave API Platform CLI

## Install

From source

```
go get github.com/ghchinoy/rwctl
go install github.com/ghchinoy/rwctl
```

On OS X, with homebrew

```
brew update
brew install ghchinoy/roguewave/rwctl
```

On Windows, with scoop


## Config file

rwctl expects a valid configuration file in [TOML](https://github.com/toml-lang/toml) format. The configuration file's default location is `$HOME/.config/roguewave/rwctl.toml`

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