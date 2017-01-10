# Rogue Wave API Platform CLI


```
rwctl is a CLI tool to manage the Rogue Wave API Platform, including
 APIs, Policies, and API platform and portal settings.

Usage:
  rwctl [command]

Available Commands:
  apis        api commands
  apps        app commands
  policies    policies commands
  portal      portal commands
  profile     profile information
  users       users commands
  version     show the current version
  zip         convenience compression

Flags:
  -h, --help   help for rwctl

Use "rwctl [command] --help" for more information about a command.
```



## Install

OS X, Linux, and Windows binaries are available.

### Option 1 - Use a package manager (preferred)


**On OS X, with homebrew**

```
brew update
brew install ghchinoy/roguewave/rwctl
```

**On Windows, with scoop:**

Add the bucket:

```
scoop bucket add roguewave https://github.com/ghchinoy/scoop-roguewave
scoop bucket list
scoop search rwctl
```

Install `rwctl`:

```
scoop install rwctl
```

Create a config (if you don't have one); see below for format (example, shown using `nano`, which can be installed via `scoop install nano`):

```
new-item -path ~\.config\roguewave -type directory
nano .config\roguewave\rwctl.toml
```

Refer to [scoop.sh](http://scoop.sh/) for more info on scoop.


### Option 2 - Download a release from GitHub

View the [releases](https://github.com/ghchinoy/rwctl/releases) page for the `rwctl` GitHub project, and find the appropriate archive for your operating system and architecture. (For OS X systems, remember to use the darwin archive.)

**OS X / Linux**

Download the archive from your browser or copy the URL and retrieve it via `wget` or `curl`:

```
# OS X
curl -L https://github.com/ghchinoy/rwctl/releases/download/v0.2.1/rwctl-0.2.1.tar.gz | tar xz

# linux, wget
wget -q0- https://github.com/ghchinoy/rwctl/releases/download/v0.2.1/rwctl-0.2.1-linux.amd64
# linux, curl
curl -L https://github.com/ghchinoy/rwctl/releases/download/v0.2.1/rwctl-0.2.1-linux.amd64
```

Rename the binary (as necessary) to `rwctl` and move the `rwctl` binary to your path. Example:

```
sudo mv ./rwctl /usr/local/bin
```

**Windows**

Download the [Windows release](https://github.com/ghchinoy/rwctl/releases/download/v0.2.1/rwctl-0.2.1-windows.amd64.exe) (link is for 0.2.1, check [releases](https://github.com/ghchinoy/rwctl/releases) for latest) and rename to `rwctl.exe`

### Option 3 - Build from source

If you have a Go environment configured, install source like so:

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

Optional elements are `theme` (defaulted to "hermosa") and `console-username` (no default).


## Examples


`rwctl` interacts with the Rogue Wave API Platform to manage various entities (APIs, Apps, etc.) as well as configure the Portal look-and-feel.

* list all APIs on the platform

```
rwctl apis list
```

* upload a file to the Portal cms

```
rwctl portal upload --path /content/home/landing my_landing_page.zip
rwctl portal upload custom.less -p /resources/theme/hermosa/less
```

* rebuild Portal styles

```
rwctl portal rebuild
```

* add an API to the platform, proxying a backend

```
rwctl apis create "JSON Placeholder" --endpoint https://jsonplaceholder.typicode.com/
```

* add an API to the platform with an OAI spec

```
rwctl apis create "Petstore from cli" --spec petstore.json
```

## Contributions

See [CONTRIBUTING.md](CONTRIBUTING.md) for more information on how to contribute.

## License

rwctl is released uner the Apache 2.0 License, see the [LICENSE](LICENSE) file for a full version of the license.