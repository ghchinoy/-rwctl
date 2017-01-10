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
	"github.com/ghchinoy/rwctl/apis"
	"github.com/spf13/viper"
	"github.com/ghchinoy/rwctl/control"
)

// metricsCmd represents the metrics command
var metricsCmd = &cobra.Command{
	Use:   "metrics <apiid>",
	Short: "display metrics for an api",
	Long: `summary metrics for an api over a period of time, shows avg, min, max, total calls, total successes, and total failures`,
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
			apis.APIMetrics(args[0], config, debug)
		}
	},
}

func init() {
	apisCmd.AddCommand(metricsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// metricsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// metricsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
