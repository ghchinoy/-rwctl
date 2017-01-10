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
	"github.com/ghchinoy/rwctl/cms"
	"github.com/spf13/viper"
	"os"
	"github.com/ghchinoy/rwctl/control"
)

// cmslistCmd represents the cmslist command
var cmslistCmd = &cobra.Command{
	Use:   "list [path]",
	Short: "list the cms path",
	Long: `list the contents of a cms path; if path not present, both cms root paths will be listed.`,
	Run: func(cmd *cobra.Command, args []string) {

		var cfgmap map[string]interface{}

		if viper.IsSet(profile) {
			cfgmap = viper.GetStringMap(profile)
		} else {
			fmt.Println("Cannot find profile", profile, " Please check configuration.")
			os.Exit(1)
		}

		config, err := control.ViperToConfiguration(cfgmap, debug)
		if err != nil {
			fmt.Println("Error translating config", err.Error())
			os.Exit(1)
		}

		if len(args) != 0 {
			cms.ListCMSPath(config, args[0], debug)
		} else {
			cms.ListTopLevelCMS(config, debug)
		}

	},
}

func init() {
	cmsCmd.AddCommand(cmslistCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cmslistCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cmslistCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
