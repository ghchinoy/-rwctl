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
	"github.com/spf13/viper"
	"os"

	"github.com/ghchinoy/rwctl/apis"
	"github.com/ghchinoy/rwctl/control"
)

var showversion bool

// apisdetailsCmd represents the apisdetails command
var apisdetailsCmd = &cobra.Command{
	Use:   "details <apiid>",
	Short: "show details of an api",
	Long: `show details about an API.`,
	Run: func(cmd *cobra.Command, args []string) {
		var cfgmap map[string]interface{}
		var config control.Configuration

		if viper.IsSet(profile) {
			cfgmap = viper.GetStringMap(profile)
		} else {
			fmt.Println("Cannot find profile", profile, " Please check configuration.")
			os.Exit(1)
		}

		if len(args) == 0 {
			fmt.Println("an API ID must be given. Please see -h help.")
			os.Exit(1)
		}

		config, err := control.ViperToConfiguration(cfgmap, debug)
		if err != nil {
			fmt.Println("Error translating config", err.Error())
		} else {
			apis.ShowDetailsforAPIID(args[0], showversion, config, debug)
		}
	},
}

func init() {
	apisCmd.AddCommand(apisdetailsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// apisdetailsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// apisdetailsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	apisdetailsCmd.Flags().BoolVarP(&showversion, "version", "v" , false,"show api version details")

}
