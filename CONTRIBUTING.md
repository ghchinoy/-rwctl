# Contributing

Guidelines on contributing to rwctl go here.

Currently, there is no Contributor License Agreement (CLA), but there may be one in the future.


## Commands vs Implementation

Commands for the cli are in the `cmd` package, but most implementations are in separate packages. Please follow this convention, so that command implementations can be tested separately from the cli commands themselves.

To add a command, use the [cobra](https://github.com/spf13/cobra) cli convention `cobra add [commandname]`


## Building for release

To build for release, use the `makedist` command.
Please update the VERSION variable in the `makedist` bash script to correspond with the version string variable in the `version/version.go` file.

