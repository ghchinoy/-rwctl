// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
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

	"github.com/spf13/viper"
	"sort"
	"strings"
)

// profilesCmd represents the profiles command
var profileCmd = &cobra.Command{
	Use:   "profile [profile]",
	Short: "profile information",
	Long: `List the details for a named profile as well as available profile configurations found in the .rwctl configuration file. Defaults to 'default'`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			profile = args[0]
		}
		// debug, list out profile information
		fmt.Printf("%7s: %s\n", "profile", profile)
		if viper.IsSet(profile) {
			p := viper.GetStringMap(profile)
			for k, v := range p {
				// don't show password
				if k != "password" {
					fmt.Printf("%7s: %s\n", k,v)
				}
			}
		} else {
			fmt.Printf("No %s profile exists in config file %s.", profile, cfgFile)
		}
		fmt.Println()
		profiles := viper.AllKeys()
		sort.Strings(profiles)
		// remove "profile" default
		profileposition := sort.SearchStrings(profiles, "profile")
		profiles = append(profiles[:profileposition], profiles[profileposition+1:]...)

		fmt.Println("Valid profiles:",  strings.Join(profiles, ", "))
	},
}

func init() {
	RootCmd.AddCommand(profileCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// profilesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// profilesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	profileCmd.PersistentFlags().StringVar(&cfgFile, "config", "", cfgHelp)
	profileCmd.PersistentFlags().StringVar(&profile, "profile", "default", "profile name")
	// Set bash-completion
	validConfigFilenames := []string{"toml", ""}
	profileCmd.PersistentFlags().SetAnnotation("config", cobra.BashCompFilenameExt, validConfigFilenames)

}
