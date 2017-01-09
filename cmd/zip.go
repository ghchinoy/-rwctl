// Copyright © 2017 G. Hussain Chinoy <ghchinoy@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"os"
	"strings"
	"github.com/ghchinoy/atmotool/zip"
)

var prefix string

// zipCmd represents the zip command
var zipCmd = &cobra.Command{
	Use:   "zip <directory>",
	Short: "convenience compression",
	Long: `compresses a directory, labeling with given prefix (flag).`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide a directory to compress. Use -h to see help.")
			os.Exit(1)
		}
		dir := args[0]
		var fn string
		if dir == "." {
			fn = "this"
		} else {
			fn = strings.Replace(dir, ".", "", -1)
			fn = strings.Replace(fn, "/", "-", -1)
		}
		dir = strings.TrimSuffix(dir, "/")
		fn = strings.TrimSuffix(fn, "-")
		fn = prefix + "_" + fn + ".zip"
		fmt.Printf("Zipping %s as %s...\n", dir, fn)
		zip.ZipFolder(dir, fn)
	},
}

func init() {
	RootCmd.AddCommand(zipCmd)

	// prefix is a flag applicable only to this command
	zipCmd.Flags().StringVarP(&prefix, "prefix", "p" , "","prefix for zip file")

}
