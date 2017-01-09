// Copyright © 2017 G. Hussain Chinoy
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

	"github.com/spf13/cobra"
)

// ApisCmd represents the apis command
var ApisCmd = &cobra.Command{
	Use:   "apis",
	Short: "api commands",
	Long: `commands to manage apis`,

}

func init() {
	RootCmd.AddCommand(ApisCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// apisCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// apisCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	ApisCmd.PersistentFlags().StringVar(&cfgFile, "config", "", cfgHelp)
	ApisCmd.PersistentFlags().StringVar(&profile, "profile", "default", "profile name")
	ApisCmd.PersistentFlags().BoolVar(&debug, "debug", false, "debug output")
	// Set bash-completion
	validConfigFilenames := []string{"toml", ""}
	ApisCmd.PersistentFlags().SetAnnotation("config", cobra.BashCompFilenameExt, validConfigFilenames)


}
