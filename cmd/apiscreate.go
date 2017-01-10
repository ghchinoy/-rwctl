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

var (
	serviceid string
	spec string
	endpoint string
)

// apiscreateCmd represents the apiscreate command
var apiscreateCmd = &cobra.Command{
	Use:   "create <name>",
	Short: "create an api on the platform",
	Long: `create an api .`,
	Run: func(cmd *cobra.Command, args []string) {
		// rwctl apis create <name> [--from <serviceID> | --spec <spec>] [--endpoint <endpoint>]

		// Create
		// rwctl apis create APINAME
		// must have a name

		if args[0] == "" {
			fmt.Println("Please provide a name for the API")
			os.Exit(1)
		}
		apiName := args[0]
		from := serviceid

		var cfgmap map[string]interface{}
		var config control.Configuration

		if viper.IsSet(profile) {
			cfgmap = viper.GetStringMap(profile)
		} else {
			fmt.Println("Cannot find profile", profile, " Please check configuration.")
			os.Exit(1)
		}

		if len(args) == 0 {
			fmt.Println("an user or list of users must be given. Please see -h help.")
			os.Exit(1)
		}
		config, err := control.ViperToConfiguration(cfgmap, debug)
		if err != nil {
			fmt.Println("Error translating config", err.Error())
		} else {

			if from != "" {
				// Create from existing service
				// .. --from APIID
				apis.CreateAPIfromExistingService(apiName, from, config, debug)
			} else if spec != "" {
				// Create using a provied spec
				// --spec SPECFILE
				// TODO
				apis.CreateAPIwithSpec(apiName, spec, config, debug)
			} else {
				// Add name only
				if endpoint != "" {
					// ... --endpoint HTTP
					apis.CreateAPINameOnlyWithEndpoint(apiName, endpoint, config, debug)
				} else {
					// rwctl apis create APINAME
					apis.CreateAPINameOnly(apiName, config, debug)
				}
			}
		}
	},
}

func init() {
	apisCmd.AddCommand(apiscreateCmd)

	apiscreateCmd.Flags().StringVar(&serviceid, "serviceid", "","id of service that exists on the platform")
	apiscreateCmd.Flags().StringVarP(&spec, "spec", "s","","api specification (oai, wsdl)")
	apiscreateCmd.Flags().StringVarP(&endpoint, "endpoint", "e", "", "endpoint for api")

}
