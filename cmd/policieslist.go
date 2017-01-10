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
	"github.com/ghchinoy/rwctl/control"
	"github.com/ghchinoy/rwctl/policies"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// policieslistCmd represents the policieslist command
var policieslistCmd = &cobra.Command{
	Use:   "list",
	Short: "list policies on the platform",
	Long:  `list policies available on the platform`,
	Run: func(cmd *cobra.Command, args []string) {

		var cfgmap map[string]interface{}
		var config control.Configuration

		if viper.IsSet(profile) {
			cfgmap = viper.GetStringMap(profile)
		}

		config, err := control.ViperToConfiguration(cfgmap, debug)
		if err != nil {
			fmt.Println("Error translating config", err.Error())
			os.Exit(1)
		}

		var policytypes string
		if len(args) == 0 {
			policytypes = "all"
		} else {
			policytypes = args[0]
		}
		err = policies.ListPolicies(policytypes, config, debug)
		if err != nil {
			fmt.Println(err.Error())
		}

	},
}

func init() {
	policiesCmd.AddCommand(policieslistCmd)

}
