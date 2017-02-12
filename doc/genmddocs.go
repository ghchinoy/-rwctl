package main

import (
"github.com/ghchinoy/rwctl/cmd"
"github.com/spf13/cobra/doc"
)

func main() {
rwctl := cmd.RootCmd
doc.GenMarkdownTree(rwctl, "./")
}
